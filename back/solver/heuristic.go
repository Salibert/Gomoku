package solver

import (
	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/server/inter"
)

// HeuristicScore ...
func (ia *IA) HeuristicScore(board board.Board, list Tlist, index int, move inter.Node) (value int) {
	report := ia.reportEval[move.Player]
	if move.Player == 0 {
		return 0
	}
	var tmp int
	for i := 0; i < index; i++ {
		tmp = 0
		node := list[i]
		board.CheckRules(node, &report)
		tmp += report.Report.IndexListCapturedStone * 50
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
	board.CheckRules(move, &report)
	if report.Report.PartyFinish == true {
		var value int
		for _, el := range report.Report.WinOrLose {
			if len(el) == 0 {
				switch move.Player {
				case ia.playerIndex:
					value += 10000 + depth
				default:
					value += -10000 - depth
				}
			}
		}
		return value
	} else if ia.playersScore[move.Player-1] == 8 && report.Report.IndexListCapturedStone != 0 {
		switch move.Player {
		case ia.playerIndex:
			return 10000 + depth
		default:
			return -10000 - depth
		}
	}
	return 0
}
