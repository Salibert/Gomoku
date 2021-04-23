package player

import (
	"github.com/Salibert/Gomoku/back/rules"
	"github.com/Salibert/Gomoku/back/server/inter"
)

type Players map[int]*Player

// Player ...
type Player struct {
	Index           int
	NextMovesOrLose [][]*inter.Node
	Score           int
	Rules           rules.Schema
}

// Clone Players
func (players Players) Clone() Players {
	new := make(Players)
	for key, value := range players {
		new[key] = value.Clone()
	}
	return new
}

// Clone Player
func (player Player) Clone() *Player {
	return &Player{
		Index: player.Index,
		Rules: *player.Rules.Clone(),
	}
}

// GetOpposentPlayer return index to opponent player
func GetOpposentPlayer(player int) (opponent int) {
	switch player {
	case 1:
		opponent = 2
	default:
		opponent = 1
	}
	return
}

// CheckIfThisMoveBlockLose ...
func (player Player) CheckIfThisMoveBlockLose(lastMove *inter.Node) (checkLose bool) {
	for _, arrayBlockedStone := range player.NextMovesOrLose {
		checkLose = true
		for _, blockedStone := range arrayBlockedStone {
			if lastMove.X == blockedStone.X && lastMove.Y == blockedStone.Y {
				checkLose = false
			}
		}
		if checkLose == true {
			return
		}
	}
	return
}
