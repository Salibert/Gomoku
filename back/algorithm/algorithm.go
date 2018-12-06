package algorithm

import (
	"fmt"

	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

type Algo struct {
	Players player.Players
}

func IA_jouer(jeu board.Board, profondeur int, players player.Players) *pb.Node {
	var max int = -10000
	var tmp, maxi, maxj int
	var i, j int
	algo := Algo{Players: players}

	for i = 0; i < len(jeu); i++ {
		for j = 0; j < len(jeu); j++ {
			if jeu[i][j] == 0 {
				jeu[i][j] = 2
				tmp = algo.Minimax(&jeu, profondeur, 1, -10000, 10000, i, j)
				// fmt.Println("Value => ", tmp, " X = ", i, " Y= ", j)
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

func (algo *Algo) Minimax(jeu *board.Board, profondeur int, maximizingPlayer int, alpha int, beta int, x int, y int) int {
	if profondeur == 0 || algo.gagnant(jeu, x, y, maximizingPlayer) != 0 {
		resultat := algo.newEval(jeu, x, y, player.GetOpposentPlayer(int32(maximizingPlayer)))
		return resultat
	}
	if maximizingPlayer == 2 {
		value := -10000
		for i := 0; i < len(*jeu); i++ {
			for j := 0; j < len(*jeu); j++ {
				if (*jeu)[i][j] == 0 {
					(*jeu)[i][j] = 2
					value = Max(value, algo.Minimax(jeu, profondeur-1, 1, alpha, beta, i, j))
					fmt.Println("ALPHA VALUE => ", beta, " value => ", value)
					if alpha >= value {
						(*jeu)[i][j] = 0
						return value
					}
					alpha = value
					(*jeu)[i][j] = 0
				}
			}
		}
		return value
	} else {
		value := 10000
		for i := 0; i < len(*jeu); i++ {
			for j := 0; j < len(*jeu); j++ {
				if (*jeu)[i][j] == 0 {
					(*jeu)[i][j] = 1
					value = Min(value, algo.Minimax(jeu, profondeur-1, 2, alpha, beta, i, j))
					fmt.Println("BETA VALUE => ", beta, " value => ", value)
					if value >= beta {
						(*jeu)[i][j] = 0
						return value
					}
					beta = value
					(*jeu)[i][j] = 0
				}
			}
		}
		return value
	}
}

func Max(jeu int, profondeur int) int {
	fmt.Println("MAX VALUE => ", jeu, " algo => ", profondeur)
	if jeu < profondeur {
		return profondeur
	}
	return jeu
}

func Min(jeu int, profondeur int) int {
	fmt.Println("MIN VALUE => ", jeu, " algo => ", profondeur)
	if jeu > profondeur {
		return profondeur
	}
	return jeu
}

func (algo *Algo) newEval(jeu *board.Board, x int, y int, player int32) int {
	value := 0
	raport := algo.Players[player].Rules

	defer raport.Report.Reset()
	jeu.CheckRules(pb.Node{X: int32(x), Y: int32(y), Player: int32(player)}, raport)

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
	if x+1 < len(*jeu) && (*jeu)[x+1][y] == 1 {
		value += 2
	}
	if x-1 > 0 && (*jeu)[x-1][y] == 1 {
		value += 2
	}
	if y+1 < 19 && (*jeu)[x][y+1] == 1 {
		value += 2
	}
	if y-1 > 0 && (*jeu)[x][y-1] == 1 {
		value += 2
	}
	if x+1 < len(*jeu) && y+1 < len(*jeu) && (*jeu)[x+1][y+1] == 1 {
		value += 2
	}
	if y-1 > 0 && x+1 < len(*jeu) && (*jeu)[x+1][y-1] == 1 {
		value += 2
	}
	if x-1 > 0 && y+1 < len(*jeu) && (*jeu)[x-1][y+1] == 1 {
		value += 2
	}
	if x-1 > 0 && y-1 > 0 && (*jeu)[x-1][y-1] == 1 {
		value += 2
	}
	if x+1 < len(*jeu) && (*jeu)[x+1][y] == 1 {
		value += 2
	}
	// func() {
	// 	fmt.Println("\n", raport.Report)
	// 	fmt.Println("Value =>")
	// 	fmt.Println("")
	// 	fmt.Println((*jeu)[0])
	// 	fmt.Println((*jeu)[1])
	// 	fmt.Println((*jeu)[2])
	// 	fmt.Println((*jeu)[3])
	// 	fmt.Println((*jeu)[4])
	// 	fmt.Println((*jeu)[5])
	// 	fmt.Println((*jeu)[6])
	// 	fmt.Println((*jeu)[7])
	// 	fmt.Println((*jeu)[8])
	// 	fmt.Println((*jeu)[9])
	// 	fmt.Println((*jeu)[10])
	// 	fmt.Println((*jeu)[11])
	// 	fmt.Println((*jeu)[12])
	// 	fmt.Println((*jeu)[13])
	// 	fmt.Println((*jeu)[14])
	// 	fmt.Println((*jeu)[15])
	// 	fmt.Println((*jeu)[16])
	// 	fmt.Println((*jeu)[17])
	// 	fmt.Println((*jeu)[18])
	// 	fmt.Println("")
	// }()
	return value
}

func (algo *Algo) gagnant(jeu *board.Board, x int, y int, playerCurrent int) int {
	raport := algo.Players[int32(playerCurrent)].Rules

	defer raport.Report.Reset()
	jeu.CheckRules(pb.Node{X: int32(x), Y: int32(y), Player: player.GetOpposentPlayer(int32(playerCurrent))}, raport)
	if raport.Report.PartyFinish == true {
		if len(raport.Report.WinOrLose[0]) == 0 {
			return playerCurrent
		}
	}
	return 0
}
