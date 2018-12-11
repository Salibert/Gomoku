package algorithm

import (
	"fmt"

	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

type Algo struct {
	players     player.Players
	bestMove    pb.Node
	currentMove pb.Node
	alpha, beta int
}

// New ...
func New(config *pb.ConfigRules) {

}

func IA_jouer(jeu board.Board, profondeur int, players player.Players) *pb.Node {
	var max int = -10000
	var tmp int
	var maxi, maxj int32
	var i, j int32
	algo := Algo{players: players}
	algo.alpha, algo.beta = -10000, 10000
	for i = 0; i < board.SizeBoard; i++ {
		for j = 0; j < board.SizeBoard; j++ {
			if jeu[i][j] == 0 {
				jeu[i][j] = 2
				algo.currentMove.X, algo.currentMove.Y, algo.currentMove.Player = j, i, int32(1)
				tmp = algo.Alpha(jeu, 2, int32(1))
				fmt.Println("tmp >= ", tmp)
				if tmp > max {
					max = tmp
					maxi = algo.bestMove.X
					maxj = algo.bestMove.Y
				}
				jeu[i][j] = 0
			}
		}
	}
	return &pb.Node{X: int32(maxi), Y: int32(maxj), Player: int32(2)}
}

func (algo *Algo) Alpha(jeu board.Board, depth int, playerIndex int32) int {
	if depth == 0 {
		resultat := algo.newEval(jeu, player.GetOpposentPlayer(playerIndex))
		return resultat
	}
	if algo.gagnant(jeu, playerIndex) != 0 {
		return 10000
	}
	var i, j int
	SizeBoard := int(board.SizeBoard)
loop:
	for i = 0; i < SizeBoard; i++ {
		for j = 0; j < SizeBoard; j++ {
			if jeu[i][j] == 0 {
				jeu[i][j] = playerIndex
				algo.currentMove.X, algo.currentMove.Y, algo.currentMove.Player = int32(i), int32(j), playerIndex
				// fmt.Println("currentMove ALPHA => ", algo.currentMove, " Player => ", playerIndex)
				tmp := algo.currentMove
				score := algo.Beta(jeu, depth-1, player.GetOpposentPlayer(playerIndex))
				fmt.Println("SCORE ALPHA => ", score)
				if score > algo.alpha {
					algo.alpha = score
					algo.bestMove = tmp
					fmt.Println("TMP ALPHA => ", tmp)
				}
				jeu[i][j] = 0
				if algo.alpha >= algo.beta {
					fmt.Println("BREAK LOOP ALPHA", algo.alpha, algo.beta)
					break loop
				}
			}
		}
	}
	return algo.alpha
}

func (algo *Algo) Beta(jeu board.Board, depth int, playerIndex int32) int {
	if depth <= 0 {
		resultat := algo.newEval(jeu, player.GetOpposentPlayer(playerIndex))
		return resultat
	}
	if algo.gagnant(jeu, playerIndex) != 0 {
		return 10000
	}
	var i, j int
	SizeBoard := int(board.SizeBoard)
loop:
	for i = 0; i < SizeBoard; i++ {
		for j = 0; j < SizeBoard; j++ {
			if jeu[i][j] == 0 {
				jeu[i][j] = playerIndex
				algo.currentMove.X, algo.currentMove.Y, algo.currentMove.Player = int32(i), int32(j), playerIndex
				tmp := algo.currentMove
				score := algo.Beta(jeu, depth-1, player.GetOpposentPlayer(playerIndex))
				fmt.Println("SCORE BETA => ", score)
				if score < algo.beta {
					algo.beta = score
					algo.bestMove = tmp
					fmt.Println("TMP BETA => ", tmp)
				}
				jeu[i][j] = 0
				if algo.alpha >= algo.beta {
					fmt.Println("BREAK LOOP BETA", algo.alpha, algo.beta)
					break loop
				}
			}
		}
	}
	return algo.beta
}

// func (algo *Algo) Minimax(jeu *board.Board, profondeur int, maximizingPlayer int, alpha int, beta int, x int, y int) int {
// 	if profondeur == 0 || algo.gagnant(jeu, x, y, maximizingPlayer) != 0 {
// 		resultat := algo.newEval(jeu, x, y, player.GetOpposentPlayer(int32(maximizingPlayer)))
// 		return resultat
// 	}
// 	var v int
// 	if maximizingPlayer == 2 {
// 		// MAX
// 		value := -10000
// 		for i := 0; i < len(*jeu); i++ {
// 			for j := 0; j < len(*jeu); j++ {
// 				if (*jeu)[i][j] == 0 {
// 					(*jeu)[i][j] = 2
// 					v = algo.Minimax(jeu, profondeur-1, 1, alpha, beta, i, j)
// 					if v > value {
// 						value = v
// 					}
// 					if v >= beta {
// 						(*jeu)[i][j] = 0
// 						return value
// 					}
// 					if v < alpha {
// 						alpha = v
// 					}
// 					(*jeu)[i][j] = 0
// 				}
// 			}
// 		}
// 		return value
// 	} else {
// 		// MIN
// 		value := 10000
// 		for i := 0; i < len(*jeu); i++ {
// 			for j := 0; j < len(*jeu); j++ {
// 				if (*jeu)[i][j] == 0 {
// 					(*jeu)[i][j] = 1
// 					v = algo.Minimax(jeu, profondeur-1, 2, alpha, beta, i, j)
// 					// fmt.Println("BETA VALUE => ", beta, " value => ", value)
// 					if v < value {
// 						value = v
// 					}
// 					if v <= alpha {
// 						(*jeu)[i][j] = 0
// 						return value
// 					}
// 					if v < beta {
// 						beta = v
// 					}
// 					(*jeu)[i][j] = 0
// 				}
// 			}
// 		}
// 		return value
// 	}
// }

func (algo *Algo) newEval(jeu board.Board, player int32) int {
	value := 0
	raport := algo.players[player].Rules.Clone()
	defer raport.Report.Reset()
	var x, y int = int(algo.currentMove.X), int(algo.currentMove.Y)
	jeu.CheckRules(algo.currentMove, *raport)
	if raport.Report.ItIsAValidMove == false {
		return -10000
	} else {
		if capture := len(raport.Report.ListCapturedStone); capture != 0 {
			value += capture * 100
		}
		value += raport.Report.NbFreeThree * 5
		if raport.Report.PartyFinish == true {
			if counter := len(raport.Report.WinOrLose); counter == 0 {
				value += 100
			}
		}
		value += raport.Report.SizeAlignment * 5
	}
	if x+1 < len(jeu) && jeu[x+1][y] != 0 {
		value += 2
	}
	if x-1 > 0 && (jeu)[x-1][y] != 0 {
		value += 2
	}
	if y+1 < 19 && (jeu)[x][y+1] != 0 {
		value += 2
	}
	if y-1 > 0 && (jeu)[x][y-1] != 0 {
		value += 2
	}
	if x+1 < len(jeu) && y+1 < len(jeu) && (jeu)[x+1][y+1] != 0 {
		value += 2
	}
	if y-1 > 0 && x+1 < len(jeu) && (jeu)[x+1][y-1] != 0 {
		value += 2
	}
	if x-1 > 0 && y+1 < len(jeu) && (jeu)[x-1][y+1] != 0 {
		value += 2
	}
	if x-1 > 0 && y-1 > 0 && (jeu)[x-1][y-1] != 0 {
		value += 2
	}
	if x+1 < len(jeu) && (jeu)[x+1][y] != 0 {
		value += 2
	}
	if player == 1 {
		value *= -1
	}
	return value
}

func (algo *Algo) gagnant(jeu board.Board, playerCurrent int32) int {
	raport := algo.players[playerCurrent].Rules.Clone()
	jeu.CheckRules(algo.currentMove, *raport)
	if raport.Report.PartyFinish == true {
		if len(raport.Report.WinOrLose[0]) == 0 {
			return int(playerCurrent)
		}
	}
	return 0
}
