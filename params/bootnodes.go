package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Kowala network.
var MainnetBootnodes = []string{}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// test network.
var TestnetBootnodes = []string{
	"enode://a46e5885d3da52bf452fcea7269b91059d93819e7e906cb5d29575508a18a4009e5d688bf29abf8cdd55aad65a9ca4beab98de8c0a226d694c8140acdedf3a55@18.219.208.190:33447",
	"enode://9791d40d97585d7241edeedd2491187e82e3273414cf1bed2aca049e8b21d96930437efe3e591d6996ee303d3e2fe4338e53551bb6e50f05dfeeaccd660eac01@18.220.38.32:33447",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
