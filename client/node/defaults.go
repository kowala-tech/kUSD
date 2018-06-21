package node

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/kowala-tech/kcoin/client/p2p"
	"github.com/kowala-tech/kcoin/client/p2p/nat"
)

const (
	DefaultHTTPHost = "localhost" // Default host interface for the HTTP RPC server
	DefaultHTTPPort = 11223       // Default TCP port for the HTTP RPC server
	DefaultWSHost   = "localhost" // Default host interface for the websocket RPC server
	DefaultWSPort   = 8546        // Default TCP port for the websocket RPC server
)

// DefaultConfig contains reasonable default settings.
var DefaultConfig = Config{
	DataDir:     DefaultDataDir(),
	HTTPPort:    DefaultHTTPPort,
	HTTPModules: []string{"net", "web3"},
	WSPort:      DefaultWSPort,
	WSModules:   []string{"net", "web3"},
	P2P: p2p.Config{
		ListenAddr:      ":22334",
		DiscoveryV5Addr: ":30304",
		MaxPeers:        25,
		NAT:             nat.Any(),
	},
}

// DefaultDataDir is the default data directory to use for the databases and other
// persistence requirements.
func DefaultDataDir() string {
	// Try to place the data folder in the user's home dir
	home := homeDir()
	if home != "" {
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "kcoin")
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "kcoin")
		} else {
			return filepath.Join(home, ".kcoin")
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
