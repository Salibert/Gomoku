package algorithm

import (
	"fmt"
	"math"

	"github.com/Salibert/Gomoku/back/rules"

	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	"github.com/Salibert/Gomoku/back/server/inter"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

type Algo struct {
	players      player.Players
	bestMove     inter.Node
	reportWin    map[int]rules.Schema
	reportEval   map[int]rules.Schema
	currentMove  inter.Node
	moveOpposent inter.Node
}

// New ...
func New(players player.Players, config pb.ConfigRules, moveOpposent inter.Node) Algo {
	algo := Algo{
		players:    players.Clone(),
		reportWin:  make(map[int]rules.Schema),
		reportEval: make(map[int]rules.Schema),
	}
	config.IsActiveRuleAlignment = true
	config.IsActiveRuleBlock = true
	config.IsActiveRuleProbableCapture = true
	configWin := pb.ConfigRules{
		IsActiveRuleWin:     config.IsActiveRuleWin,
		IsActiveRuleCapture: config.IsActiveRuleCapture,
	}
	algo.reportWin[1] = rules.New(1, 2, configWin)
	algo.reportWin[2] = rules.New(2, 1, configWin)
	algo.reportEval[1] = rules.New(1, 2, config)
	algo.reportEval[2] = rules.New(2, 1, config)
	algo.moveOpposent = moveOpposent
	return algo
}

func IA_jouer(jeu board.Board, profondeur int, players player.Players, config pb.ConfigRules, moveOpposent inter.Node) *inter.Node {
	MovesList := jeu.CreateSearchSpace(moveOpposent)
	var max int = -10000
	var tmp int
	// var maxi, maxj int
	var i, j int
	algo := New(players, config, moveOpposent)
	alpha, beta := -10000, 10000
	for _, move := range MovesList {
		if jeu[move.X][move.Y] == 0 {
			jeu[move.X][move.Y] = 2
			algo.currentMove.X, algo.currentMove.Y, algo.currentMove.Player = move.X, move.Y, 2
			tmp = algo.Beta(jeu, MovesList, 2, alpha, beta)
			jeu[move.X][move.Y] = 0
			if tmp > max {
				max = tmp
				// maxi = i
				// maxj = j
			}
			jeu[i][j] = 0
		}

	}
	// res := &inter.Node{X: maxi, Y: maxj, Player: 2}
	fmt.Println("BEST MOVE > ", algo.bestMove)
	algo.bestMove.Player = 2
	return &algo.bestMove
}

func distance(moveOpposent, currentMove inter.Node) int {
	return int(math.Sqrt(math.Pow(float64(moveOpposent.X-currentMove.X), 2) +
		math.Pow(float64(moveOpposent.Y-currentMove.Y), 2)))
}

func (algo *Algo) Alpha(jeu board.Board, MovesList []inter.Node, depth, alpha, beta int) int {
	if algo.gagnant(jeu) != 0 {
		return 999
	} else if depth <= 0 {
		return algo.HeuristicScore(jeu)
	}
	playerIndex := player.GetOpposentPlayer(algo.currentMove.Player)
loop:
	for _, move := range MovesList {
		if jeu[move.X][move.Y] == 0 {
			algo.currentMove.X, algo.currentMove.Y, algo.currentMove.Player = move.X, move.Y, playerIndex
			jeu[move.X][move.Y] = playerIndex
			score := algo.Beta(jeu, MovesList, depth-1, alpha, beta)
			jeu[move.X][move.Y] = 0
			if score > alpha {
				alpha = score
				algo.bestMove = move
				if alpha >= beta {
					fmt.Println("PRUNE ALPHA")
					break loop
				}
			}

		}
	}
	return alpha
}

func (algo *Algo) Beta(jeu board.Board, MovesList []inter.Node, depth, alpha, beta int) int {
	if algo.gagnant(jeu) != 0 {
		return 10000
	} else if depth <= 0 {
		return algo.HeuristicScore(jeu)
	}
	playerIndex := player.GetOpposentPlayer(algo.currentMove.Player)
loop:
	for _, move := range MovesList {
		if jeu[move.X][move.Y] == 0 {
			algo.currentMove.X, algo.currentMove.Y, algo.currentMove.Player = move.X, move.Y, playerIndex
			jeu[move.X][move.Y] = playerIndex
			score := algo.Alpha(jeu, MovesList, depth-1, alpha, beta)
			jeu[move.X][move.Y] = 0
			if score < beta {
				beta = score
				algo.bestMove = move
				if alpha >= beta {
					fmt.Println("PRUNE BETA")
					break loop
				}
			}
		}
	}
	return beta
}
