package p2p

// Peer is an interface that represents the remote node.
type Peer interface{}

// transport is anything that handles the communication btw the nodes in the network. This can be of the form (TCP, UDP, Websockets, ...)
type Tranport interface {
	ListenAndAccept() error
}
