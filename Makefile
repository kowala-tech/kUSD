# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

PWD   := $(shell pwd)
GOBIN = $(PWD)/client/build/bin

.PHONY: kcoin
kcoin: bindings
	cd client; build/env.sh go run build/ci.go install ./cmd/kcoin
	@echo "Done building."
	@echo "Run \"$(GOBIN)/kcoin\" to launch kcoin."

.PHONY: control
control:
	cd client; build/env.sh go run build/ci.go install ./cmd/control
	@echo "Done building."
	@echo "Run \"$(GOBIN)/control\" to launch control."

.PHONY: bootnode
bootnode: bindings
	cd client; build/env.sh go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch bootnode."

.PHONY: faucet
faucet: bindings
	cd client; build/env.sh go run build/ci.go install ./cmd/faucet
	@echo "Done building."
	@echo "Run \"$(GOBIN)/faucet\" to launch faucet."

.PHONY: evm
evm:
	cd client; build/env.sh go run build/ci.go install ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm\" to start the evm."

.PHONY: genesis
genesis:
	cd client; build/env.sh go run build/ci.go install ./cmd/genesis
	@echo "Done building."
	@echo "Run \"$(GOBIN)/genesis\" to generate genesis files."

.PHONY: all
all:
	cd client; build/env.sh go run build/ci.go install

.PHONY: android
android:
	cd client; build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/kcoin.aar\" to use the library."

.PHONY: ios
ios:
	cd client; build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Kusd.framework\" to use the library."

.PHONY: test
test: all
	cd client; build/env.sh go run build/ci.go test

.PHONY: test_genesis
test_genesis: all
	cd client/knode/genesis && go test . || curl -X POST -H 'Content-type: application/json' --data '{"attachments":[{"actions":[{"type":"button","text":"Build link","url":"${DRONE_BUILD_LINK}"}],"title":"Build failure","pretext":"Some days, you just cant get rid of a bomb!","text":"*Network config has changed!* Regenerate golden files, this will also require a network restart.","mrkdwn_in":["text","pretext","title"],"color":"#ff0000"}]}' ${SLACK_APP_WEBHOOK}

.PHONY: test_notifications
test_notifications: dep
	cd notifications && \
	$(GOPATH)/bin/dep ensure --vendor-only && \
	go test ./... -tags=integration

.PHONY: lint
lint: all
	cd client; build/env.sh go run build/ci.go lint

.PHONY: clean
clean:
	rm -fr client/build/_workspace/pkg/ $(GOBIN)/* client/build/bin/abigen client/contracts/truffle/node_modules

# Bindings tools

# FILES is the list of binding files that would be created when generating the bindings
FILES=$(shell egrep -ir "go:generate" client/contracts/bindings | grep abigen | sed -E 's/^client\/contracts\/bindings\/(.*)\/.*\.go.*-out\ \.?\/?(.*)/client\/contracts\/bindings\/\1\/\2/' )
$(FILES):
	$(MAKE) -j 5 stringer go-bindata gencodec client/build/bin/abigen client/contracts/truffle/node_modules
	go generate ./client/contracts/bindings/...
.PHONY: bindings
bindings: | $(FILES)

client/contracts/truffle/node_modules:
	cd client/contracts/truffle && npm i

client/build/bin/abigen:
	cd client; build/env.sh go run build/ci.go install ./cmd/abigen

# Cross Compilation Targets (xgo)

.PHONY: kcoin_cross
kcoin_cross: kcoin_cross_build kcoin_cross_compress kcoin_cross_rename
	@echo "Full cross compilation done."

.PHONY: kcoin_cross_build
kcoin_cross_build: bindings
	cd client; build/env.sh go run build/ci.go xgo -- --go=latest --targets=linux/amd64,linux/arm64,darwin/amd64,windows/amd64 -v ./cmd/kcoin
	mv client/build/bin/kcoin-darwin-10.6-amd64 client/build/bin/kcoin-osx-10.6-amd64

.PHONY: kcoin_cross_compress
kcoin_cross_compress:
	cd client/build/bin; for f in kcoin*; do zip $$f.zip $$f; rm $$f; done; cd -

.PHONY: kcoin_cross_rename
kcoin_cross_rename:
ifdef DRONE_TAG
	cd client/build/bin && for f in kcoin-*; do \
		release=$$(echo $$f | awk '{ gsub("kcoin", "kcoin-stable"); print }');\
		version=$$(echo $$f | awk '{ gsub("kcoin", "kcoin-$(DRONE_TAG)"); print }');\
		cp $$f $$release;\
		mv $$f $$version;\
	done;
else
	cd client/build/bin && for f in kcoin-*; do \
		release=$$(echo $$f | awk '{ gsub("kcoin", "kcoin-unstable"); print }');\
		version=$$(echo $$f | awk '{ gsub("kcoin", "kcoin-$(DRONE_COMMIT_SHA)"); print }');\
		cp $$f $$release;\
		mv $$f $$version;\
	done;
endif

## E2E tests

.PHONY: e2e
e2e: dep
	cd e2e && \
	$(GOPATH)/bin/dep ensure --vendor-only && \
	go build -a && \
	./e2e --features ./features

## Wallet app

.PHONY: wallet_app_tests
wallet_app_tests:
	@cd wallet-app; \
	yarn install --network-concurrency 1 && \
	yarn run lint && \
	yarn run test

## Docs
BUILD_DOCS := mkdocs build --clean --strict -d site
.PHONY: build_docs
build_docs:
	@cd docs; $(BUILD_DOCS)

.PHONY: build_docs_with_docker
build_docs_with_docker:
	@docker run --rm -v $(PWD)/docs:/documents kowalatech/mkdocs $(BUILD_DOCS)

## Dev docker images

.PHONY: dev_docker_images
dev_docker_images: dev_explorer_docker_image dev_explorer_sync_docker_image dev_kusd_docker_image dev_bootnode_docker_image dev_faucet_docker_image dev_wallet_backend_docker_image dev_transactions_persistance_docker_image dev_transactions_publisher_docker_image dev_backend_api_docker_image

.PHONY: dev_kusd_docker_image
dev_kusd_docker_image:
	docker build -t kowalatech/kusd:dev -f client/release/kcoin.Dockerfile .

.PHONY: dev_bootnode_docker_image
dev_bootnode_docker_image:
	docker build -t kowalatech/bootnode:dev -f client/release/bootnode.Dockerfile .

.PHONY: dev_faucet_docker_image
dev_faucet_docker_image:
	docker build -t kowalatech/faucet:dev -f client/release/faucet.Dockerfile .

.PHONY: dev_wallet_backend_docker_image
dev_wallet_backend_docker_image:
	docker build -t kowalatech/wallet_backend:dev -f wallet-backend/Dockerfile .

.PHONY: dev_transactions_persistance_docker_image
dev_transactions_persistance_docker_image:
	docker build -t kowalatech/transactions_persistance:dev -f notifications/transactions_db_synchronize.Dockerfile .

.PHONY: dev_transactions_publisher_docker_image
dev_transactions_publisher_docker_image:
	docker build -t kowalatech/transactions_publisher:dev -f notifications/transactions_publisher.Dockerfile .

.PHONY: dev_backend_api_docker_image
dev_backend_api_docker_image:
	docker build -t kowalatech/backend_api:dev -f notifications/api.Dockerfile .

.PHONY: dev_explorer_docker_image
dev_explorer_docker_image:
	docker build -t kowalatech/kexplorer -f explorer/web.Dockerfile .

.PHONY: dev_explorer_sync_docker_image
dev_explorer_sync_docker_image:
	docker build -t kowalatech/kexplorersync -f explorer/sync.Dockerfile .

# Tools

.PHONY: dep
DEP_BIN := $(shell command -v dep 2> /dev/null)
dep:
ifndef DEP_BIN
	@echo "Installing dep..."
	@go get github.com/golang/dep/cmd/dep
endif

.PHONY: stringer
STRINGER_BIN := $(shell command -v stringer 2> /dev/null)
stringer:
ifndef STRINGER_BIN
	@echo "Installing stringer..."
	@go get golang.org/x/tools/cmd/stringer
endif

.PHONY: go-bindata
GO_BINDATA_BIN := $(shell command -v go-bindata 2> /dev/null)
go-bindata:
ifndef GO_BINDATA_BIN
	@echo "Installing go-bindata..."
	@go get github.com/jteeuwen/go-bindata/go-bindata
endif

.PHONY: gencodec
GENCODEC_BIN := $(shell command -v gencodec 2> /dev/null)
gencodec:
ifndef GENCODEC_BIN
	@echo "Installing gencodec..."
	@go get github.com/fjl/gencodec
endif
