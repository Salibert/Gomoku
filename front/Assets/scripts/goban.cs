using UnityEngine;
using System.Collections;
using System.Collections.Generic;
public class goban : MonoBehaviour
{
    public gameMaster.nodes[][] board;
    void Start()
    {
        board = new gameMaster.nodes[19][];
        Transform line = null;
        Transform lines = transform.Find("lines").transform;
        intersections scriptInter = transform.Find("intersections").GetComponent<intersections>();
        for (int i = 0; i < lines.childCount; i++)
        {
            line = lines.GetChild(i).transform;
            if (line.rotation.y == 0) {
                board[i] = findIntersections(line);
                scriptInter.createIntersection(board[i].Length, board[i]);
            }
        }
    }

    private gameMaster.nodes[] findIntersections(Transform line) {
        Vector3 fwd = transform.TransformDirection(transform.forward);
        List<gameMaster.nodes> nodes = new List<gameMaster.nodes>();
        List<RaycastHit> hitsAll = new List<RaycastHit>();

        hitsAll.AddRange(Physics.RaycastAll(line.position, fwd, 1000));
        hitsAll.AddRange(Physics.RaycastAll(line.position, fwd * -1, 1000));
        // hitsAll.Sort(el => {  })
        hitsAll.ForEach(el => {
            nodes.Add(new gameMaster.nodes() {
                x = line.position.x,
                y = el.transform.position.z,
                player = 0,
                captered = false });
        });
        return nodes.ToArray();
    }


}
