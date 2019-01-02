package solver

import (
	"github.com/Salibert/Gomoku/back/rules"
	"github.com/Salibert/Gomoku/back/server/inter"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// IA ...
type IA struct {
	SearchZone         []inter.Node
	reportWin          map[int]rules.Schema
	reportEval         map[int]rules.Schema
	minMax             Algo
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
	return regis
}
