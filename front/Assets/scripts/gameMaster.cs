using UnityEngine;
public class gameMaster : MonoBehaviour
{
    public struct node {
        public Vector3 position;
        public Vector2Int posArray;
        public int index;
        public int player;
        public bool captered;
    }
    public static int player;

    void Start() { player = 1; }

    void nextPlayer() { player = player == 1 ? 2 : 1; }
}
