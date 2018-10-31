using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class stone : MonoBehaviour
{
    private MeshRenderer rend;
    private Collider gravity;
    private static int index = 0;
    private gameMaster.node node;
    private bool isCreate;

    public void initNode(ref gameMaster.node n) { node = n; }
    void Start() {
        index++;
        rend = GetComponent<MeshRenderer>();
        gravity = GetComponent<Collider>();
    }

    void OnMouseDown() {
        if (!isCreate) {
            rend.enabled = true;
            isCreate = true;
            gravity.attachedRigidbody.useGravity = true;
        }
    }
    void OnMouseEnter() { if (!isCreate) rend.enabled = true; }

    void OnMouseExit() { if (!isCreate) rend.enabled = false; }

    public int getIndex() { return index; }
}
