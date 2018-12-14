package solver

import (
	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/server/inter"
)

// HeuristicScore ...
func (ia *IA) HeuristicScore(board board.Board, move inter.Node) int {
	value := 0
	report := ia.reportEval[move.Player]
	defer report.Report.Reset()
	board.CheckRules(move, report)
	if report.Report.ItIsAValidMove == false {
		return -1000
	} else {
		if len(ia.minMax.players[move.Player].NextMovesOrLose) != 0 {
			if report.Report.PartyFinish = ia.minMax.players[move.Player].CheckIfThisMoveBlockLose(&move); report.Report.PartyFinish == true {
				return 1000
			}
		}
		if capture := len(report.Report.ListCapturedStone); capture != 0 {
			value += capture * 25
		}
		value += report.Report.NbFreeThree * 5
		if report.Report.PartyFinish == true {
			if counter := len(report.Report.WinOrLose); counter == 0 {
				value += 100
			}
		}
		value += report.Report.SizeAlignment
		value += report.Report.NbBlockStone
		value -= report.Report.LevelCapture * 2
	}
	return value
}

func (ia *IA) isWin(board board.Board, move inter.Node) int {
	report := ia.reportWin[move.Player]
	defer report.Report.Reset()
	board.CheckRules(move, report)
	if report.Report.PartyFinish == true {
		if len(report.Report.WinOrLose[0]) == 0 {
			return 1
		}
	} else if ia.minMax.players[move.Player].Score == 8 && len(report.Report.ListCapturedStone) != 0 {
		return 1
	}
	return 0
}
