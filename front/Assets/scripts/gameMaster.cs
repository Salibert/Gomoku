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
    private player CurrentPlayer;
    public player player1;
    public player player2;

    protected Channel channel;
    protected string GameID;
    protected Game.GameClient Client;
    void Awake() {
        channel = new Channel("127.0.0.1:50051", ChannelCredentials.Insecure);
        Client = new Game.GameClient(channel);
        GameID = Convert.ToBase64String(Guid.NewGuid().ToByteArray());
        CurrentPlayer = player1;
    }

    public void NextPlayer() {
        if (CurrentPlayer.index == player1.index) {
            CurrentPlayer = player2;
        } else {
            CurrentPlayer = player1;
        }
    }

    public int GetplayerTurn() { return CurrentPlayer.index; }
    public Material GetCurrentMaterial() { return CurrentPlayer.material; }

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
            if (reply.NbStonedCaptured != 0) {
                CurrentPlayer.SetScore(CurrentPlayer.GetScore() + reply.NbStonedCaptured);
                int index;
                GomokuBuffer.Node elementNode;
                foreach(GomokuBuffer.Node capturedStone in reply.Captured) {
                    index = 0;
                    foreach(stone el in goban.board) {
                        elementNode = el.GetNode();
                        if (elementNode.X == capturedStone.X && elementNode.Y == capturedStone.Y) {
                            el.Reset();
                            break;
                        }
                        index++;
                    }
                    goban.board.RemoveAt(index);
                }
            }
            return reply.IsPossible;
        } catch (Exception e) {
            Debug.Log("RPC failed" + e);
            throw;
        }
    }
}