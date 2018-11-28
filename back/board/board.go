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
func (board Board) CheckRulesAndCaptured(initialStone pb.Node) *rules.ReportCheckRules {
	report := rules.New(initialStone.Player)
	defer board.updateBoardAfterCapture(&report)
	board.proccessRulesByAxes(report.ProccessCheckRules, initialStone)
	if report.Report.NbFreeThree > 1 {
		return nil
	} else if lenWinOrLose := len(report.Report.WinOrLose); lenWinOrLose != 0 {
	loop:
		for i := 0; i < lenWinOrLose; i++ {
			for _, checkedStone := range report.Report.WinOrLose[i] {
				board.proccessRulesByAxes(report.CheckIfPartyIsFinish, *checkedStone)
			}
			if len(report.Report.NextMovesOrLose) == 0 {
				report.Report.PartyFinish = true
				break loop
			}
			report.Report.WinOrLose[i] = report.Report.NextMovesOrLose
			report.Report.NextMovesOrLose = make([]*pb.Node, 0, 0)
		}
	}
	board[initialStone.X][initialStone.Y] = initialStone.Player
	return report.Report
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

func (board Board) proccessRulesByAxes(m func(list []*pb.Node, index int), initialStone pb.Node) {
	for index := 0; index < 4; index++ {
		m(board.createListCheckStone(index, initialStone))
	}
}

func (board Board) updateBoardAfterCapture(report *rules.Schema) {
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
}
