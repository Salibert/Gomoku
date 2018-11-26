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
    protected string GameID;
    protected Game.GameClient Client;
    void Awake() {
        channel = new Channel("127.0.0.1:50051", ChannelCredentials.Insecure);
        Client = new Game.GameClient(channel);
        GameID = Convert.ToBase64String(Guid.NewGuid().ToByteArray());
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
        return GameID;
    }
    async public void GetCDGame() {    
        try {
            GomokuBuffer.CDGameResponse reply = await Client.CDGameAsync( new GomokuBuffer.CDGameRequest(){ GameID= GameID });
            if (reply.IsSuccess == false)
                Debug.Log("NONONONO");
        } catch (Exception e) {
            Debug.Log("RPC failed" + e);
            throw;
        }
    }

    async public void GetPlayed(GomokuBuffer.Node node) {
        try {
            GomokuBuffer.StonePlayed reply = await Client.PlayedAsync(
                new GomokuBuffer.StonePlayed(){ CurrentPlayerMove=node.Clone(), GameID=GameID  });
            Transform stone = goban.GetStone(reply.CurrentPlayerMove);
            stone.transform.GetComponent<stone>().SetStone();
        } catch (Exception e) {
            Debug.Log("RPC failed" + e);
            throw;
        }
    }

    async public Task<bool> GetCheckRules(GomokuBuffer.Node node, int player) {
        try {
            node.Player = player;
            GomokuBuffer.CheckRulesResponse reply = await Client.CheckRulesAsync(
                new GomokuBuffer.StonePlayed(){ CurrentPlayerMove=node.Clone(), GameID=GameID  });
            if (reply.Captured) {
                int indexElementDelete;
                foreach(GomokuBuffer.Node capturedStone in reply.Captured) {
                    goban.board.ForEach((index, el) => {
                        if (el.transform.node.X == capturedStone.X && el.transform.node.Y == capturedStone.Y) {
                            el.transform.GetComponent<stone>().Reset();
                            indexElementDelete = index;
                            break;
                        }
                    });
                    goban.board.Remove(indexElementDelete);
                }
            }
            return reply.IsPossible;
        } catch (Exception e) {
            Debug.Log("RPC failed" + e);
            throw;
        }
    }
}