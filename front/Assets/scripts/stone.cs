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

    delegate void Played();
    Played modePlayed;

    delegate void Render();
    Render renderStone;
    public void initNode(ref GomokuBuffer.Node n) { node = n; }

    void Start() {
        rend = GetComponent<Renderer>();
        meshRend = GetComponent<MeshRenderer>();
        gravity = GetComponent<Collider>();
        if (mainMenu.modeGame == 1) {
            modePlayed = playedModeIA;
            renderStone = renderStoneIA;
        } else {
            modePlayed = playedMode1vs1;
            renderStone = renderStone1vs1;
        }
    }

    async void playedMode1vs1() {
        if (await goban.GM.GetCheckRules(node, goban.GM.GetPlayerTurn())) {
            SetStone();
            goban.board.Add(transform.GetComponent<stone>());
        } else {
            Debug.Log("IMPOSSIBLE");
        }
    }
    async void playedModeIA() {
        if (goban.GM.GetPlayerTurn() == 1) {
            if (await goban.GM.GetCheckRules(node, goban.GM.GetPlayerTurn())) {
                SetStone();
                goban.board.Add(transform.GetComponent<stone>());
                goban.GM.GetPlayed(node);
            } else {
                Debug.Log("IMPOSSIBLE");
            }
        }
    }

    void renderStoneIA() {
        if (goban.GM.GetPlayerTurn() == 1) {
            rend.material = goban.GM.GetCurrentMaterial();
            meshRend.enabled = true;
        }
    }
    void renderStone1vs1() {
        rend.material = goban.GM.GetCurrentMaterial();
        meshRend.enabled = true;
    }
    void OnMouseDown() {
        if (!isCreate) {
            modePlayed();
        }
    }
    void OnMouseEnter() {
        if (!isCreate) {
            renderStone();
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
        Vector3 up = transform.position;
        up.y += 0.1f;
        transform.position = up;
        gravity.attachedRigidbody.useGravity = false;
    }
    public void SetStone() {
        rend.material = goban.GM.GetCurrentMaterial();
        node.Player = goban.GM.GetPlayerTurn();
        goban.GM.NextPlayer();
        Vector3 up = transform.position;
        up.y += 0.9f;
        transform.position = up;
        isCreate = true;
        gravity.attachedRigidbody.useGravity = true;
        meshRend.enabled = true;
    }

    public GomokuBuffer.Node GetNode() {
        return node;
    }
}
