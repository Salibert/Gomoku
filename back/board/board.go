package board

import (
	"github.com/Salibert/Gomoku/back/axis"
	"github.com/Salibert/Gomoku/back/rules"
	"github.com/Salibert/Gomoku/back/server/inter"
)

// SizeBoard is the size of the board
const SizeBoard int = 19

// Board ...
type Board [SizeBoard][SizeBoard]int

// UpdateBoard board with new stone
func (board *Board) UpdateBoard(stone inter.Node) {
	board[stone.X][stone.Y] = stone.Player
}

// CheckRules check all rules. Modify the report passed in params
func (board *Board) CheckRules(initialStone inter.Node, report rules.Schema) {
	board.proccessRulesByAxes(report.ProccessCheckRules, initialStone)
	if report.Report.NbFreeThree > 1 {
		report.Report.ItIsAValidMove = false
		return
	}
	report.Report.ItIsAValidMove = true
	if lenWinOrLose := len(report.Report.WinOrLose); lenWinOrLose != 0 {
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
			report.Report.NextMovesOrLose = make([]*inter.Node, 0, 10)
		}
	}
}

func (board *Board) sendRadius(m func(list *axis.Radius, index, lenRadius int), index int, initialStone inter.Node) {
	listCheckStone := &axis.Radius{}
	var indexStone, indexRadius int
	axisCheck := axis.DialRightAxes[index]
	Y, X := initialStone.Y+(axisCheck.Y*5), initialStone.X+(axisCheck.X*5)
	for i := 5; i > 0; i-- {
		if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard {
			listCheckStone[indexRadius] = inter.Node{X: X, Y: Y, Player: board[X][Y]}
			indexRadius++
		}
		Y -= axisCheck.Y
		X -= axisCheck.X
	}
	indexStone = indexRadius
	indexRadius++
	listCheckStone[indexStone] = initialStone
	axisCheck = axisCheck.Inverse()
	Y, X = initialStone.Y+axisCheck.Y, initialStone.X+axisCheck.X
	for i := 0; i < 5; i++ {
		if Y >= 0 && X >= 0 && Y < SizeBoard && X < SizeBoard {
			listCheckStone[indexRadius] = inter.Node{X: X, Y: Y, Player: board[X][Y]}
			indexRadius++
		}
		Y += axisCheck.Y
		X += axisCheck.X
	}
	m(listCheckStone, indexStone, indexRadius)
}

func (board *Board) proccessRulesByAxes(m func(list *axis.Radius, index, lenRadius int), initialStone inter.Node) {
	for index := 0; index < 4; index++ {
		board.sendRadius(m, index, initialStone)
	}
}

// UpdateBoardAfterCapture ...
func (board *Board) UpdateBoardAfterCapture(report *rules.Schema) {
	if len := len(report.Report.ListCapturedStone); len != 0 {
		func(list []*inter.Node, len int) {
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

func searchIfZoneExist(searchZone []inter.Node, x, y int) bool {
	for _, el := range searchZone {
		if el.X == x && el.Y == y {
			return true
		}
	}
	return false
}

func (board *Board) UpdateSearchSpace(searchZone *[]inter.Node, lastMove inter.Node, size int) {
	lenSearchZone := len(*searchZone)
	for i := 0; i < lenSearchZone; i++ {
		if (*searchZone)[i].Y == lastMove.Y && (*searchZone)[i].Y == lastMove.Y {
			copy((*searchZone)[i:], (*searchZone)[i+1:])
			(*searchZone)[lenSearchZone-1] = inter.Node{}
			(*searchZone) = (*searchZone)[:lenSearchZone-1]
			break
		}
	}
	tmpX, tmpY := 0, 0
	sizeMax := size + 1
	var isExist bool
	for x := -size; x < sizeMax; x++ {
		for y := -size; y < sizeMax; y++ {
			tmpX, tmpY = x+lastMove.X, y+lastMove.Y
			if tmpX >= 0 && tmpX < SizeBoard && tmpY >= 0 && tmpY < SizeBoard {
				if isExist = searchIfZoneExist(*searchZone, tmpX, tmpY); isExist == false {
					if board[tmpX][tmpY] == 0 {
						*searchZone = append(*searchZone, inter.Node{X: tmpX, Y: tmpY})
					}
				}
			}
		}
	}
}
