using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Threading.Tasks;
using GomokuBuffer;

public class stone : MonoBehaviour
{
    private MeshRenderer meshRend;
    private Collider gravity;
    private Renderer rend;
    private GomokuBuffer.Node node;
    private bool isCreate;

    public void initNode(ref GomokuBuffer.Node n) { node = n; }
    void Start() {
        rend = GetComponent<Renderer>();
        meshRend = GetComponent<MeshRenderer>();
        gravity = GetComponent<Collider>();
    }

    async void OnMouseDown() {
        if (!isCreate) {
            if (await goban.GM.GetCheckRules(node, goban.GM.GetplayerTurn())) {
                SetStone();
                goban.board.Add(transform);
            } else {
                Debug.Log("IMPOSSIBLE");
            }
        }
    }
    void OnMouseEnter() {
        if (!isCreate) {
            rend.material = goban.GM.GetCurrentMaterial();
            meshRend.enabled = true;
        }
    }

    void OnMouseExit() {
        if (!isCreate) {
            meshRend.enabled = false;
        }
    }

    public void Reset() {
        meshRend.enabled = false;
        node.Player = 0;
        isCreate = false;
        gravity.attachedRigidbody.useGravity = false;
    }
    public void SetStone() {
        rend.material = goban.GM.GetCurrentMaterial();
        node.Player = goban.GM.GetplayerTurn();
        goban.GM.NextPlayer();
        isCreate = true;
        gravity.attachedRigidbody.useGravity = true;
        meshRend.enabled = true;
    }
}
