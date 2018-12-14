package solver

import (
	"fmt"

	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	"github.com/Salibert/Gomoku/back/server/inter"
)

type Algo struct {
	players  player.Players
	bestMove inter.Node
}

func (ia *IA) Play(board board.Board, depth int, players player.Players) *inter.Node {
	fmt.Println("LEN +> ", len(ia.searchZone))
	ia.minMax.players = players
	alpha, beta := -10000, 10000
	for _, move := range ia.searchZone {
		if board[move.X][move.Y] == 0 {
			board[move.X][move.Y] = 2
			move.Player = 2
			_ = ia.Beta(board, move, 2, alpha, beta)
			board[move.X][move.Y] = 0
		}

	}
	ia.minMax.bestMove.Player = ia.playerIndex
	return &ia.minMax.bestMove
}

func (ia *IA) Alpha(board board.Board, move inter.Node, depth, alpha, beta int) int {
	if ia.isWin(board, move) != 0 {
		return 999
	} else if depth <= 0 {
		return ia.HeuristicScore(board, move)
	}
	playerIndex := player.GetOpposentPlayer(move.Player)
loop:
	for _, move = range ia.searchZone {
		if board[move.X][move.Y] == 0 {
			board[move.X][move.Y] = playerIndex
			move.Player = playerIndex
			score := ia.Beta(board, move, depth-1, alpha, beta)
			board[move.X][move.Y] = 0
			if score > alpha {
				alpha = score
				ia.minMax.bestMove = move
				if alpha >= beta {
					break loop
				}
			}

		}
	}
	return alpha
}

func (ia *IA) Beta(board board.Board, move inter.Node, depth, alpha, beta int) int {
	if ia.isWin(board, move) != 0 {
		return 999
	} else if depth <= 0 {
		return ia.HeuristicScore(board, move)
	}
	playerIndex := player.GetOpposentPlayer(move.Player)
loop:
	for _, move = range ia.searchZone {
		if board[move.X][move.Y] == 0 {
			move.Player = playerIndex
			board[move.X][move.Y] = playerIndex
			score := ia.Alpha(board, move, depth-1, alpha, beta)
			board[move.X][move.Y] = 0
			if score < beta {
				beta = score
				ia.minMax.bestMove = move
				if alpha >= beta {
					break loop
				}
			}
		}
	}
	return beta
}
