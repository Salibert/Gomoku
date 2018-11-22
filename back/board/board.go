package board

import (
	"github.com/Salibert/Gomoku/back/axis"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// SizeBoard is the size of the board
const SizeBoard int32 = 19

// Board ...
type Board [][]int

// CheckRulesAndCaptured ...
func (board Board) CheckRulesAndCaptured(initialStone pb.Node) []*int {
	listCapture := make([]pb.Node, 0, 16)
	var listCheckStone []pb.Node
	var indexStone int
	report := rules.New(initialStone.Player)
	for index := 0; index < 4; index++ {
		report.ProccessCheckRules(board.createListCheckStone(index, initialStone))
	}
	return listCapture
}

func (board Board) createListCheckStone(index int, initialStone pb.Node) []pb.Node, int {
	listCheckStone := make([]pb.Node, 1, 11)
	var indexStone int
	axisCheck := axis.DialRightAxes[index]
	Y, X := initialStone.Y + (axisCheck.Y * 5), initialStone.X + (axisCheck.X * 5)
	for i := 5; i > 0; i-- {
		if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard {
			listCheckStone = append(listCheckStone, pb.Node{X: X, Y: Y, Player: board[X][Y]})
		}
		Y -= axisCheck.Y
		X -= axisCheck.X
	}
	indexStone = len(listCheckStone)
	listCheckStone = append(listCheckStone, initialStone)
	axisCheck = axisCheck.Inverse()
	Y, X := initialStone.Y + (axisCheck.Y * 5), initialStone.X + (axisCheck.X * 5)
	for i := 5; i > 0; i-- {
		if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard {
			listCheckStone = append(listCheckStone, pb.Node{X: X, Y: Y, Player: board[X][Y]})
		}
		Y -= axisCheck.Y
		X -= axisCheck.X
	}
	return listCheckStone, indexStone
}
