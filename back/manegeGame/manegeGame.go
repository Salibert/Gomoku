package manegeGame

import (
	"context"
	"sync"

	"github.com/Salibert/Gomoku/back/board"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// Game contains all the meta data of a part
type Game struct {
	rwmux                  sync.RWMutex
	board                  board.Board
	nbStoneCapturedPlayer1 int32
	nbStoneCapturedPlayer2 int32
}

// Games contains all current games
type Games struct {
	rwmux sync.RWMutex
	game  map[string]*Game
}

var (
	// CurrentGames Public var
	CurrentGames *Games
)

func init() {
	CurrentGames = &Games{game: make(map[string]*Game)}
}

// AddNewGame method for create a game after a call front
func (CurrentGames *Games) AddNewGame(gameID string) (*pb.CDGameResponse, error) {
	CurrentGames.rwmux.Lock()
	if _, ok := CurrentGames.game[gameID]; ok == true {
		return &pb.CDGameResponse{IsSuccess: false, Message: "GameID Already exists"}, nil
	}
	board := make(Board, 19, 19)
	CurrentGames.game[gameID] = &Game{board: board}
	for i := 0; i < 19; i++ {
		board[i] = make([]pb.Node, 19, 19)
	}
	CurrentGames.rwmux.Unlock()
	return &pb.CDGameResponse{IsSuccess: true}, nil
}

// DeleteGame method for delete a game in the map
func (CurrentGames *Games) DeleteGame(gameID string) (*pb.CDGameResponse, error) {
	CurrentGames.rwmux.Lock()
	defer CurrentGames.rwmux.Unlock()
	delete(CurrentGames.game, gameID)
	if _, ok := CurrentGames.game[gameID]; ok == false {
		return &pb.CDGameResponse{IsSuccess: true}, nil
	}
	return &pb.CDGameResponse{IsSuccess: false}, nil
}

// ProccessRules ...
func (CurrentGames *Games) ProccessRules(context context.Context, in *pb.StonePlayed) (checkRules *pb.CheckRulesResponse, err error) {

}

// PlayedAI choose the best move for win
func (CurrentGames *Games) PlayedAI(context context.Context, in *pb.StonePlayed) (*pb.StonePlayed, error) {
	game := CurrentGames.game[in.GameID]
	game.board[in.CurrentPlayerMove.X][in.CurrentPlayerMove.Y] = *in.CurrentPlayerMove.Player
	return &pb.StonePlayed{CurrentPlayerMove: &pb.Node{X: in.CurrentPlayerMove.X, Y: in.CurrentPlayerMove.Y + 1, Player: 2}}, nil
}
