package game

import (
	"sync"

	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// Game contains all the meta data of a part
type Game struct {
	rwmux   sync.RWMutex
	board   board.Board
	players map[int32]*player.Player
}

// New create new instance of Game
func New() *Game {
	game := &Game{
		board:   board.New(),
		players: make(map[int32]*player.Player),
	}
	game.players[1] = &player.Player{Index: 1}
	game.players[2] = &player.Player{Index: 2}
	return game
}

// ProccessRules ...
func (game *Game) ProccessRules(initialStone *pb.Node) (*pb.CheckRulesResponse, error) {
	game.rwmux.Lock()
	defer game.rwmux.Unlock()
	res := &pb.CheckRulesResponse{}
	if game.board[initialStone.X][initialStone.Y] != 0 {
		res.IsPossible = false
		return res, nil
	}
	if report := game.board.CheckRulesAndCaptured(*initialStone); report != nil {
		currentPlayer := game.players[initialStone.Player]
		lenListCapture := int32(len(report.ListCapturedStone))
		if lenListCapture != 0 {
			currentPlayer.Score += lenListCapture
			res.NbStonedCaptured = lenListCapture
			res.Captured = report.ListCapturedStone
			if currentPlayer.Score == 10 {
				res.PartyFinish = true
				res.WinIs = currentPlayer.Index
				return res, nil
			}
		} else {
			res.PartyFinish = report.PartyFinish
		}
		if len(currentPlayer.NextMovesOrLose) != 0 {
			if res.PartyFinish = currentPlayer.CheckIfThisMoveBlockLose(initialStone); res.PartyFinish == true {
				res.WinIs = player.GetOpposentPlayer(currentPlayer.Index)
			}
		}
		if report.PartyFinish == false && len(report.WinOrLose) != 0 {
			opposentPlayer := game.players[player.GetOpposentPlayer(initialStone.Player)]
			opposentPlayer.NextMovesOrLose = report.WinOrLose
		}
		res.IsPossible = true
	}
	return res, nil
}
