using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class stone : MonoBehaviour
{
    private MeshRenderer meshRend;
    private Collider gravity;
    private Renderer rend;
    private gameMaster.node node;
    private bool isCreate;

    public void initNode(ref gameMaster.node n) { node = n; }
    void Start() {
        rend = GetComponent<Renderer>();
        meshRend = GetComponent<MeshRenderer>();
        gravity = GetComponent<Collider>();
    }

    void OnMouseDown() {
        if (!isCreate) {
            rend.material = goban.currentGM.getCurrentMaterial();
            node.player = goban.currentGM.getplayerTurn();
            goban.currentGM.nextPlayer();
            meshRend.enabled = true;
            isCreate = true;
            gravity.attachedRigidbody.useGravity = true;
        }
    }
    void OnMouseEnter() {
        if (!isCreate) {
            rend.material = goban.currentGM.getCurrentMaterial();
            meshRend.enabled = true;
        }
    }

    void OnMouseExit() { if (!isCreate) meshRend.enabled = false;}
}
