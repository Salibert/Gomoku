package board

import (
	"github.com/Salibert/Gomoku/back/axis"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// SizeBoard is the size of the board
const SizeBoard int32 = 19

// Board ...
type Board [][]pb.Node

// CheckRulesAndCaptured ...
func (board Board) CheckRulesAndCaptured(initialStone pb.Node) []*pb.Node {
	listCapture := make([]*pb.Node, 0, 16)
	var freeThree int
	var X, Y int32
	for _, axis := range axis.Axes {
		X, Y = initialStone.X+axis.X, initialStone.Y+axis.Y
		if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard {
			stone := &board[X][Y]
			switch stone.Player {
			case 0:
				freeThree += board.checkFreeThree(stone, axis)
			case initialStone.Player:
				freeThree += board.checkFreeThreeOrWin(stone, axis)
			default:
				board.captured(stone, axis, listCapture)
			}
		}
	}
	return listCapture
}

// Captured check if the new stone captured other adverse stone
func (board Board) captured(stone *pb.Node, axis axis.Axis, listCapture []*pb.Node) {
	X, Y := stone.X+axis.X, stone.Y+axis.Y
	if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard {
		stoneFriendly := &board[X][Y]
		if stoneFriendly.Player != 0 && stoneFriendly.Player != stone.Player {
			if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard {
				X, Y = X+axis.X, Y+axis.Y
				stoneCaptured := &board[X][Y]
				if stoneCaptured.Player != 0 && stoneCaptured.Player != stone.Player {
					listCapture = append(listCapture, stone, stoneCaptured)
				}
			}
		}
	}
	return listCapture
}

func (board Board) findStone(initialStone pb.Node, axis axis.Axis, stone *pb.Node) {
	stone.X = initialStone.X + axis.X
	stone.Y = initialStone.Y + axis.Y
	if stone.X >= 0 && stone.X < SizeBoard && stone.Y >= 0 && stone.Y < SizeBoard {
		stone.Player = board[stone.X][stone.Y].Player
	} else {
		stone.Player = 0
	}
}
