package algorithm

import (
	"fmt"
	"math"

	"github.com/Salibert/Gomoku/back/rules"

	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	"github.com/Salibert/Gomoku/back/server/inter"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

type Algo struct {
	players      player.Players
	bestMove     inter.Node
	reportWin    map[int]rules.Schema
	reportEval   map[int]rules.Schema
	currentMove  inter.Node
	moveOpposent inter.Node
}

// New ...
func New(players player.Players, config pb.ConfigRules, moveOpposent inter.Node) Algo {
	algo := Algo{
		players:    players.Clone(),
		reportWin:  make(map[int]rules.Schema),
		reportEval: make(map[int]rules.Schema),
	}
	config.IsActiveRuleAlignment = true
	config.IsActiveRuleBlock = true
	config.IsActiveRuleProbableCapture = true
	configWin := pb.ConfigRules{
		IsActiveRuleWin:     config.IsActiveRuleWin,
		IsActiveRuleCapture: config.IsActiveRuleCapture,
	}
	algo.reportWin[1] = rules.New(1, 2, configWin)
	algo.reportWin[2] = rules.New(2, 1, configWin)
	algo.reportEval[1] = rules.New(1, 2, config)
	algo.reportEval[2] = rules.New(2, 1, config)
	algo.moveOpposent = moveOpposent
	return algo
}

func IA_jouer(jeu board.Board, profondeur int, players player.Players, config pb.ConfigRules, moveOpposent inter.Node) *inter.Node {
	var max int = -10000
	var tmp int
	var maxi, maxj int
	var best inter.Node
	var i, j int
	algo := New(players, config, moveOpposent)
	alpha, beta := -10000, 10000
	for i = 0; i < board.SizeBoard; i++ {
		for j = 0; j < board.SizeBoard; j++ {
			if jeu[i][j] == 0 {
				jeu[i][j] = 2
				algo.currentMove.X, algo.currentMove.Y, algo.currentMove.Player = i, j, 2
				tmp = algo.alphabeta(jeu, 1, alpha, beta)
				fmt.Println("TMP ", tmp, " MAX ", max)
				if tmp > max {
					max = tmp
					best = algo.bestMove
					maxi = i
					maxj = j
				}
				jeu[i][j] = 0
			}
		}
	}
	res := &inter.Node{X: maxi, Y: maxj, Player: 2}
	fmt.Println("BEST => ", best, " RES => ", res)
	return res
}

func (algo *Algo) alphabeta(jeu board.Board, depth, alpha, beta int) int {
	if algo.gagnant(jeu) != 0 {
		return 10000
	} else if depth <= 0 {
		return algo.newEval(jeu)
	}
	var i, j int
	playerIndex := player.GetOpposentPlayer(algo.currentMove.Player)
	SizeBoard := int(board.SizeBoard)
loop:
	for i = 0; i < SizeBoard; i++ {
		for j = 0; j < SizeBoard; j++ {
			if jeu[i][j] == 0 {
				algo.currentMove.X, algo.currentMove.Y, algo.currentMove.Player = i, j, playerIndex
				jeu[i][j] = playerIndex
				tmp := algo.currentMove
				score := -algo.alphabeta(jeu, depth-1, -beta, -alpha)
				jeu[i][j] = 0
				if score >= alpha {
					alpha = score
					algo.bestMove = tmp
					if alpha >= beta {
						fmt.Println("PRUNE ALPHA")
						break loop
					}
				}
			}
		}
	}
	return alpha
}

// func (algo *Algo) alphabeta(jeu board.Board, depth, alpha, beta int) int {
// 	if algo.gagnant(jeu) != 0 {
// 		return 10000
// 	} else if depth <= 0 {
// 		return algo.newEval(jeu)
// 	}
// 	playerIndex := algo.currentMove.Player
// 	current := algo.alphabeta(jeu, depth-1, -beta, -alpha)
// 	SizeBoard := int(board.SizeBoard)
// 	var i, j int
// 	if current >= alpha {
// 		alpha = current
// 	}
// 	if current < beta {
// 	loop:
// 		for i = 0; i < SizeBoard; i++ {
// 			for j = 0; j < SizeBoard; j++ {
// 				algo.currentMove.X, algo.currentMove.Y, algo.currentMove.Player = i, j, playerIndex
// 				tmp := algo.currentMove
// 				jeu[i][j] = playerIndex
// 				score := -algo.alphabeta(jeu, depth-1, -beta, -alpha)
// 				if score > alpha && score < beta {
// 					score = -algo.alphabeta(jeu, depth-1, -beta, -alpha)
// 				}
// 				jeu[i][j] = 0
// 				if score >= current {
// 					current = score
// 					algo.bestMove = tmp
// 					if score >= alpha {
// 						alpha = score
// 						if score >= beta {
// 							break loop
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return current
// }

// func (algo *Algo) Alpha(jeu board.Board, depth, alpha, beta int) int {
// 	if algo.gagnant(jeu) != 0 {
// 		return 10000
// 	} else if depth <= 0 {
// 		return algo.newEval(jeu)
// 	}
// 	var i, j int
// 	playerIndex := player.GetOpposentPlayer(algo.currentMove.Player)
// 	SizeBoard := int(board.SizeBoard)
// loop:
// 	for i = 0; i < SizeBoard; i++ {
// 		for j = 0; j < SizeBoard; j++ {
// 			if jeu[i][j] == 0 {
// 				algo.currentMove.X, algo.currentMove.Y, algo.currentMove.Player = i, j, playerIndex
// 				jeu[i][j] = playerIndex
// 				tmp := algo.currentMove
// 				score := algo.Beta(jeu, depth-1, alpha, beta)
// 				jeu[i][j] = 0
// 				if score > alpha {
// 					alpha = score
// 					algo.bestMove = tmp
// 					if alpha >= beta {
// 						fmt.Println("PRUNE ALPHA")
// 						break loop
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return alpha
// }

// func (algo *Algo) Beta(jeu board.Board, depth, alpha, beta int) int {
// 	if algo.gagnant(jeu) != 0 {
// 		return 10000
// 	} else if depth <= 0 {
// 		return algo.newEval(jeu)
// 	}
// 	var i, j int
// 	playerIndex := player.GetOpposentPlayer(algo.currentMove.Player)
// 	SizeBoard := int(board.SizeBoard)

// loop:
// 	for i = 0; i < SizeBoard; i++ {
// 		for j = 0; j < SizeBoard; j++ {
// 			if jeu[i][j] == 0 {
// 				algo.currentMove.X, algo.currentMove.Y, algo.currentMove.Player = i, j, playerIndex
// 				jeu[i][j] = playerIndex
// 				tmp := algo.currentMove
// 				score := algo.Alpha(jeu, depth-1, alpha, beta)
// 				jeu[i][j] = 0
// 				if score < algo.beta {
// 					beta = score
// 					algo.bestMove = tmp
// 					if alpha >= beta {
// 						fmt.Println("PRUNE BETA")
// 						break loop
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return beta
// }

// // MTDF ...
// func MTDF(jeu board.Board, players player.Players, config pb.ConfigRules, moveOpposent inter.Node, depth, initG int) int {
// 	g := initG
// 	algo := New(players, config, moveOpposent)
// 	var beta int
// 	upperbound := math.MaxInt64
// 	lowerbound := math.MinInt64
// 	for {
// 		if g == lowerbound {
// 			beta = g + 1
// 		} else {
// 			beta = g
// 		}
// 		g = algo.AlphaBetaWithMemory(depth, beta-1, beta)
// 		if g < beta {
// 			upperbound = g
// 		} else {
// 			lowerbound = g
// 		}
// 	}
// 	return g
// }

// // AlphaBetaWithMemory ...
// func (algo Algo) AlphaBetaWithMemory(depth, alpha, beta int) int {
// 	if
// }

func distance(moveOpposent, currentMove inter.Node) int {
	return int(math.Sqrt(math.Pow(float64(moveOpposent.X-currentMove.X), 2) +
		math.Pow(float64(moveOpposent.Y-currentMove.Y), 2)))
}

func (algo *Algo) newEval(jeu board.Board) int {
	value := 0
	raport := algo.reportEval[algo.currentMove.Player]
	defer raport.Report.Reset()
	jeu.CheckRules(algo.currentMove, raport)
	if raport.Report.ItIsAValidMove == false {
		return -10000
	} else {
		if len(algo.players[algo.currentMove.Player].NextMovesOrLose) != 0 {
			if raport.Report.PartyFinish = algo.players[algo.currentMove.Player].CheckIfThisMoveBlockLose(&algo.currentMove); raport.Report.PartyFinish == true {
				return 10000
			}
		}
		if capture := len(raport.Report.ListCapturedStone); capture != 0 {
			value += capture * 100
		}
		value += raport.Report.NbFreeThree * 15
		if raport.Report.PartyFinish == true {
			if counter := len(raport.Report.WinOrLose); counter == 0 {
				value += 100
			}
		}
		value += raport.Report.SizeAlignment * 5
		value += raport.Report.NbBlockStone * 5
		value -= raport.Report.LevelCapture * 10
	}
	value += 19 - distance(algo.moveOpposent, algo.currentMove)
	// if value != 0 {
	func() {
		fmt.Println("NEW EVAL:",
			"\n X ", algo.currentMove.X,
			"\n Y ", algo.currentMove.Y,
			"\n Player ", algo.currentMove.Player,
			"\n REPORT ", raport.Report,
			"\nSCORE +> ", value,
			"\n", jeu[0],
			"\n", jeu[1],
			"\n", jeu[2],
			"\n", jeu[3],
			"\n", jeu[4],
			"\n", jeu[5],
			"\n", jeu[6],
			"\n", jeu[7],
			"\n", jeu[8],
			"\n", jeu[9],
			"\n", jeu[10],
			"\n", jeu[11],
			"\n", jeu[12],
			"\n", jeu[13],
			"\n", jeu[14],
			"\n", jeu[15],
			"\n", jeu[16],
			"\n", jeu[17],
			"\n", jeu[18],
			"\n")
	}()
	// }
	return value
}

func (algo *Algo) gagnant(jeu board.Board) int {
	raport := algo.reportWin[algo.currentMove.Player]
	defer raport.Report.Reset()
	jeu.CheckRules(algo.currentMove, raport)
	if raport.Report.PartyFinish == true {
		if len(raport.Report.WinOrLose[0]) == 0 {
			return 1
		}
	} else if algo.players[algo.currentMove.Player].Score == 8 && len(raport.Report.ListCapturedStone) != 0 {
		return 1
	}
	return 0
}
