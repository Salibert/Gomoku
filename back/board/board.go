package board

import (
	"github.com/Salibert/Gomoku/back/axis"
	"github.com/Salibert/Gomoku/back/rules"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// SizeBoard is the size of the board
const SizeBoard int32 = 19

// Board ...
type Board [][]int32

// New ...
func New() Board {
	board := make(Board, 19, 19)
	for i := 0; i < 19; i++ {
		board[i] = make([]int32, 19, 19)
	}
	return board
}

// CheckRulesAndCaptured ...
func (board Board) CheckRulesAndCaptured(initialStone pb.Node) *pb.CheckRulesResponse {
	res := &pb.CheckRulesResponse{}
	report := rules.New(initialStone.Player)
	defer func() {
		if len := len(report.Report.ListCapturedStone); len != 0 {
			go func(list []*pb.Node, len int) {
				var node *pb.Node
				for len > 0 {
					node = list[len-1]
					board[node.X][node.Y] = node.Player
					len--
				}
			}(report.Report.ListCapturedStone, len)
		}
	}()
	for index := 0; index < 4; index++ {
		report.ProccessCheckRules(board.createListCheckStone(index, initialStone))
	}
	if report.Report.NbFreeThree > 1 {
		return res
	}
	board[initialStone.X][initialStone.Y] = initialStone.Player
	res.IsPossible = true
	res.Captered = report.Report.ListCapturedStone
	return res
}

func (board Board) createListCheckStone(index int, initialStone pb.Node) ([]*pb.Node, int) {
	listCheckStone := make([]*pb.Node, 0, 11)
	var indexStone int
	axisCheck := axis.DialRightAxes[index]
	Y, X := initialStone.Y+(axisCheck.Y*5), initialStone.X+(axisCheck.X*5)
	for i := 5; i > 0; i-- {
		if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard {
			listCheckStone = append(listCheckStone, &pb.Node{X: X, Y: Y, Player: board[X][Y]})
		}
		Y -= axisCheck.Y
		X -= axisCheck.X
	}
	indexStone = len(listCheckStone)
	listCheckStone = append(listCheckStone, &initialStone)
	axisCheck = axisCheck.Inverse()
	Y, X = initialStone.Y+axisCheck.Y, initialStone.X+axisCheck.X
	for i := 0; i < 5; i++ {
		if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard {
			listCheckStone = append(listCheckStone, &pb.Node{X: X, Y: Y, Player: board[X][Y]})
		}
		Y += axisCheck.Y
		X += axisCheck.X
	}
	return listCheckStone, indexStone
}
