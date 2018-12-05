package algorithm

import (
	"fmt"

	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/rules"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

type Algo struct {
	Report *rules.Schema
}

func IA_jouer(jeu board.Board, profondeur int, schema rules.Schema) *pb.Node {
	var max int = -10000
	var tmp, maxi, maxj int
	var i, j int
	algo := Algo{Report: &schema}

	fmt.Println("jeu = ", jeu)
	for i = 0; i < len(jeu); i++ {
		for j = 0; j < len(jeu); j++ {
			if jeu[i][j] == 0 {
				fmt.Println("[x] = ", i, " && [y] = ", j, " == [x][y] = ", jeu[i][j])
				jeu[i][j] = 2
				tmp = algo.Minimax(jeu, profondeur, 2, -10000, 10000, i, j)
				if tmp >= max {
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

func (algo *Algo) Minimax(jeu board.Board, profondeur int, maximizingPlayer int, alpha int, beta int, x int, y int) int {
	if profondeur == 0 || algo.gagnant(jeu, x, y, maximizingPlayer) != 0 {
		return algo.eval(jeu, x, y, maximizingPlayer)
	}
	if maximizingPlayer == 2 {
		value := -10000
		for i := 0; i < len(jeu); i++ {
			for j := 0; j < len(jeu); j++ {
				if jeu[i][j] == 0 {
					jeu[i][j] = 2
					value = Max(value, algo.Minimax(jeu, profondeur-1, 1, alpha, beta, i, j))
					if alpha >= value {
						return value
					}
					alpha = Max(alpha, value)
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
					value = Min(value, algo.Minimax(jeu, profondeur-1, 2, alpha, beta, i, j))
					if value >= beta {
						return value
					}
					beta = Min(beta, value)
					jeu[i][j] = 0
				}
			}
		}
		return value
	}
}

func Max(jeu int, profondeur int) int {
	if jeu > profondeur {
		return profondeur
	}
	return jeu
}

func Min(jeu int, profondeur int) int {
	if jeu < profondeur {
		return profondeur
	}
	return jeu
}

func nb_series(jeu board.Board, series_j1 *int, series_j2 *int, n int) int { //Compte le nombre de séries de n pions alignés de chacun des joueurs
	var compteur1, compteur2 int

	*series_j1 = 0
	*series_j2 = 0

	compteur1 = 0
	compteur2 = 0

	//Diagonale descendante
	for i := 0; i < len(jeu); i++ {
		for j := 0; j < len(jeu)-i; j++ {
			if jeu[i][j] == 1 {
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

func (algo *Algo) eval(jeu board.Board, x int, y int, player int) int {
	nb_de_pions := 0

	if vainqueur := algo.gagnant(jeu, x, y, player); vainqueur != 0 {
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

func (algo *Algo) gagnant(jeu board.Board, x int, y int, player int) int {

	jeu.CheckRules(pb.Node{X: int32(x), Y: int32(y), Player: int32(player)}, *algo.Report)
	jeu[x][y] = int32(0)
	if algo.Report.Report.PartyFinish == true && len(algo.Report.Report.WinOrLose) == 0 {
		if player == 1 {
			return 1
		}
		return 2
	} else if algo.Report.Report.PartyFinish == true && len(algo.Report.Report.WinOrLose) != 0 {
		if player == 1 {
			return 2
		}
		return 1
	}
	return 0
}
