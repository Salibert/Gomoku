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
		best.Player = ia.playerIndex
	}
	best.Player = ia.playerIndex
	return &best
}

func (ia *IA) Boost(board board.Board, list, searchList Tlist, indexListMoves, indexSearchList int) (current int, best inter.Node) {
	current = math.MinInt64
	var finish bool
	alpha, beta := -100000, 1000000
	// halfSearchList := int(indexSearchList / 2)
	move := inter.Node{Player: player.GetOpposentPlayer(ia.playerIndex)}
	// for i := 0; i < indexSearchList; i++ {
	// 	move = searchList[i]
	// 	if board[move.X][move.Y] == 0 {

	// 		board[move.X][move.Y] = ia.playerIndex
	// 		move.Player = ia.playerIndex
	// 		list[indexListMoves] = move
	// 		score, _ := ia.Min(board, list, searchList, move, 0, alpha, beta, indexListMoves+1, indexSearchList)
	// 		board[move.X][move.Y] = 0
	// 		if score > current {
	// 			current = score
	// 			best = move
	// 		}
	// 		if score > alpha {
	// 			alpha = score
	// 			best = move
	// 		}
	// 		if alpha >= beta {
	// 			fmt.Println("SALUT ")
	// 			// finish = true
	// 			break
	// 		}
	// 	}
	// }
	for i := 0; i < indexSearchList; i++ {
		move = searchList[i]
		if board[move.X][move.Y] == 0 {

			board[move.X][move.Y] = ia.playerIndex
			move.Player = ia.playerIndex
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
				fmt.Println("SALUT ")
				// finish = true
				break
			}
		}
	}
	if finish == true {
		maxToken := 20
		token := maxToken
		chanScores := make(chan int, maxToken)
	loop:
		for i := 0; i < indexSearchList; i++ {
			move = searchList[i]
			if board[move.X][move.Y] == 0 {

				board[move.X][move.Y] = ia.playerIndex
				move.Player = ia.playerIndex
				list[indexListMoves] = move
				token--
				newIa := ia.Pool.Get().(*IA)
				fmt.Println("OK => ", ici)
				ici++
				go func() {
					score, _ := newIa.Min(board, list, searchList, move, ia.depth-1, alpha, beta, indexListMoves+1, indexSearchList)
					chanScores <- score
					ia.Pool.Put(newIa)
				}()
				board[move.X][move.Y] = 0
				if token == 0 {
					for true {
						if token > int(maxToken/2) {
							break
						}
						select {
						case score := <-chanScores:
							token++
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
						default:
						}
					}

				}
			}
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

			board[move.X][move.Y] = ia.playerIndex
			move.Player = ia.playerIndex
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
	playerIndex := player.GetOpposentPlayer(ia.playerIndex)
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

// func (ia *IA) MinMax(board board.Board, list [sizeListMoves]inter.Node, move inter.Node, depth, alpha, beta, index int, max bool) (current int, best inter.Node) {
// 	if result := ia.isWin(board, depth, move); result != 0 {
// 		return result, move
// 	} else if depth <= 0 {
// 		return ia.HeuristicScore(board, list, index, move), move
// 	}
// 	if max == true {
// 		current = math.MinInt64
// 		for _, move = range ia.SearchZone {
// 			if board[move.X][move.Y] == 0 {

// 				board[move.X][move.Y] = ia.playerIndex
// 				move.Player = ia.playerIndex
// 				list[index] = move
// 				score, _ := ia.MinMax(board, list, move, depth-1, alpha, beta, index+1, false)
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
// 				list[index] = move
// 				score, _ := ia.MinMax(board, list, move, depth-1, alpha, beta, index+1, true)
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
