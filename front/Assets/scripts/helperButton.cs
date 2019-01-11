using System.Collections;
using System.Collections.Generic;
using UnityEngine;
public class helperButton : MonoBehaviour {
    private bool lockHelper;

    void Start() {
        lockHelper = false;
    }
    async private void OnMouseDown(){
        if (lockHelper != true && goban.GM.GetPlayerIndexIA() != 0  && goban.GM.GetPlayerTurn() != goban.GM.GetPlayerIndexIA()) {
            lockHelper = true;
            GomokuBuffer.Node node = new GomokuBuffer.Node(){Player=goban.GM.GetPlayerTurn()};
            await goban.GM.GetPlayedHelp(node);
            lockHelper = false;
        }
    }

    public bool GetLockerHelpher() {
        return lockHelper;
    }
}