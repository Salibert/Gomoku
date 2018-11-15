using UnityEngine;
using UnityEngine.UI;
using Grpc.Core;
using GomokuBuffer;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

public class gameMaster : MonoBehaviour
{
    private int playerTurn;
    private Material materialCurrentPlayer;
    public Material player1;
    public Material player2;

    protected Channel channel;
    protected string gameID;
    protected Game.GameClient Client;
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
    async public void GetInitGame(List<GomokuBuffer.Node> sentedBoard) {    
        try {
            GomokuBuffer.InitGameResponse reply = await Client.InitGameAsync(
                new GomokuBuffer.InitGameRequest(){ Board= { sentedBoard.ToArray() }, GameId= gameID
            });
            Debug.Log(reply.Message);
        } catch (Exception e) {
            Debug.Log("RPC failed" + e);
            throw;
        }
    }

    async public void GetPlayed(GomokuBuffer.Node node, string message) {
        try {
            GomokuBuffer.StonePlayed reply = await Client.PlayedAsync(
                new GomokuBuffer.StonePlayed(){ CurrentPlayerMove=node.Clone(), Message=message
            });
            Transform stone = goban.GetStone(reply.CurrentPlayerMove);
            stone.transform.GetComponent<stone>().SetStone();
        } catch (Exception e) {
            Debug.Log("RPC failed" + e);
            throw;
        }
    }
}