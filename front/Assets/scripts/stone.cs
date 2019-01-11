using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Threading.Tasks;
using GomokuBuffer;

public class stone : MonoBehaviour
{
    private MeshRenderer meshRend;
    public int X;
    public int Y;
    private Collider gravity;
    private Renderer rend;
    private GomokuBuffer.Node node;
    private bool isCreate;

    delegate void Played();
    Played modePlayed;

    delegate void Render();
    Render renderStone;
    public void initNode(ref GomokuBuffer.Node n) { node = n; X = n.X; Y = n.Y; }

    void Awake() {
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
            if (goban.GM.GetPlayerTurn() == 1) {
                goban.GM.manager.SwitchHouseToGround();
            } else {
                goban.GM.manager.SwitchGroundToHouse();
            }
            goban.board.Add(transform.GetComponent<stone>());
        } else {
            Debug.Log("IMPOSSIBLE");
        }
    }
    async void playedModeIA() {
        if (goban.GM.GetPlayerTurn() != goban.GM.GetPlayerIndexIA()) {
            if (await goban.GM.GetCheckRules(node, goban.GM.GetPlayerTurn())) {
                SetStone();
                goban.board.Add(transform.GetComponent<stone>());
                if (!goban.GM.GetGameIsFinish()) {
                    goban.GM.manager.SwitchGroundToHouse();
                    await goban.GM.GetPlayed(node);
                    goban.GM.manager.SwitchHouseToGround();
                }
            } else {
                Debug.Log("IMPOSSIBLE");
            }
        }
    }

    void renderStoneIA() {
        if (goban.GM.GetPlayerTurn() != goban.GM.GetPlayerIndexIA()) {
            rend.material = goban.GM.GetCurrentMaterial();
            meshRend.enabled = true;
        }
    }
    void renderStone1vs1() {
        rend.material = goban.GM.GetCurrentMaterial();
        meshRend.enabled = true;
    }
    void OnMouseDown() {
        if (!isCreate && !pauseMenu.GameIsPaused && !goban.GM.helper.GetLockerHelpher()) {
            modePlayed();
        }
    }
    void OnMouseEnter() {
        if (!isCreate && !pauseMenu.GameIsPaused && !goban.GM.helper.GetLockerHelpher()) {
            renderStone();
        }
    }

    void OnMouseExit() {
        if (!isCreate && !pauseMenu.GameIsPaused && !goban.GM.helper.GetLockerHelpher()) {
            meshRend.enabled = false;
        }
    }

    public void Reset() {
        meshRend.enabled = false;
        node.Player = 0;
        isCreate = false;
        Vector3 up = transform.position;
        up.y += 0.2f;
        transform.position = up;
        gravity.attachedRigidbody.useGravity = true;
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

    public void SetMaterial(Material material) {
        rend.material = material;
    }
    public GomokuBuffer.Node GetNode() {
        return node;
    }
}
