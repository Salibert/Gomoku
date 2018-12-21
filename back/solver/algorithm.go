package solver

import (
	"fmt"
	"math"

	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	"github.com/Salibert/Gomoku/back/server/inter"
)

type Algo struct {
	players  player.Players
	bestMove inter.Node
}

var MAX_DEPTH = 2

func (ia *IA) Play(board board.Board, depth int, players player.Players) *inter.Node {
	ia.minMax.players = players
	var best inter.Node
	if len(ia.SearchZone) != 0 {
		var test int
		// test, best = ia.PVS(board, inter.Node{Player: player.GetOpposentPlayer(ia.playerIndex)}, MAX_DEPTH, -10000, 10000)
		test, best = ia.NegaMax(board, inter.Node{Player: player.GetOpposentPlayer(ia.playerIndex)}, MAX_DEPTH, -10000, 10000)
		// test, best = ia.MinMax(board, inter.Node{Player: player.GetOpposentPlayer(ia.playerIndex)}, MAX_DEPTH, -100000, 1000000, true)
		fmt.Println("TEST +> ", test)
	} else {
		best.X, best.Y, best.Player = 10, 10, 2
	}
	fmt.Println(best)
	fmt.Println(ia.minMax.bestMove)
	fmt.Println(ia.SearchZone)
	best.Player = ia.playerIndex
	ia.minMax.bestMove.Player = ia.playerIndex
	return &best
}

// func (ia *IA) MinMax(board board.Board, move inter.Node, depth int, alpha, beta int, max bool) (current int, best inter.Node) {
// 	if result := ia.isWin(board, depth, move); result != 0 {
// 		return result, move
// 	} else if depth <= 0 {
// 		return ia.HeuristicScore(board, depth, move), move
// 	}
// 	if max == true {
// 		current = math.MinInt64
// 		for _, move = range ia.SearchZone {
// 			if board[move.X][move.Y] == 0 {

// 				board[move.X][move.Y] = ia.playerIndex
// 				move.Player = ia.playerIndex
// 				score, _ := ia.MinMax(board, move, depth-1, alpha, beta, false)
// 				board[move.X][move.Y] = 0
// 				if score > current {
// 					current = score
// 					best = move
// 				}
// 				if score > alpha {
// 					alpha = score
// 					best = move
// 				}
// 				if alpha >= beta {
// 					break
// 				}
// 			}
// 		}
// 	} else {
// 		current = math.MaxInt64
// 		playerIndex := player.GetOpposentPlayer(ia.playerIndex)
// 		for _, move = range ia.SearchZone {
// 			if board[move.X][move.Y] == 0 {

// 				board[move.X][move.Y] = playerIndex
// 				move.Player = playerIndex
// 				score, _ := ia.MinMax(board, move, depth-1, alpha, beta, true)
// 				board[move.X][move.Y] = 0
// 				if score < current {
// 					current = score
// 					best = move
// 				}
// 				if score < beta {
// 					beta = score
// 					best = move
// 				}
// 				if alpha >= beta {
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return
// }

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (ia *IA) NegaMax(board board.Board, move inter.Node, depth, alpha, beta int) (_ int, best inter.Node) {
	if result := ia.isWin(board, depth, move); result != 0 {
		return result, move
	} else if depth <= 0 {
		return ia.HeuristicScore(board, depth, move), move
	}
	var score int
	maximun := math.MinInt64
	playerIndex := player.GetOpposentPlayer(move.Player)
	for _, move = range ia.SearchZone {
		if board[move.X][move.Y] == 0 {
			board[move.X][move.Y] = playerIndex
			move.Player = playerIndex
			score, _ = ia.NegaMax(board, move, depth-1, -beta, -alpha)
			maximun = max(maximun, -score)
			board[move.X][move.Y] = 0
			if maximun > alpha {
				alpha = maximun
				best = move
			}
			if alpha >= beta {
				break
			}
		}
	}
	return alpha, best
}

// var static int

// func (ia *IA) PVS(board board.Board, move inter.Node, depth, alpha, beta int) (_ int, best inter.Node) {
// 	// fmt.Println("MOVE ", move)
// 	if result := ia.isWin(board, depth, move); result != 0 {
// 		return result, move
// 	} else if depth <= 0 {
// 		return ia.HeuristicScore(board, depth, move), move
// 	}
// 	var score, key int
// 	playerIndex := player.GetOpposentPlayer(move.Player)

// 	for key, move = range ia.SearchZone {
// 		if board[move.X][move.Y] == 0 {
// 			board[move.X][move.Y] = playerIndex
// 			move.Player = playerIndex
// 			if key == 0 {
// 				score, _ = ia.PVS(board, move, depth-1, -beta, -alpha)
// 				score = -score
// 			} else {
// 				score, _ = ia.PVS(board, move, depth-1, -(alpha + 1), -alpha)
// 				score = -score
// 				if alpha < score && score < beta {
// 					score, _ = ia.PVS(board, move, depth-1, -beta, -score)
// 					score = -score
// 				}
// 			}
// 			board[move.X][move.Y] = 0
// 			if score > alpha {
// 				alpha = score
// 				best = move
// 			}
// 			// if depth == MAX_DEPTH {
// 			// 	fmt.Println("SCORE ", score, "ALPHA ", alpha, " BETA ", beta, " DEPTH", depth, " MOVE", move, " BEST ", best)
// 			// }
// 			// if depth == MAX_DEPTH-1 {
// 			// 	fmt.Println("SCORE ", score, "ALPHA ", alpha, " BETA ", beta, " DEPTH", depth, " MOVE", move, " BEST ", best)
// 			// }
// 			// if depth == MAX_DEPTH-2 {
// 			// 	fmt.Println("SCORE ", score, "ALPHA ", alpha, " BETA ", beta, " DEPTH", depth, " MOVE", move, " BEST ", best)
// 			// }
// 			if alpha >= beta {
// 				// fmt.Println("ALPHA ", alpha, " BETA ", beta, " DEPTH", depth)
// 				fmt.Println("BREAK")
// 				break
// 			}
// 		}
// 	}
// 	return alpha, best
// }

// func (ia *IA) PVS(board board.Board, move inter.Node, depth, alpha, beta int) (current int, best inter.Node) {
// 	if result := ia.isWin(board, depth, move); result != 0 {
// 		return result, move
// 	} else if depth <= 0 {
// 		return ia.HeuristicScore(board, depth, move), move
// 	}
// 	current, best = ia.PVS(board, move, depth-1, -beta, -alpha)
// 	current = -current
// 	score := 0
// 	playerIndex := player.GetOpposentPlayer(move.Player)
// 	if current >= alpha {
// 		alpha = current
// 		if current < beta {
// 			for _, move = range ia.SearchZone {
// 				if board[move.X][move.Y] == 0 {
// 					board[move.X][move.Y] = playerIndex
// 					move.Player = playerIndex
// 					score, _ = ia.PVS(board, move, depth-1, -(alpha + 1), -alpha)
// 					score = -score
// 					fmt.Println("score > alpha && score < beta ", score, alpha, beta)
// 					if score > alpha && score < beta {
// 						score, _ = ia.PVS(board, move, depth-1, -beta, -alpha)
// 						score = -score
// 						board[move.X][move.Y] = 0
// 						fmt.Println("score >= current ", score, alpha, beta)
// 						if score >= current {
// 							current = score
// 							best = move
// 							ia.minMax.bestMove = move
// 							fmt.Println("score >= alpha ", score, alpha, beta)
// 							if score >= alpha {
// 								alpha = score
// 								best = move
// 								ia.minMax.bestMove = move
// 								if score >= beta {
// 									fmt.Println("BREAK")
// 									break
// 								}
// 							}
// 						}
// 					} else {
// 						board[move.X][move.Y] = 0
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return current, best
// }
