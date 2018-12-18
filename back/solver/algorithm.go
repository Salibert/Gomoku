package solver

import (
	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	"github.com/Salibert/Gomoku/back/server/inter"
)

type Algo struct {
	players  player.Players
	bestMove inter.Node
}

func (ia *IA) Play(board board.Board, depth int, players player.Players) *inter.Node {
	ia.minMax.players = players
	// fmt.Println("SearchZone => ", ia.SearchZone, " LEN +> ", len(ia.SearchZone))
	alpha, beta := -10000, 10000
	var best inter.Node
	if len(ia.SearchZone) != 0 {
		var current, score int
		move := ia.SearchZone[0]
		playerIndex := ia.playerIndex
		move.Player = playerIndex
		current, best = ia.PVS(board, move, depth-1, alpha, beta)
		// var tmpMove inter.Node
		for _, move := range ia.SearchZone {
			if board[move.X][move.Y] == 0 {
				board[move.X][move.Y] = playerIndex
				move.Player = playerIndex
				score, _ = ia.PVS(board, move, depth-1, -(alpha + 1), -alpha)
				score = -score
				if score > alpha && score < beta {
					score, _ = ia.PVS(board, move, depth-1, -beta, -alpha)
					score = -score
					board[move.X][move.Y] = 0
					if score >= current {
						current = score
						best = move
						if score >= alpha {
							alpha = score
							if score >= beta {
								break
							}
						}
					}
				} else {
					board[move.X][move.Y] = 0
				}
			}
		}
		_, best = ia.PVS(board, inter.Node{}, depth, alpha, beta)
	} else {
		ia.minMax.bestMove.X, ia.minMax.bestMove.Y, ia.minMax.bestMove.Player = 10, 10, 2
	}
	best.Player = ia.playerIndex
	return &best
}

func (ia *IA) PVS(board board.Board, move inter.Node, depth, alpha, beta int) (current int, best inter.Node) {
	if result := ia.isWin(board, move); result != 0 {
		return result, move
	} else if depth <= 0 {
		return ia.Heuristic(board, move), move
	}
	current, best = ia.PVS(board, move, depth-1, -beta, -alpha)
	current = -current
	score := 0
	// var tmpMove inter.Node
	playerIndex := player.GetOpposentPlayer(move.Player)
	if current >= alpha {
		alpha = current
		if current < beta {
			for _, move = range ia.SearchZone {
				if board[move.X][move.Y] == 0 {
					board[move.X][move.Y] = playerIndex
					move.Player = playerIndex
					score, _ = ia.PVS(board, move, depth-1, -(alpha + 1), -alpha)
					score = -score
					if score > alpha && score < beta {
						score, _ = ia.PVS(board, move, depth-1, -beta, -alpha)
						score = -score
						board[move.X][move.Y] = 0
						if score >= current {
							current = score
							best = move
							if score >= alpha {
								alpha = score
								if score >= beta {
									break
								}
							}
						}
					} else {
						board[move.X][move.Y] = 0
					}
				}
			}
		}
	}
	return current, best
}

// func (ia *IA) Alpha(board board.Board, move inter.Node, depth, alpha, beta int) int {
// 	if ia.isWin(board, move) != 0 {
// 		return 999
// 	} else if depth <= 0 {
// 		return ia.HeuristicScore(board, move)
// 	}
// 	playerIndex := player.GetOpposentPlayer(move.Player)
// loop:
// 	for _, move = range ia.SearchZone {
// 		if board[move.X][move.Y] == 0 {
// 			board[move.X][move.Y] = playerIndex
// 			move.Player = playerIndex
// 			score := ia.Beta(board, move, depth-1, alpha, beta)
// 			board[move.X][move.Y] = 0
// 			if score > alpha {
// 				alpha = score
// 				ia.minMax.bestMove = move
// 				if alpha >= beta {
// 					break loop
// 				}
// 			}

// 		}
// 	}
// 	return alpha
// }

// func (ia *IA) Beta(board board.Board, move inter.Node, depth, alpha, beta int) int {
// 	if ia.isWin(board, move) != 0 {
// 		return 999
// 	} else if depth <= 0 {
// 		return ia.HeuristicScore(board, move)
// 	}
// 	playerIndex := player.GetOpposentPlayer(move.Player)
// loop:
// 	for _, move = range ia.SearchZone {
// 		if board[move.X][move.Y] == 0 {
// 			move.Player = playerIndex
// 			board[move.X][move.Y] = playerIndex
// 			score := ia.Alpha(board, move, depth-1, alpha, beta)
// 			board[move.X][move.Y] = 0
// 			if score < beta {
// 				beta = score
// 				ia.minMax.bestMove = move
// 				if alpha >= beta {
// 					break loop
// 				}
// 			}
// 		}
// 	}
// 	return beta
// }
