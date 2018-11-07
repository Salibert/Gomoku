using UnityEngine;
public class gameMaster : MonoBehaviour
{
    public struct node {
        public Vector2Int posArray;
        public int player;
        public bool captered;
    }
    private int playerTurn;
    private Material materialCurrentPlayer;
    public Material player1;
    public Material player2;
    void Start() {
        playerTurn = 1;
        materialCurrentPlayer = player1;
    }

    public void nextPlayer() {
        if (playerTurn == 1) {
            playerTurn = 2;
            materialCurrentPlayer = player2;
        } else {
            playerTurn = 1;
            materialCurrentPlayer = player1;
        }
    }

    public int getplayerTurn() { return playerTurn; }
    public Material getCurrentMaterial() { return materialCurrentPlayer; }
}