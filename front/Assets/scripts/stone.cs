using System.Collections;
using System.Collections.Generic;
using UnityEngine;
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

    void OnMouseDown() {
        if (!isCreate && goban.GM.GetplayerTurn() == 1) {
            SetStone();
            goban.GM.GetPlayed(node, "OKOK");
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

    public void SetStone() {
        rend.material = goban.GM.GetCurrentMaterial();
        node.Player = goban.GM.GetplayerTurn();
        goban.GM.NextPlayer();
        isCreate = true;
        gravity.attachedRigidbody.useGravity = true;
        meshRend.enabled = true;
    }
}
