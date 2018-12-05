package algorithm

import (
	"fmt"

	"github.com/Salibert/Gomoku/back/board"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// Struct to store attempt for the search algorithm
type Tentatives struct {
	poid  int
	board board.Board
}

//TODO: Reprendre le tuto de openclassroom !

func IA_jouer(jeu board.Board, profondeur int) *pb.Node {
	var max int = -10000
	var tmp, maxi, maxj int
	var i, j int

	for i = 0; i < len(jeu); i++ {
		for j = 0; j < len(jeu); j++ {
			if jeu[i][j] == 0 {
				jeu[i][j] = 1
				tmp = Min(jeu, profondeur-1)
				if tmp > max {
					max = tmp
					maxi = i
					maxj = j
				}
				jeu[i][j] = 0
			}
		}
	}
	return &pb.Node{X: int32(maxi), Y: int32(maxj), Player: int32(2)}
}

func Max(jeu board.Board, profondeur int) int {
	fmt.Println("Max profondeur = ", profondeur)
	if profondeur == 0 || gagnant(jeu) != 0 {
		//fmt.Println("Eval Max")
		return eval(jeu)
	}

	var max int = -10000
	var tmp int

	for i := 0; i < len(jeu); i++ {
		for j := 0; j < len(jeu); j++ {
			fmt.Println("MAX ===> i = ", i, " && j = ", j)
			fmt.Println("Max profondeur = ", profondeur)
			if jeu[i][j] == 0 {
				jeu[i][j] = 2
				tmp = Min(jeu, profondeur-1)
				if tmp > max {
					max = tmp
				}
				jeu[i][j] = 0
			}
		}
	}
	fmt.Println("======== Fin MAX ========")
	return max
}

func Min(jeu board.Board, profondeur int) int {
	//fmt.Println("Min profondeur = ", profondeur)
	if profondeur == 0 || gagnant(jeu) != 0 {
		//fmt.Println("Eval Min")
		return eval(jeu)
	}
	var min int = 10000
	var tmp int

	for i := 0; i < len(jeu); i++ {
		for j := 0; j < len(jeu); j++ {
			fmt.Println("MIN ===> i = ", i, " && j = ", j)
			if jeu[i][j] == 0 {
				jeu[i][j] = 1
				tmp = Max(jeu, profondeur-1)
				if tmp < min {
					min = tmp
				}
				jeu[i][j] = 0
			}
		}
	}
	fmt.Println("======== Fin MIN ========")
	return min
}

func nb_series(jeu board.Board, series_j1 *int, series_j2 *int, n int) int { //Compte le nombre de séries de n pions alignés de chacun des joueurs

	*series_j1 = 0
	*series_j2 = 0

	return 0
}

func eval(jeu board.Board) int {
	nb_de_pions := 0

	if vainqueur := gagnant(jeu); vainqueur != 0 {
		//On compte le nombre de pions présents sur le plateau
		for i := 0; i < len(jeu); i++ {
			for j := 0; j < len(jeu); j++ {
				if jeu[i][j] != 0 {
					nb_de_pions++
				}
			}
		}
		if vainqueur == 1 {
			return 1000 - nb_de_pions
		} else if vainqueur == 2 {
			return -1000 + nb_de_pions
		} else {
			return 0
		}
	}
	//On compte le nombre de séries de 2 pions alignés de chacun des joueurs
	series_j1, series_j2 := 0, 0
	nb_series(jeu, &series_j1, &series_j2, 2)
	return series_j1 - series_j2

}

func gagnant(jeu board.Board) int {
	var j1, j2 int

	nb_series(jeu, &j1, &j2, 3)

	if j1 != 0 {
		return 1
	} else if j2 != 0 {
		return 2
	} else {
		//Si le jeu n'est pas fini et que personne n'a gagné, on renvoie 0
		for i := 0; i < len(jeu); i++ {
			for j := 0; j < len(jeu); j++ {
				if jeu[i][j] == 0 {
					return 0
				}
			}
		}
	}
	//Si le jeu est fini et que personne n'a gagné, on renvoie 3
	return 3
}

// func gomokuShapeScore(consecutive int, openEnds int, currentTurn int) int {
// 	if openEnds == 0 && consecutive < 5 {
// 		return 0
// 	}
// 	switch consecutive {
// 	case 4:
// 		switch openEnds {
// 		case 1:
// 			if currentTurn == 1 {
// 				return 100000000
// 			}
// 			return 50
// 		case 2:
// 			if currentTurn == 1 {
// 				return 100000000

// 			}
// 			return 500000
// 		}
// 	case 3:
// 		switch openEnds {
// 		case 1:
// 			if currentTurn == 1 {
// 				return 7
// 			}
// 			return 5
// 		case 2:
// 			if currentTurn == 1 {
// 				return 10000
// 			}
// 			return 50
// 		}
// 	case 2:
// 		switch openEnds {
// 		case 1:
// 			return 3
// 		case 2:
// 			return 5
// 		}
// 	case 1:
// 		switch openEnds {
// 		case 1:
// 			return 1
// 		case 2:
// 			return 2
// 		}
// 	default:
// 		return 200000000
// 	}
// 	return 0
// }

// func analyzeHorizontalSetsForBlack(current_turn int, board board.Board) int {
// 	var score = 0
// 	var countConsecutive = 0
// 	var openEnds = 0

// 	for i := 0; i < len(board); i++ {
// 		for a := 0; a < len(board[i]); a++ {
// 			if board[i][a] == 2 {
// 				countConsecutive++
// 			} else if board[i][a] == 0 && countConsecutive > 0 {
// 				openEnds++
// 				score += gomokuShapeScore(countConsecutive,
// 					openEnds, current_turn)
// 				countConsecutive = 0
// 				openEnds = 1
// 			} else if board[i][a] == 0 {
// 				openEnds = 1
// 			} else if countConsecutive > 0 {
// 				score += gomokuShapeScore(countConsecutive,
// 					openEnds, current_turn)
// 				countConsecutive = 0
// 				openEnds = 0
// 			} else {
// 				openEnds = 0
// 			}
// 		}
// 		if countConsecutive > 0 {
// 			score += gomokuShapeScore(countConsecutive,
// 				openEnds, current_turn)
// 		}
// 		countConsecutive = 0
// 		openEnds = 0
// 	}
// 	return score
// }

// func bestGomokuMove(bturn int, depth int, board board.Board) *pb.Node {
// 	var xBest = -1
// 	var yBest = -1
// 	if bturn != 0 {
// 		var bestScore = -1000000000
// 		var color = 1
// 	} else {
// 		var bestScore = 1000000000
// 		var color = 2
// 	}
// 	var analysis, response int
// 	if depth%2 == 0 {
// 		var analTurn = 2
// 	} else {
// 		var analTurn = 1
// 	}
// 	var moves = get_moves()

// 	for i := moves.length - 1; i > moves.length-aiMoveCheck-1 && i >= 0; i-- {
// 		board[moves[i][1]][moves[i][2]] = color
// 		if depth == 1 {
// 			analysis = analyzeGomoku(analTurn, board)
// 		} else {
// 			response = bestGomokuMove(!bturn, depth-1)
// 			analysis = response[2]
// 		}
// 		board[moves[i][1]][moves[i][2]] = 0
// 		if (analysis > bestScore && bturn) ||
// 			(analysis < bestScore && !bturn) {
// 			bestScore = analysis
// 			xBest = moves[i][1]
// 			yBest = moves[i][2]
// 		}
// 	}

// 	return &pb.Node{X: xBest, Y: yBest, Player: 2}
// }
