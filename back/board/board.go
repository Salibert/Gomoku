package board

import (
	"github.com/Salibert/Gomoku/back/axis"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// SizeBoard is the size of the board
const SizeBoard int32 = 19

// Board ...
type Board [][]pb.Node

// Captured check if the new stone captured other adverse stone
func (board Board) Captured(initialStone pb.StonePlayed) []pb.Node {
	var stone *pb.Node
	for axis := range axis.Axes {
		if isFind := board.findStone(initialStone, axis.Multiply(2), stone) == true {
			
		}
	}
}

func (board Board) findStone(initialStone pb.Node, axis axis.Axis, stone *pb.Node) bool {
	stone.X = initialStone.X + axis.X
	stone.Y = initialStone.Y + axis.Y
	if stone.X >= 0 && stone.X < SizeBoard && stone.Y >= 0 && stone.Y < SizeBoard {
		stone.Player = board[stone.X][stone.Y].Player
		return true
	}
	return false
}
