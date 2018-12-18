package solver

import (
	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/server/inter"
)

func (ia *IA) Heuristic(board board.Board, move inter.Node) int {
	return ia.HeuristicScore(board, move)
}

// HeuristicScore ...
func (ia *IA) HeuristicScore(board board.Board, move inter.Node) (value int) {
	report := ia.reportEval[move.Player]
	if move.Player == 0 {
		return 0
	}
	defer report.Report.Reset()
	board.CheckRules(move, report)
	if report.Report.ItIsAValidMove == false {
		return 0
	} else {
		if capture := len(report.Report.ListCapturedStone); capture != 0 {
			value += capture * 25
		}
		value += report.Report.NbFreeThree
		value += report.Report.SizeAlignment
		value += report.Report.NbBlockStone
		value -= report.Report.LevelCapture * 10
	}
	switch move.Player {
	case ia.playerIndex:
		return -value
	default:
		return value
	}
}

func (ia *IA) isWin(board board.Board, move inter.Node) int {
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
				return 999
			default:
				return -999
			}
		}
	} else if ia.minMax.players[move.Player].Score == 8 && len(report.Report.ListCapturedStone) != 0 {
		switch move.Player {
		case ia.playerIndex:
			return 999
		default:
			return -999
		}
	}
	return 0
}
