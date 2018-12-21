package solver

import (
	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/server/inter"
)

// HeuristicScore ...
func (ia *IA) HeuristicScore(board board.Board, depth int, move inter.Node) (value int) {
	report := ia.reportEval[move.Player]
	defer report.Report.Reset()
	if move.Player == 0 {
		return 0
	}
	var tmp, player int
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			player = board[i][j]
			if player != 0 {
				tmp = 0
				// fmt.Println("BEFORE CLEAN ", report.Report)
				report.Report.Reset()
				// fmt.Println("AFTER CLEAN ", report.Report)
				test := inter.Node{X: i, Y: j, Player: player}
				board.CheckRules(test, report)
				if capture := len(report.Report.ListCapturedStone); capture != 0 {
					tmp += capture * 30
				}
				tmp += report.Report.NbFreeThree * 5
				tmp += report.Report.SizeAlignment * 5
				tmp += report.Report.NbBlockStone * 5
				tmp += report.Report.AmbientScore
				tmp -= report.Report.LevelCapture * 100
				if report.Report.ItIsAValidMove == false {
					return 0
				}
				switch player {
				case ia.playerIndex:
					value += tmp
				default:
					value -= tmp
				}
			}
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
	} else if ia.minMax.players[move.Player].Score == 8 && len(report.Report.ListCapturedStone) != 0 {
		switch move.Player {
		case ia.playerIndex:
			return 10000 + depth
		default:
			return -10000 - depth
		}
	}
	return 0
}
