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

const (
	sizeListMoves int = 361
	sizeMap       int = 19
)

type ABoard [sizeMap][sizeMap]int

func createListMoves(NewListMoves *[sizeListMoves]inter.Node, listMoves []inter.Node) {
	for key, el := range listMoves {
		NewListMoves[key] = el
	}
}

func createABoard(newBoard *ABoard, board board.Board) {
	for x, el := range board {
		for y, point := range el {
			newBoard[x][y] = point
		}
	}
}

func createBoard(aBoard *ABoard, board board.Board) {
	for x, el := range aBoard {
		for y, point := range el {
			board[x][y] = point
		}
	}
}

func (ia *IA) Play(board board.Board, players player.Players) *inter.Node {
	ia.players = players
	var best inter.Node
	var listMoves [sizeListMoves]inter.Node
	var newBoard ABoard
	createListMoves(&listMoves, ia.ListMoves)
	createABoard(&newBoard, board)
	if len(ia.SearchZone) != 0 {
		start := time.Now()
		_, best = ia.MinMax(newBoard, listMoves, inter.Node{Player: player.GetOpposentPlayer(ia.playerIndex)}, ia.depth, -100000, 1000000, len(ia.ListMoves), true)
		t := time.Now()
		fmt.Println(t.Sub(start))
	} else {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		best.X = r.Intn(13) + 3
		best.Y = r.Intn(13) + 3
		best.Player = ia.playerIndex
	}
	best.Player = ia.playerIndex
	return &best
}

func (ia *IA) MinMax(board ABoard, list [sizeListMoves]inter.Node, move inter.Node, depth, alpha, beta, index int, max bool) (current int, best inter.Node) {
	if result := ia.isWin(&board, depth, move); result != 0 {
		return result, move
	} else if depth <= 0 {
		return ia.HeuristicScore(&board, list, index, move), move
	}
	if max == true {
		current = math.MinInt64
		for _, move = range ia.SearchZone {
			if board[move.X][move.Y] == 0 {

				board[move.X][move.Y] = ia.playerIndex
				move.Player = ia.playerIndex
				list[index] = move
				score, _ := ia.MinMax(board, list, move, depth-1, alpha, beta, index+1, false)
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
				list[index] = move
				score, _ := ia.MinMax(board, list, move, depth-1, alpha, beta, index+1, true)
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
