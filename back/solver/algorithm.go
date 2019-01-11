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

type Tlist [361]inter.Node

func createListMoves(NewListMoves *Tlist, listMoves []inter.Node) {
	for key, el := range listMoves {
		NewListMoves[key] = el
	}
}

func createSearchList(NewListMoves *Tlist, listMoves []inter.Node) {
	lenListMoves := len(listMoves)
	lenListMoves--
	for _, el := range listMoves {
		NewListMoves[lenListMoves] = el
		lenListMoves--
	}
}

func createSearchListByNewListMoves(NewListMoves *Tlist, list [][]*inter.Node) int {
	var index int
	for _, el := range list {
		for _, element := range el {
			NewListMoves[index] = *element
			index++
		}
	}
	return index
}

func pushFront(list *Tlist, node inter.Node, lenList int) {
	var tmp, move inter.Node
	var offset int
	move = list[0]
	list[0] = node
	list[0].Player = 0
	for i := 1; i < lenList; i++ {
		tmp = list[i]
		if node.X == tmp.X && node.Y == tmp.Y {
			offset = 1
			continue
		}
		list[i-offset] = move
		move = tmp
	}
	list[lenList] = inter.Node{}
}

func (ia *IA) Play(board *board.Board, players player.Players) *inter.Node {
	ia.playersScore[0] = players[1].Score
	ia.playersScore[1] = players[2].Score
	var best inter.Node
	var listMoves, searchList Tlist
	var lenSearchZone int
	createListMoves(&listMoves, ia.ListMoves)
	if len(players[ia.PlayerIndex].NextMovesOrLose) != 0 {
		lenSearchZone = createSearchListByNewListMoves(&searchList, players[ia.PlayerIndex].NextMovesOrLose)
	} else {
		createSearchList(&searchList, ia.SearchZone)
		lenSearchZone = len(ia.SearchZone)
	}
	if len(ia.SearchZone) != 0 {
		start := time.Now()
		_, best = ia.Boost(*board, listMoves, searchList, len(ia.ListMoves), lenSearchZone)
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
	if ia.Depth > 3 {
		for i := 0; i < indexSearchList; i++ {
			move = searchList[i]
			if board[move.X][move.Y] == 0 {
				board[move.X][move.Y] = ia.PlayerIndex
				move.Player = ia.PlayerIndex
				list[indexListMoves] = move
				score, _ := ia.Min(board, list, searchList, move, 3-1, alpha, beta, indexListMoves+1, indexSearchList)
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
		pushFront(&searchList, best, indexSearchList+1)
		current, best = ia.Max(board, list, searchList, move, ia.Depth, alpha, beta, indexListMoves, indexSearchList)
	} else {
		current, best = ia.Max(board, list, searchList, move, ia.Depth, alpha, beta, indexListMoves, indexSearchList)
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
