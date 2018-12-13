package algorithm

import "github.com/Salibert/Gomoku/back/board"

// Heuristic ...
func (algo *Algo) Heuristic(jeu board.Board) int {
	return algo.HeuristicScore(jeu) - algo.HeuristicScore(jeu)
}

// HeuristicScore ...
func (algo *Algo) HeuristicScore(jeu board.Board) int {
	value := 0
	raport := algo.reportEval[algo.currentMove.Player]
	defer raport.Report.Reset()
	jeu.CheckRules(algo.currentMove, raport)
	if raport.Report.ItIsAValidMove == false {
		return -1000
	} else {
		if len(algo.players[algo.currentMove.Player].NextMovesOrLose) != 0 {
			if raport.Report.PartyFinish = algo.players[algo.currentMove.Player].CheckIfThisMoveBlockLose(&algo.currentMove); raport.Report.PartyFinish == true {
				return 1000
			}
		}
		if capture := len(raport.Report.ListCapturedStone); capture != 0 {
			value += capture * 25
		}
		value += raport.Report.NbFreeThree * 5
		if raport.Report.PartyFinish == true {
			if counter := len(raport.Report.WinOrLose); counter == 0 {
				value += 100
			}
		}
		value += raport.Report.SizeAlignment
		value += raport.Report.NbBlockStone
		value -= raport.Report.LevelCapture
	}
	return value
}

func (algo *Algo) gagnant(jeu board.Board) int {
	raport := algo.reportWin[algo.currentMove.Player]
	defer raport.Report.Reset()
	jeu.CheckRules(algo.currentMove, raport)
	if raport.Report.PartyFinish == true {
		if len(raport.Report.WinOrLose[0]) == 0 {
			return 1
		}
	} else if algo.players[algo.currentMove.Player].Score == 8 && len(raport.Report.ListCapturedStone) != 0 {
		return 1
	}
	return 0
}
