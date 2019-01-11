package game

import (
	"fmt"
	"sync"

	"github.com/Salibert/Gomoku/back/board"
	"github.com/Salibert/Gomoku/back/player"
	"github.com/Salibert/Gomoku/back/rules"
	"github.com/Salibert/Gomoku/back/server/inter"
	pb "github.com/Salibert/Gomoku/back/server/pb"
	"github.com/Salibert/Gomoku/back/solver"
)

// Game contains all the meta data of a part
type Game struct {
	rwmux   sync.RWMutex
	board   *board.Board
	players player.Players
	IA      *solver.IA
}

// New create new instance of Game
func New(config pb.ConfigRules) *Game {
	game := &Game{
		board:   &board.Board{},
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
	if config.PlayerIndexIA != 0 {
		game.IA = solver.New(config, int(config.PlayerIndexIA))
	} else {
		game.IA = nil
	}
	return game
}

// UpdateGame proccess all Update after a modif board
func (game *Game) UpdateGame(player *player.Player, initialStone *inter.Node) {
	game.board.UpdateBoardAfterCapture(&player.Rules)
	if game.IA != nil {
		if game.IA.Depth > 3 {
			game.board.UpdateSearchSpace(&game.IA.SearchZone, *initialStone, 1)
		} else {
			game.board.UpdateSearchSpace(&game.IA.SearchZone, *initialStone, 2)
		}
		game.IA.UpdateListMove(player.Rules.Report.ListCapturedStone, *initialStone)
	}
	player.Rules.Report.Reset()
	game.rwmux.Unlock()
}

// ProccessRules ...
func (game *Game) ProccessRules(initialStone *inter.Node) (*pb.CheckRulesResponse, error) {
	game.rwmux.Lock()
	res := &pb.CheckRulesResponse{}
	currentPlayer := game.players[initialStone.Player]
	defer func() {
		go game.UpdateGame(currentPlayer, initialStone)
	}()
	if game.board[initialStone.X][initialStone.Y] != 0 {
		res.IsPossible = false
		return res, nil
	}
	game.board.CheckRules(*initialStone, currentPlayer.Rules)
	if currentPlayer.Rules.Report.ItIsAValidMove == true {
		lenListCapture := len(currentPlayer.Rules.Report.ListCapturedStone)
		if lenListCapture != 0 {
			currentPlayer.Score += lenListCapture
			res.NbStonedCaptured = int32(lenListCapture)
			res.Captured = inter.ConvertArrayNode(currentPlayer.Rules.Report.ListCapturedStone)
			if currentPlayer.Score == 10 {
				res.PartyFinish = true
				res.IsWin = int32(currentPlayer.Index)
				return res, nil
			}
		} else {
			res.PartyFinish = currentPlayer.Rules.Report.PartyFinish
			if res.PartyFinish == true {
				res.IsWin = int32(currentPlayer.Index)
			}
		}
		if len(currentPlayer.NextMovesOrLose) != 0 {
			if res.PartyFinish = currentPlayer.CheckIfThisMoveBlockLose(initialStone); res.PartyFinish == true {
				res.IsWin = int32(player.GetOpposentPlayer(currentPlayer.Index))
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

func (game *Game) PlayIA(in *inter.Node, isHelp bool) *inter.Node {
	if isHelp == true {
		depth := game.IA.Depth
		game.IA.PlayerIndex = player.GetOpposentPlayer(game.IA.PlayerIndex)
		game.IA.Depth = 3
		defer func() {
			game.IA.PlayerIndex = player.GetOpposentPlayer(game.IA.PlayerIndex)
			game.IA.Depth = depth
		}()
	}
	fmt.Println("LEN => ", len(game.IA.SearchZone))
	return game.IA.Play(game.board, game.players)
}
