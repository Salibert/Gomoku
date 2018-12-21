package solver

import (
	"math"

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
					tmp += capture * 25
				}
				tmp += report.Report.NbFreeThree
				tmp += report.Report.SizeAlignment * 2
				tmp += report.Report.NbBlockStone * 5
				tmp += report.Report.AmbientScore
				tmp -= report.Report.LevelCapture * 20
				if report.Report.ItIsAValidMove == false {
					return 0
				}
				// if test.X == 9 && test.Y == 8 {
				// 	fmt.Println("REPORT ", tmp)
				// 	fmt.Println(report.Report.NbFreeThree)
				// 	fmt.Println(report.Report.SizeAlignment)
				// 	fmt.Println(report.Report.NbBlockStone)
				// 	fmt.Println(report.Report.AmbientScore)
				// 	fmt.Println(report.Report.LevelCapture)
				// 	fmt.Println("")

				// }
				switch player {
				case ia.playerIndex:
					value += tmp
				default:
					value -= tmp
				}
			}
		}
	}
	// if move.X == 9 && move.Y == 8 {
	// 	fmt.Println(" vaule ", value)
	// }
	// fmt.Println("move => ", move.X, move.Y, move.Player, "detpth ", depth, " VALUE ", value)
	// for i := 0; i < 19; i++ {
	// 	fmt.Println(board[i])
	// }
	// fmt.Println("\n")
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
				return math.MaxInt64 + depth
			default:
				return math.MinInt64 - depth
			}
		}
	} else if ia.minMax.players[move.Player].Score == 8 && len(report.Report.ListCapturedStone) != 0 {
		switch move.Player {
		case ia.playerIndex:
			return math.MaxInt64 + depth
		default:
			return math.MinInt64 - depth
		}
	}
	return 0
}
