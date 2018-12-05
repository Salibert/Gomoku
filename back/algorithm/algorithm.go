package algorithm

import (
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

func minimax(jeu board.Board, profondeur int, maximizingPlayer int) int {
	if profondeur == 0 || gagnant(jeu) != 0 {
		return eval(jeu)
	}
	if maximizingPlayer == 2 {
		value := -10000
		for i := 0; i < len(jeu); i++ {
			for j := 0; j < len(jeu); j++ {
				if jeu[i][j] == 0 {
					jeu[i][j] = 2
					value = Max(value, minimax(jeu, depth-1, 1))
					jeu[i][j] = 0
				}
			}
		}
		return value
	} else {
		value := 10000
		for i := 0; i < len(jeu); i++ {
			for j := 0; j < len(jeu); j++ {
				if jeu[i][j] == 0 {
					jeu[i][j] = 1
					value = Min(value, minimax(jeu, depth-1, 2))
					jeu[i][j] = 0
				}
			}
		}
		return value
	}
}

func Max(jeu board.Board, profondeur int) int {
	if profondeur == 0 || gagnant(jeu) != 0 {
		return eval(jeu)
	}
	var max int = -10000
	var tmp int

	for i := 0; i < len(jeu); i++ {
		for j := 0; j < len(jeu); j++ {
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
	return max
}

func Min(jeu board.Board, profondeur int) int {
	if profondeur == 0 || gagnant(jeu) != 0 {
		return eval(jeu)
	}
	var min int = 10000
	var tmp int

	for i := 0; i < len(jeu); i++ {
		for j := 0; j < len(jeu); j++ {
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
	return min
}

func nb_series(jeu board.Board, series_j1 *int, series_j2 *int, n int) int { //Compte le nombre de séries de n pions alignés de chacun des joueurs
	var compteur1, compteur2 int

	*series_j1 = 0
	*series_j2 = 0

	compteur1 = 0
	compteur2 = 0

	//Diagonale descendante
	for i := 0; i < len(jeu); i++ {
		if jeu[i][i] == 1 {
			compteur1++
			compteur2 = 0
			if compteur1 == n {
				*series_j1++
			}
		} else if jeu[i][i] == 2 {
			compteur2++
			compteur1 = 0
			if compteur2 == n {
				*series_j2++
			}
		}
	}
	compteur1 = 0
	compteur2 = 0

	//Diagonale montante
	for i := 0; i < len(jeu); i++ {
		// fmt.Println("i = ", i, "\nlen = ", len(jeu))
		if jeu[i][(len(jeu)-1)-i] == 1 {
			compteur1++
			compteur2 = 0
			if compteur1 == n {
				*series_j1++
			}
		} else if jeu[i][(len(jeu)-1)-i] == 2 {
			compteur2++
			compteur1 = 0
			if compteur2 == n {
				*series_j2++
			}
		}
	}

	//En ligne
	for i := 0; i < len(jeu); i++ {
		compteur1 = 0
		compteur2 = 0

		//Horizontalement
		for j := 0; j < len(jeu); j++ {
			if jeu[i][j] == 1 {
				compteur1++
				compteur2 = 0
				if compteur1 == n {
					*series_j1++
				}
			} else if jeu[i][j] == 2 {
				compteur2++
				compteur1 = 0
				if compteur2 == n {
					*series_j2++
				}
			}
		}

		compteur1 = 0
		compteur2 = 0

		//Verticalement
		for j := 0; j < len(jeu); j++ {
			if jeu[j][i] == 1 {
				compteur1++
				compteur2 = 0
				if compteur1 == n {
					*series_j1++
				}
			} else if jeu[j][i] == 2 {
				compteur2++
				compteur1 = 0
				if compteur2 == n {
					*series_j2++
				}
			}
		}
	}

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
