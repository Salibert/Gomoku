package board

import (
	"fmt"

	"github.com/Salibert/Gomoku/back/axis"
	"github.com/Salibert/Gomoku/back/rules"
	"github.com/Salibert/Gomoku/back/server/inter"
)

// SizeBoard is the size of the board
const SizeBoard int = 19

// Board ...
type Board [][]int

// New ...
func New() Board {
	board := make(Board, 19, 19)
	for i := 0; i < 19; i++ {
		board[i] = make([]int, 19, 19)
	}
	return board
}

// UpdateBoard board with new stone
func (board Board) UpdateBoard(stone inter.Node) {
	board[stone.X][stone.Y] = stone.Player
}

// CheckRules check all rules. Modify the report passed in params
func (board Board) CheckRules(initialStone inter.Node, report rules.Schema) {
	board.proccessRulesByAxes(report.ProccessCheckRules, initialStone)
	if report.Report.NbFreeThree > 1 {
		report.Report.ItIsAValidMove = false
		return
	}
	report.Report.ItIsAValidMove = true
	if lenWinOrLose := len(report.Report.WinOrLose); lenWinOrLose != 0 {
		if report.Schema[rules.Capture] != nil {
		loop:
			for i := 0; i < lenWinOrLose; i++ {
				for _, checkedStone := range report.Report.WinOrLose[i] {
					board.proccessRulesByAxes(report.CheckIfPartyIsFinish, *checkedStone)
				}
				if len(report.Report.NextMovesOrLose) == 0 {
					report.Report.PartyFinish = true
					report.Report.WinOrLose[i] = report.Report.WinOrLose[i][:0]
					break loop
				}
				report.Report.WinOrLose[i] = report.Report.NextMovesOrLose
				report.Report.NextMovesOrLose = make([]*inter.Node, 0, 0)
			}
		}
	}
}

func (board Board) createListCheckStone(index int, initialStone inter.Node) ([]*inter.Node, int) {
	listCheckStone := make([]*inter.Node, 0, 11)
	var indexStone int
	axisCheck := axis.DialRightAxes[index]
	Y, X := initialStone.Y+(axisCheck.Y*5), initialStone.X+(axisCheck.X*5)
	for i := 5; i > 0; i-- {
		if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard {

			listCheckStone = append(listCheckStone, &inter.Node{X: X, Y: Y, Player: board[X][Y]})
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
			listCheckStone = append(listCheckStone, &inter.Node{X: X, Y: Y, Player: board[X][Y]})
		}
		Y += axisCheck.Y
		X += axisCheck.X
	}
	return listCheckStone, indexStone
}

func (board Board) proccessRulesByAxes(m func(list []*inter.Node, index int), initialStone inter.Node) {
	for index := 0; index < 4; index++ {
		m(board.createListCheckStone(index, initialStone))
	}
}

// UpdateBoardAfterCapture ...
func (board Board) UpdateBoardAfterCapture(report *rules.Schema) {
	if len := len(report.Report.ListCapturedStone); len != 0 {
		go func(list []*inter.Node, len int) {
			var node *inter.Node
			for len > 0 {
				node = list[len-1]
				board[node.X][node.Y] = 0
				len--
			}
		}(report.Report.ListCapturedStone, len)
	}
}

func trim(X, Y int) (newX, newY int) {
	trimXzero := X - 5
	trimYzero := Y - 5
	trimX := X + 5
	trimY := Y + 5
	if trimXzero < 0 {
		X = 6
	} else if trimX > SizeBoard {
		X = SizeBoard - 6
	}
	if trimYzero < 0 {
		Y = 6
	} else if trimY > SizeBoard {
		Y = SizeBoard - 6
	}
	return X, Y
}

// CreateSearchSpace ...
func (board Board) CreateSearchSpace(lastMoveOpposent inter.Node) []inter.Node {
	X, Y := trim(lastMoveOpposent.X, lastMoveOpposent.Y)
	X, Y = X-6, Y-6
	tmpX, tmpY := 0, 0
	MovesList := make([]inter.Node, 0, 121)
	for x := 0; x < 11; x++ {
		for y := 0; y < 11; y++ {
			tmpX, tmpY = x+X, y+Y
			fmt.Println("TMP +> X", tmpX, " TMP +> Y ", tmpY)
			if board[tmpX][tmpY] == 0 {
				MovesList = append(MovesList, inter.Node{X: tmpX, Y: tmpY})
			}
		}
	}
	return MovesList
}

// func (board Board) UpdateSearchSpace(searhZone []inter.Node, lastMove inter.Node, sizeMax int) {
// 	if searhZone == nil {
// 		searhZone = make([]inter.Node, 0, 361)
// 	} else {
// 		lenSearchZone := len(searhZone)
// 		for i := 0; i < lenSearchZone; i++ {
// 			if searhZone[i].Y == lastMove.Y && searhZone[i].Y == lastMove.Y {
// 				copy(searhZone[i:], searhZone[i+1:])
// 				searhZone[lenSearchZone-1] = inter.Node{}
// 				searhZone = searhZone[:lenSearchZone-1]
// 				break
// 			}
// 		}
// 	}
// 	tmpX, tmpY := 0, 0
// 	for size := sizeMax; size > 0; size++ {
// 		for x := 0; x < sizeMax; x++ {
// 			for y := 0; y < sizeMax; y++ {

// 			}
// 		}
// 	}
// }
