package algorithm

import (
	"github.com/Salibert/Gomoku/back/board"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// Struct to store attempt for the search algorithm
type Tentatives struct {
	poid  int32
	board board.Board
}

// heuristic
func heuristic(x int32, y int32, CurrentBoard board.Board) {
	var res *pb.CheckRulesResponse
	poid := 0
	node := pb.Node{X: x, Y: y, Player: 2}

	if res = CurrentBoard.CheckRulesAndCaptured(node); res.IsPossible == true {
		poid += len(res.Captured)
	}
}

// Min Max algo to find the best move (FAUX)
func Minmax(CurrentBoard board.Board) {
	allMoves := make([]Tentatives, 0)
	for x, col := range CurrentBoard {
		for y, position := range col {
			if position == 0 {
				append(allMoves, heuristic(x, y, CurrentBoard))
			}
		}
	}
}

// Min Max algo to find the best move (FAUX)
func Simulate(CurrentBoard board.Board) {
	allMoves := make([]Tentatives, 0)
	for x, col := range CurrentBoard {
		for y, position := range col {
			if position == 0 {
				append(allMoves, heuristic(x, y, CurrentBoard))
			}
		}
	}
}

//TODO: Reprendre le tuto de openclassroom !
