package manegeGame

import (
	"sync"

	"github.com/Salibert/Gomoku/back/game"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// Games contains all current games
type Games struct {
	rwmux sync.RWMutex
	game  map[string]*game.Game
}

var (
	// CurrentGames Public var
	CurrentGames *Games
)

func init() {
	CurrentGames = &Games{game: make(map[string]*game.Game)}
}

// AddNewGame method for create a game after a call front
func (CurrentGames *Games) AddNewGame(gameID string) (*pb.CDGameResponse, error) {
	CurrentGames.rwmux.Lock()
	defer CurrentGames.rwmux.Unlock()
	if _, ok := CurrentGames.game[gameID]; ok == true {
		return &pb.CDGameResponse{IsSuccess: false, Message: "GameID Already exists"}, nil
	}
	CurrentGames.game[gameID] = game.New()
	return &pb.CDGameResponse{IsSuccess: true}, nil
}

// DeleteGame method for delete a game in the map
func (CurrentGames *Games) DeleteGame(gameID string) (res *pb.CDGameResponse, err error) {
	CurrentGames.rwmux.Lock()
	defer CurrentGames.rwmux.Unlock()
	delete(CurrentGames.game, gameID)
	if _, ok := CurrentGames.game[gameID]; ok == false {
		res.IsSuccess = true
	}
	return res, err
}

// ProccessRules ...
func (CurrentGames *Games) ProccessRules(in *pb.StonePlayed) (*pb.CheckRulesResponse, error) {
	CurrentGames.rwmux.RLock()
	defer CurrentGames.rwmux.RUnlock()
	return CurrentGames.game[in.GameID].ProccessRules(in.CurrentPlayerMove)
}

// PlayedAI choose the best move for win
func (CurrentGames *Games) PlayedAI(in *pb.StonePlayed) (*pb.StonePlayed, error) {
	// game := CurrentGames.game[in.GameID]
	// game.board[in.CurrentPlayerMove.X][in.CurrentPlayerMove.Y] = in.CurrentPlayerMove.Player
	return &pb.StonePlayed{CurrentPlayerMove: &pb.Node{X: in.CurrentPlayerMove.X, Y: in.CurrentPlayerMove.Y + 1, Player: 2}}, nil
}
