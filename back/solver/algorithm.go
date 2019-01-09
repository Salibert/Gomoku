package solver

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	"github.com/Salibert/Gomoku/back/server/inter"
)

const (
	sizeListMoves int = 361
	sizeMap       int = 19
)

type Tlist [361]inter.Node
type IASelfService struct {
	billy IA
	rwmux sync.RWMutex
}

var ici int

func createListMoves(NewListMoves *Tlist, listMoves []inter.Node) {
	for key, el := range listMoves {
		NewListMoves[key] = el
	}
}

func createSearchList(NewListMoves *Tlist, listMoves []inter.Node) {
	for key, el := range listMoves {
		NewListMoves[key] = el
	}
}

func (ia *IA) Play(board *board.Board, players player.Players) *inter.Node {
	ia.playersScore[0] = players[1].Score
	ia.playersScore[1] = players[2].Score
	var best inter.Node
	var listMoves, searchList Tlist
	createListMoves(&listMoves, ia.ListMoves)
	createSearchList(&searchList, ia.SearchZone)
	if len(ia.SearchZone) != 0 {
		start := time.Now()
		_, best = ia.Boost(*board, listMoves, searchList, len(ia.ListMoves), len(ia.SearchZone))
		t := time.Now()
		fmt.Println("TIME +> ", t.Sub(start))
	} else {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		best.X = r.Intn(13) + 3
		best.Y = r.Intn(13) + 3
		best.Player = ia.PlayerIndex
	}
	best.Player = ia.PlayerIndex
	return &best
}

func (ia *IA) Boost(board board.Board, list, searchList Tlist, indexListMoves, indexSearchList int) (current int, best inter.Node) {
	current = math.MinInt64
	alpha, beta := -100000, 1000000
	move := inter.Node{Player: player.GetOpposentPlayer(ia.PlayerIndex)}
	if ia.Depth > 3 && indexSearchList > 50 {
		for i := 0; i < indexSearchList; i++ {
			move = searchList[i]
			if board[move.X][move.Y] == 0 {

				board[move.X][move.Y] = ia.PlayerIndex
				move.Player = ia.PlayerIndex
				list[indexListMoves] = move
				score, _ := ia.GMin(board, list, searchList, move, ia.Depth-1, alpha, beta, indexListMoves+1, indexSearchList)
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
				break
			}
		}
		chanScores := make(chan int, indexSearchList)
		for i := 1; i < indexSearchList; i++ {
			if i < indexSearchList {
				move = searchList[i]
				if board[move.X][move.Y] != 0 {
					i++
					continue
				}
				board[move.X][move.Y] = ia.PlayerIndex
				move.Player = ia.PlayerIndex
				list[indexListMoves] = move
				newIA := ia.Pool.Get().(*IA)
				go func(score int) {
					score, _ = newIA.GMin(board, list, searchList, move, ia.Depth-1, alpha, beta, indexListMoves+1, indexSearchList)
					chanScores <- score
					ia.Pool.Put(newIA)
				}(0)
				board[move.X][move.Y] = 0
			}
		}
	loop:
		for true {
			select {
			case score := <-chanScores:
				if score > current {
					current = score
					best = move
				}
				if score > alpha {
					alpha = score
					best = move
				}
				if alpha >= beta {
					break loop
				}
				if len(chanScores) == 0 {
					break loop
				}
			default:
			}
		}
	} else {
		for i := 0; i < indexSearchList; i++ {
			move = searchList[i]
			if board[move.X][move.Y] == 0 {

				board[move.X][move.Y] = ia.PlayerIndex
				move.Player = ia.PlayerIndex
				list[indexListMoves] = move
				score, _ := ia.Min(board, list, searchList, move, ia.Depth-1, alpha, beta, indexListMoves+1, indexSearchList)
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
	}
	return
}

func (ia *IA) GMax(board board.Board, list, searchList Tlist, move inter.Node, depth, alpha, beta, indexListMoves, indexSearchList int) (current int, best inter.Node) {
	if result := ia.isWin(board, depth, move); result != 0 {
		return result, move
	} else if depth <= 0 {
		return ia.HeuristicScore(board, list, indexListMoves, move), move
	}
	current = math.MinInt64
	for i := 0; i < indexSearchList; i++ {
		move = searchList[i]
		if board[move.X][move.Y] == 0 {

			board[move.X][move.Y] = ia.PlayerIndex
			move.Player = ia.PlayerIndex
			list[indexListMoves] = move
			score, _ := ia.GMin(board, list, searchList, move, depth-1, alpha, beta, indexListMoves+1, indexSearchList)
			board[move.X][move.Y] = 0
			if score > current {
				current = score
				best = move
				if score > alpha {
					alpha = score
					best = move
				}
				if alpha >= beta {
					break
				}
			}
			break
		}
	}
	chanScores := make(chan int, indexSearchList)
	for i := 1; i < indexSearchList; i++ {
		if i < indexSearchList {
			move = searchList[i]
			if board[move.X][move.Y] != 0 {
				i++
				continue
			}
			board[move.X][move.Y] = ia.PlayerIndex
			move.Player = ia.PlayerIndex
			list[indexListMoves] = move
			newIA := ia.Pool.Get().(*IA)
			go func(score int) {
				score, _ = newIA.GMin(board, list, searchList, move, ia.Depth-1, alpha, beta, indexListMoves+1, indexSearchList)
				chanScores <- score
				ia.Pool.Put(newIA)
			}(0)
			board[move.X][move.Y] = 0
		}
	}
loop:
	for true {
		select {
		case score := <-chanScores:
			if score > current {
				current = score
				best = move
			}
			if score > alpha {
				alpha = score
				best = move
			}
			if alpha >= beta {
				break loop
			}
			if len(chanScores) == 0 {
				break loop
			}
		default:
		}
	}
	return
}

func (ia *IA) GMin(board board.Board, list, searchList Tlist, move inter.Node, depth, alpha, beta, indexListMoves, indexSearchList int) (current int, best inter.Node) {
	if result := ia.isWin(board, depth, move); result != 0 {
		return result, move
	} else if depth <= 0 {
		return ia.HeuristicScore(board, list, indexListMoves, move), move
	}
	current = math.MaxInt64
	playerIndex := player.GetOpposentPlayer(ia.PlayerIndex)
	for i := 0; i < indexSearchList; i++ {
		move = searchList[i]
		if board[move.X][move.Y] == 0 {
			board[move.X][move.Y] = playerIndex
			move.Player = playerIndex
			list[indexListMoves] = move
			score, _ := ia.GMax(board, list, searchList, move, depth-1, alpha, beta, indexListMoves+1, indexSearchList)
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
			break
		}
	}
	chanScores := make(chan int, indexSearchList)
	for i := 1; i < indexSearchList; i++ {
		if i < indexSearchList {
			move = searchList[i]
			if board[move.X][move.Y] != 0 {
				i++
				continue
			}
			board[move.X][move.Y] = ia.PlayerIndex
			move.Player = ia.PlayerIndex
			list[indexListMoves] = move
			newIA := ia.Pool.Get().(*IA)
			go func(score int) {
				score, _ = newIA.GMax(board, list, searchList, move, ia.Depth-1, alpha, beta, indexListMoves+1, indexSearchList)
				chanScores <- score
				ia.Pool.Put(newIA)
			}(0)
			board[move.X][move.Y] = 0
		}
	}
loop:
	for true {
		select {
		case score := <-chanScores:
			if score < current {
				current = score
				best = move
			}
			if score < beta {
				beta = score
				best = move
			}
			if alpha >= beta {
				break loop
			}
			if len(chanScores) == 0 {
				break loop
			}
		default:
		}
	}
	return
}

func (ia *IA) Max(board board.Board, list, searchList Tlist, move inter.Node, depth, alpha, beta, indexListMoves, indexSearchList int) (current int, best inter.Node) {
	if result := ia.isWin(board, depth, move); result != 0 {
		return result, move
	} else if depth <= 0 {
		return ia.HeuristicScore(board, list, indexListMoves, move), move
	}
	current = math.MinInt64
	for i := 0; i < indexSearchList; i++ {
		move = searchList[i]
		if board[move.X][move.Y] == 0 {

			board[move.X][move.Y] = ia.PlayerIndex
			move.Player = ia.PlayerIndex
			list[indexListMoves] = move
			score, _ := ia.Min(board, list, searchList, move, depth-1, alpha, beta, indexListMoves+1, indexSearchList)
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
	return
}

func (ia *IA) Min(board board.Board, list, searchList Tlist, move inter.Node, depth, alpha, beta, indexListMoves, indexSearchList int) (current int, best inter.Node) {
	if result := ia.isWin(board, depth, move); result != 0 {
		return result, move
	} else if depth <= 0 {
		return ia.HeuristicScore(board, list, indexListMoves, move), move
	}
	current = math.MaxInt64
	playerIndex := player.GetOpposentPlayer(ia.PlayerIndex)
	for i := 0; i < indexSearchList; i++ {
		move = searchList[i]
		if board[move.X][move.Y] == 0 {
			board[move.X][move.Y] = playerIndex
			move.Player = playerIndex
			list[indexListMoves] = move
			score, _ := ia.Max(board, list, searchList, move, depth-1, alpha, beta, indexListMoves+1, indexSearchList)
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
	return
}
