package downloader

import (
	"fmt"

	"github.com/kowala-tech/kUSD/core/types"
)

// peerDropFn is a callback type for dropping a peer detected as malicious.
type peerDropFn func(id string)

// dataPack is a data message returned by a peer for some query.
type dataPack interface {
	PeerId() string
	Items() int
	Stats() string
}

// headerPack is a batch of block headers returned by a peer.
type headerPack struct {
	peerId  string
	headers []*types.Header
}

func (p *headerPack) PeerId() string { return p.peerId }
func (p *headerPack) Items() int     { return len(p.headers) }
func (p *headerPack) Stats() string  { return fmt.Sprintf("%d", len(p.headers)) }

// bodyPack is a batch of block bodies returned by a peer.
type bodyPack struct {
	peerId       string
	transactions [][]*types.Transaction
}

func (p *bodyPack) PeerId() string { return p.peerId }
func (p *bodyPack) Items() int     { return len(p.transactions) }
func (p *bodyPack) Stats() string  { return fmt.Sprintf("%d", len(p.transactions) }

// receiptPack is a batch of receipts returned by a peer.
type receiptPack struct {
	peerId   string
	receipts [][]*types.Receipt
}

func (p *receiptPack) PeerId() string { return p.peerId }
func (p *receiptPack) Items() int     { return len(p.receipts) }
func (p *receiptPack) Stats() string  { return fmt.Sprintf("%d", len(p.receipts)) }

// statePack is a batch of states returned by a peer.
type statePack struct {
	peerId string
	states [][]byte
}

func (p *statePack) PeerId() string { return p.peerId }
func (p *statePack) Items() int     { return len(p.states) }
func (p *statePack) Stats() string  { return fmt.Sprintf("%d", len(p.states)) }
