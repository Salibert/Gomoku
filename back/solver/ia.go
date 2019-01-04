package solver

import (
	"sync"

	"github.com/Salibert/Gomoku/back/board"

	"github.com/Salibert/Gomoku/back/rules"
	"github.com/Salibert/Gomoku/back/server/inter"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

var Pool *sync.Pool
var salut int

func init() {
	Pool = &sync.Pool{
		New: func() interface{} {
			// board := make(board.Board, 19, 19)
			// for index := 0; index < 19; index++ {
			// 	board[index] = make([]int, 19, 19)
			// }
			return board.Board{}
		},
	}
}

// IA ...
type IA struct {
	SearchZone         []inter.Node
	ListMoves          []inter.Node
	reportWin          map[int]rules.Schema
	reportEval         map[int]rules.Schema
	playersScore       [2]int
	playerIndex, depth int
}

// New ...
func New(config pb.ConfigRules, playerIndex int) *IA {
	regis := &IA{
		reportWin:  make(map[int]rules.Schema),
		reportEval: make(map[int]rules.Schema),
	}
	regis.playerIndex = playerIndex
	config.IsActiveRuleAlignment = true
	config.IsActiveRuleBlock = true
	if config.IsActiveRuleCapture == true {
		config.IsActiveRuleProbableCapture = true
	}
	config.IsActiveRuleAmbientSearch = true
	configWin := pb.ConfigRules{
		IsActiveRuleWin:     config.IsActiveRuleWin,
		IsActiveRuleCapture: config.IsActiveRuleCapture,
	}
	regis.depth = int(config.DepthIA)
	regis.playerIndex = int(config.PlayerIndexIA)
	regis.reportWin[1] = rules.New(1, 2, configWin)
	regis.reportWin[2] = rules.New(2, 1, configWin)
	regis.reportEval[1] = rules.New(1, 2, config)
	regis.reportEval[2] = rules.New(2, 1, config)
	regis.SearchZone = make([]inter.Node, 0, 361)
	regis.ListMoves = make([]inter.Node, 0, 361)
	return regis
}

func (ia *IA) UpdateListMove(listCapture []*inter.Node, lastMove inter.Node) {
	if len(listCapture) != 0 {
		for _, capture := range listCapture {
			for i, list := range ia.ListMoves {
				if list == *capture {
					lenListMoves := len(ia.ListMoves)
					copy((ia.ListMoves)[i:], (ia.ListMoves)[i+1:])
					(ia.ListMoves)[lenListMoves-1] = inter.Node{}
					(ia.ListMoves) = (ia.ListMoves)[:lenListMoves-1]
					break
				}
			}
		}
	}
	ia.ListMoves = append(ia.ListMoves, lastMove)
}
