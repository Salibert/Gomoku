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
    public Transform Player1;
    public Transform Player2;
    private player player1;
    private player player2;

    protected Channel channel;
    protected string GameID;
    protected Game.GameClient Client;
    public GameObject finishGameUI;
    public Text winner;
    public Text time;

    void Awake() {
        channel = new Channel("127.0.0.1:50051", ChannelCredentials.Insecure);
        Client = new Game.GameClient(channel);
        GameID = Convert.ToBase64String(Guid.NewGuid().ToByteArray());
        CurrentPlayer = Player1.GetComponent<player>();
        finishGameUI.SetActive(false);
        player1 = CurrentPlayer;
        player2 = Player2.GetComponent<player>();
    }

	public void PartyFinish(int player) {
        winner.text = "Player " + player.ToString();
		finishGameUI.SetActive(true);
        GetCDGame(true);
	}
    public void NextPlayer() {
        if (CurrentPlayer.GetIndex() == player1.GetIndex()) {
            CurrentPlayer = player2;
        } else {
            CurrentPlayer = player1;
        }
    }

    public int GetPlayerTurn() { return CurrentPlayer.GetIndex(); }
    public Material GetCurrentMaterial() { return CurrentPlayer.GetMaterial(); }

    public Game.GameClient GetClient() {
        return Client;
    }
    public Channel GetChannel() {
        return channel;
    }
    public string GetGameID() {
        return GameID;
    }
    async public void GetCDGame(bool delete) {    
        try {
            Debug.Log(mainMenu.config);
            GomokuBuffer.CDGameResponse reply = await Client.CDGameAsync(
                new GomokuBuffer.CDGameRequest(){
                    GameID= GameID,
                    Rules= mainMenu.config.Clone(),
                    Delete= delete,
                });
            if (reply.IsSuccess == false)
                Debug.Log("NONONONO");
        } catch (Exception e) {
            Debug.Log("RPC failed" + e);
            throw;
        }
    }

    async public void GetPlayed(GomokuBuffer.Node node) {
        DateTime start = DateTime.Now;
        try {
            GomokuBuffer.StonePlayed reply = await Client.PlayedAsync(
                new GomokuBuffer.StonePlayed(){ CurrentPlayerMove=node.Clone(), GameID=GameID  });
            TimeSpan end = DateTime.Now.Subtract(start);
            time.text = end.ToString();
            await GetCheckRules(reply.CurrentPlayerMove, reply.CurrentPlayerMove.Player);
            Transform stone = goban.GetStone(reply.CurrentPlayerMove);
            stone.transform.GetComponent<stone>().SetStone();
            goban.board.Add(stone.transform.GetComponent<stone>());
        } catch (Exception e) {
            Debug.Log("RPC failed" + e);
            throw;
        }
    }

    private void updateCapture(GomokuBuffer.CheckRulesResponse reply) {
        if (reply.NbStonedCaptured != 0) {
            CurrentPlayer.SetScore(CurrentPlayer.GetScore() + reply.NbStonedCaptured);
            Debug.Log("SALUT TOI" + reply.NbStonedCaptured);
            goban.zoneCapture[CurrentPlayer.GetIndex()].GetComponent<zoneCapture>().AddStone(CurrentPlayer.GetScore());
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
                if (index <= goban.board.Count)
                    goban.board.RemoveAt(index);
            }
        }
    }

    async public Task<bool> GetCheckRules(GomokuBuffer.Node node, int player) {
        try {
            node.Player = player;
            GomokuBuffer.CheckRulesResponse reply = await Client.CheckRulesAsync(
                new GomokuBuffer.StonePlayed(){ CurrentPlayerMove=node.Clone(), GameID=GameID });
            Debug.Log("List capture "+ reply.Captured);
            updateCapture(reply);
            if (reply.PartyFinish == true) {
                PartyFinish(reply.IsWin);
            }
            return reply.IsPossible;
        } catch (Exception e) {
            Debug.Log("RPC failed" + e);
            throw;
        }
    }
    public player GetPlayer(int index) {
        if (index == 1){
            return player1;
        } else {
            return player2;
        }
    }
}