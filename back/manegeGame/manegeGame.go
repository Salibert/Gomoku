package manegeGame

import (
	"errors"
	"sync"

	"github.com/Salibert/Gomoku/back/game"
	"github.com/Salibert/Gomoku/back/server/inter"
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
func (CurrentGames *Games) AddNewGame(in *pb.CDGameRequest) (*pb.CDGameResponse, error) {
	CurrentGames.rwmux.Lock()
	defer CurrentGames.rwmux.Unlock()
	if _, ok := CurrentGames.game[in.GameID]; ok == true {
		return &pb.CDGameResponse{IsSuccess: false, Message: "GameID Already exists"}, nil
	}
	CurrentGames.game[in.GameID] = game.New(*in.Rules)
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
	CurrentGames.rwmux.Lock()
	defer CurrentGames.rwmux.Unlock()
	if game, ok := CurrentGames.game[in.GameID]; ok == true {
		return game.ProccessRules(inter.NewNode(in.CurrentPlayerMove))
	}
	return nil, errors.New("partie not found")
}

// PlayedAI choose the best move for win
func (CurrentGames *Games) PlayedIA(in *pb.StonePlayed) (*pb.StonePlayed, error) {
	game := CurrentGames.game[in.GameID]
	return &pb.StonePlayed{CurrentPlayerMove: game.PlayIA(inter.NewNode(in.CurrentPlayerMove)).Convert()}, nil
}
