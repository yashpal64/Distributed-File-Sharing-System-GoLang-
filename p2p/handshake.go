package p2p

//HanshakeFunc is
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
