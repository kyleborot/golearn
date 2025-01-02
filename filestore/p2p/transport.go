package p2p

//Peer is an interface that represents the remote node
//(the node we are connecting to)
type Peer interface {
}

//Transport is anything that handles the communication
//between the nodes in the network
//form (TCP, UDP, websockets, etc.)
type Transport interface {
}
