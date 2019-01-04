using System.Collections;
using System.Collections.Generic;
using UnityEngine;
public class helperButton : MonoBehaviour {
    private void OnMouseDown(){
        GomokuBuffer.Node node = new GomokuBuffer.Node();
        goban.GM.GetPlayed(node);
        Debug.Log(name + " Game Object Clicked!");
    }
}