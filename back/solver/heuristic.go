package solver

import (
	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/server/inter"
)

// HeuristicScore ...
func (ia *IA) HeuristicScore(board board.Board, list [sizeListMoves]inter.Node, index int, move inter.Node) (value int) {
	report := ia.reportEval[move.Player]
	defer func() {
		report.Report.Reset()
	}()
	if move.Player == 0 {
		return 0
	}
	var tmp int
	for i := 0; i < index; i++ {
		tmp = 0
		report.Report.Reset()
		node := list[i]
		board.CheckRules(node, report)
		if capture := len(report.Report.ListCapturedStone); capture != 0 {
			tmp += capture * 50
		}
		tmp += report.Report.NbFreeThree * 5
		tmp += report.Report.SizeAlignment * 7
		tmp += report.Report.NbBlockStone * 7
		tmp += report.Report.AmbientScore
		tmp -= report.Report.LevelCapture * 110
		if report.Report.ItIsAValidMove == false {
			return 0
		}
		switch node.Player {
		case ia.playerIndex:
			value += tmp
		default:
			value -= tmp
		}
	}
	return value
}

func (ia *IA) isWin(board board.Board, depth int, move inter.Node) int {
	report := ia.reportWin[move.Player]
	if move.Player == 0 {
		return 0
	}
	defer func() {
		report.Report.Reset()
	}()
	board.CheckRules(move, report)
	if report.Report.PartyFinish == true {
		if len(report.Report.WinOrLose[0]) == 0 {
			switch move.Player {
			case ia.playerIndex:
				return 10000 + depth
			default:
				return -10000 - depth
			}
		}
	} else if ia.playersScore[move.Player-1] == 8 && len(report.Report.ListCapturedStone) != 0 {
		switch move.Player {
		case ia.playerIndex:
			return 10000 + depth
		default:
			return -10000 - depth
		}
	}
	return 0
}
