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

func (ia *IA) Play(board *board.Board, players player.Players) *inter.Node {
	ia.playersScore[0] = players[1].Score
	ia.playersScore[1] = players[2].Score
	var best inter.Node
	var listMoves [sizeListMoves]inter.Node
	createListMoves(&listMoves, ia.ListMoves)
	if len(ia.SearchZone) != 0 {
		start := time.Now()
		_, best = ia.MinMax(*board, listMoves, inter.Node{Player: player.GetOpposentPlayer(ia.PlayerIndex)}, ia.Depth, -100000, 1000000, len(ia.ListMoves), true)
		t := time.Now()
		fmt.Println(t.Sub(start))
	} else {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		best.X = r.Intn(13) + 3
		best.Y = r.Intn(13) + 3
		best.Player = ia.PlayerIndex
	}
	best.Player = ia.PlayerIndex
	return &best
}

func (ia *IA) MinMax(board board.Board, list [sizeListMoves]inter.Node, move inter.Node, depth, alpha, beta, index int, max bool) (current int, best inter.Node) {
	if result := ia.isWin(board, depth, move); result != 0 {
		return result, move
	} else if depth <= 0 {
		return ia.HeuristicScore(board, list, index, move), move
	}
	if max == true {
		current = math.MinInt64
		for _, move = range ia.SearchZone {
			if board[move.X][move.Y] == 0 {

				board[move.X][move.Y] = ia.PlayerIndex
				move.Player = ia.PlayerIndex
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
		playerIndex := player.GetOpposentPlayer(ia.PlayerIndex)
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
