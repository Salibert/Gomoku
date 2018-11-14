using UnityEngine;
using UnityEngine.UI;
using Grpc.Core;
using GomokuBuffer;
using System;

public class gameMaster : MonoBehaviour
{
    private int playerTurn;
    private Material materialCurrentPlayer;
    public Material player1;
    public Material player2;

    private Channel channel;
    private string gameID;
    private Game.GameClient Client;
    void Awake() {
        channel = new Channel("127.0.0.1:50051", ChannelCredentials.Insecure);
        Client = new Game.GameClient(channel);
        gameID = Convert.ToBase64String(Guid.NewGuid().ToByteArray());
        playerTurn = 1;
        materialCurrentPlayer = player1;
    }

    public void NextPlayer() {
        if (playerTurn == 1) {
            playerTurn = 2;
            materialCurrentPlayer = player2;
        } else {
            playerTurn = 1;
            materialCurrentPlayer = player1;
        }
    }

    public int GetplayerTurn() { return playerTurn; }
    public Material GetCurrentMaterial() { return materialCurrentPlayer; }

    public Game.GameClient GetClient() {
        return Client;
    }
    public Channel GetChannel() {
        return channel;
    }

    public string GetGameID() {
        return gameID;
    }
}