package player

import (
	"github.com/Salibert/Gomoku/back/rules"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

type Players map[int32]*Player

// Player ...
type Player struct {
	Index           int32
	NextMovesOrLose [][]*pb.Node
	Score           int32
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

// GetOpposentPlayer return index to opposent player
func GetOpposentPlayer(player int32) (opposent int32) {
	switch player {
	case 1:
		opposent = 2
	default:
		opposent = 1
	}
	return
}

// CheckIfThisMoveBlockLose ...
func (player Player) CheckIfThisMoveBlockLose(lastMove *pb.Node) (checkLose bool) {
	for _, arrayBlockedStone := range player.NextMovesOrLose {
		checkLose = true
		for _, blockedStone := range arrayBlockedStone {
			if lastMove == blockedStone {
				checkLose = false
			}
		}
		if checkLose == true {
			return
		}
	}
	return
}
