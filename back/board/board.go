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
	var listStone []*pb.Node
	var backStone *pb.Node
	var freeThree int32
	for _, axis := range axis.Axes {
		backStone, listStone = board.createArrayStoneCheck(&initialStone, axis)
		switch listStone[1].Player {
		case 0:
			freeThree += checkFreeThree(listStone, backStone)
		case initialStone.Player:
			freeThree += board.checkFreeThreeOrWin(listStone, axis, backStone)
		default:
			captured(listStone, listCapture)
		}
	}
	return listCapture
}

func (board Board) createArrayStoneCheck(initialStone *pb.Node, axis axis.Axis) (*pb.Node, []*pb.Node) {
	listStone := make([]*pb.Node, 1, 5)
	listStone[0] = initialStone
	len := len(listStone)
	X, Y := initialStone.X, initialStone.Y
loop:
	for i := 0; i < len; i++ {
		X, Y = X+axis.X, Y+axis.Y
		if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard {
			listStone = append(listStone, &board[X][Y])
		} else {
			break loop
		}
	}
	X, Y = initialStone.X-axis.X, initialStone.Y-axis.Y
	if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard && board[X][Y].Player != 0 {
		return &board[X][Y], listStone
	}
	return nil, listStone
}

func checkFreeThree(listStone []*pb.Node, backStone *pb.Node) int32 {
	len := len(listStone)
	if len >= 4 && (backStone == nil) &&
		listStone[2].Player == listStone[0].Player &&
		listStone[3].Player == listStone[0].Player &&
		(len == 5 && (listStone[4].Player == listStone[0].Player || listStone[4].Player == 0)) {
		return 1
	}
	return 0

}

func (board Board) checkFreeThreeOrWin(listStone []*pb.Node, axis axis.Axis, backStone *pb.Node) (int32, []*pb.Node) {
	nbStone := 1
	len := len(listStone)
loop:
	for i := 1; i < len; i++ {
		if listStone[i].Player != listStone[0].Player {
			break loop
		}
		nbStone++
	}
	if (nbStone == 3 || nbStone == 4) && backStone == nil {
		return 1, false
	} else if nbStone == 5 ||
		(nbStone == 4 && backStone != nil && backStone.Player == listStone[0].Player) {
		return 0, true
	}
	return 0, false
}

func captured(listStone []*pb.Node, listCapture []*pb.Node) {
	if len(listStone) > 3 &&
		listStone[3].Player != 0 &&
		listStone[3].Player != listStone[0].Player &&
		listStone[2].Player != 0 &&
		listStone[2].Player != listStone[0].Player {
		listCapture = append(listCapture, listStone[1], listStone[2])
	}
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
