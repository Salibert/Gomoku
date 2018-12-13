package inter

import (
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// Node is an interface to avoid typing grpc
type Node struct {
	X      int
	Y      int
	Player int
}

// New pass grpc Node to inter.Node
func NewNode(n *pb.Node) *Node {
	return &Node{
		X:      int(n.X),
		Y:      int(n.Y),
		Player: int(n.Player),
	}
}

// Convert pass inter.Node to grpc Node
func (node *Node) Convert() *pb.Node {
	return &pb.Node{X: int32(node.X), Y: int32(node.Y), Player: int32(node.Player)}
}

// ConvertArrayNode ...
func ConvertArrayNode(array []*Node) []*pb.Node {
	lenArray := len(array)
	newArray := make([]*pb.Node, lenArray, lenArray)
	for i := 0; i < lenArray; i++ {
		newArray[i] = array[i].Convert()
	}
	return newArray
}
