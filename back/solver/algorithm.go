package solver

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	"github.com/Salibert/Gomoku/back/server/inter"
)

type Algo struct {
	players  player.Players
	bestMove inter.Node
}

func (ia *IA) Play(board board.Board, players player.Players) *inter.Node {
	ia.minMax.players = players
	var best inter.Node
	if len(ia.SearchZone) != 0 {
		_, best = ia.MinMax(board, inter.Node{Player: player.GetOpposentPlayer(ia.playerIndex)}, ia.depth, -100000, 1000000, true)
	} else {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		tmp1 := r.Intn(13)
		best.X = tmp1 + 3
		tmp2 := r.Intn(13)
		best.Y = tmp2 + 3
		best.Player = ia.playerIndex
		fmt.Println("TMP +> 1: ", tmp1, " 2:", tmp2)
	}
	best.Player = ia.playerIndex
	return &best
}

func (ia *IA) MinMax(board board.Board, move inter.Node, depth int, alpha, beta int, max bool) (current int, best inter.Node) {
	if result := ia.isWin(board, depth, move); result != 0 {
		return result, move
	} else if depth <= 0 {
		return ia.HeuristicScore(board, depth, move), move
	}
	if max == true {
		current = math.MinInt64
		for _, move = range ia.SearchZone {
			if board[move.X][move.Y] == 0 {

				board[move.X][move.Y] = ia.playerIndex
				move.Player = ia.playerIndex
				score, _ := ia.MinMax(board, move, depth-1, alpha, beta, false)
				board[move.X][move.Y] = 0
				if score > current {
					current = score
					best = move
				}
				if score > alpha {
					alpha = score
					best = move
				}
				if alpha >= beta {
					break
				}
			}
		}
	} else {
		current = math.MaxInt64
		playerIndex := player.GetOpposentPlayer(ia.playerIndex)
		for _, move = range ia.SearchZone {
			if board[move.X][move.Y] == 0 {
				board[move.X][move.Y] = playerIndex
				move.Player = playerIndex
				score, _ := ia.MinMax(board, move, depth-1, alpha, beta, true)
				board[move.X][move.Y] = 0
				if score < current {
					current = score
					best = move
				}
				if score < beta {
					beta = score
					best = move
				}
				if alpha >= beta {
					break
				}
			}
		}
	}
	return
}
