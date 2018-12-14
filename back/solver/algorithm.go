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
	if len(ia.SearchZone) != 0 {
		alpha, beta := -10000, 10000
		for _, move := range ia.SearchZone {
			if board[move.X][move.Y] == 0 {
				// move := ia.SearchZone[0]
				board[move.X][move.Y] = 2
				move.Player = 2
				_ = ia.PVS(board, move, 2, alpha, beta)
				board[move.X][move.Y] = 0
			}
		}
	} else {
		ia.minMax.bestMove.X, ia.minMax.bestMove.Y, ia.minMax.bestMove.Player = 8, 8, 2
	}

	ia.minMax.bestMove.Player = ia.playerIndex
	return &ia.minMax.bestMove
}

func (ia *IA) PVS(board board.Board, move inter.Node, depth, alpha, beta int) int {
	if ia.isWin(board, move) != 0 {
		return 999
	} else if depth <= 0 {
		return ia.HeuristicScore(board, move)
	}
	current := -ia.PVS(board, move, depth-1, -beta, -alpha)
	playerIndex := player.GetOpposentPlayer(move.Player)
	if current >= alpha {
		alpha = current
		if current < beta {
			for _, move = range ia.SearchZone {
				if board[move.X][move.Y] == 0 {
					board[move.X][move.Y] = playerIndex
					move.Player = playerIndex
					score := -ia.PVS(board, move, depth-1, -(alpha + 1), -alpha)
					if score > alpha && score < beta {
						score = -ia.PVS(board, move, depth-1, -beta, -alpha)
						board[move.X][move.Y] = 0
						if score >= current {
							current = score
							ia.minMax.bestMove = move
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
	return current
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
