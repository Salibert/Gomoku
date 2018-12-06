package game

import (
	"sync"

	"github.com/Salibert/Gomoku/back/algorithm"
	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	"github.com/Salibert/Gomoku/back/rules"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// Game contains all the meta data of a part
type Game struct {
	rwmux   sync.RWMutex
	board   board.Board
	players player.Players
}

// New create new instance of Game
func New(config pb.ConfigRules) *Game {
	game := &Game{
		board:   board.New(),
		players: make(player.Players),
	}
	game.players[1] = &player.Player{
		Index: 1,
		Rules: rules.New(1, 2, config),
	}
	game.players[2] = &player.Player{
		Index: 2,
		Rules: rules.New(2, 1, config),
	}
	return game
}

// ProccessRules ...
func (game *Game) ProccessRules(initialStone *pb.Node) (*pb.CheckRulesResponse, error) {
	game.rwmux.Lock()
	res := &pb.CheckRulesResponse{}
	currentPlayer := game.players[initialStone.Player]
	defer func() {
		game.board.UpdateBoardAfterCapture(&currentPlayer.Rules)
		currentPlayer.Rules.Report.Reset()
		game.rwmux.Unlock()
	}()
	if game.board[initialStone.X][initialStone.Y] != 0 {
		res.IsPossible = false
		return res, nil
	}
	game.board.CheckRules(*initialStone, currentPlayer.Rules)
	if currentPlayer.Rules.Report.ItIsAValidMove == true {
		lenListCapture := int32(len(currentPlayer.Rules.Report.ListCapturedStone))
		if lenListCapture != 0 {
			currentPlayer.Score += lenListCapture
			res.NbStonedCaptured = lenListCapture
			res.Captured = currentPlayer.Rules.Report.ListCapturedStone
			if currentPlayer.Score == 10 {
				res.PartyFinish = true
				res.WinIs = currentPlayer.Index
				return res, nil
			}
		} else {
			res.PartyFinish = currentPlayer.Rules.Report.PartyFinish
		}
		if len(currentPlayer.NextMovesOrLose) != 0 {
			if res.PartyFinish = currentPlayer.CheckIfThisMoveBlockLose(initialStone); res.PartyFinish == true {
				res.WinIs = player.GetOpposentPlayer(currentPlayer.Index)
			}
		}
		if currentPlayer.Rules.Report.PartyFinish == false &&
			len(currentPlayer.Rules.Report.WinOrLose) != 0 {
			opposentPlayer := game.players[player.GetOpposentPlayer(initialStone.Player)]
			opposentPlayer.NextMovesOrLose = currentPlayer.Rules.Report.WinOrLose
		}
		game.board.UpdateBoard(*initialStone)
		res.IsPossible = true
	}
	return res, nil
}

func (game *Game) PlayIA(in *pb.Node) *pb.Node {
	node := algorithm.IA_jouer(game.board, 2, game.players)
	return node
}
