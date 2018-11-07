using UnityEngine;
using System.Collections;
using System.Collections.Generic;
public class goban : MonoBehaviour
{
    public Transform stonePrefab;
    public static gameMaster.node[][] board;
    private Transform inter;
    public static gameMaster currentGM;
    void Start()
    {
        currentGM = GameObject.Find("gameMaster").GetComponent<gameMaster>();
        board = new gameMaster.node[19][];
        Transform line = null;
        Transform lines = transform.Find("lines").transform;
        inter = transform.Find("stones");
        for (int i = 0; i < lines.childCount; i++)
        {
            line = lines.GetChild(i).transform;
            if (line.rotation.y == 0)
                board[i] = createStones(findIntersections(line), i);
        }
    }

    private gameMaster.node[] createStones(Vector3[] pos, int x) {
        Transform newInstance;
        gameMaster.node[] lines = new gameMaster.node[19];
        for (int i = 0; i < pos.Length; i++){
            lines[i].posArray.x = x;
            lines[i].posArray.y = i;
            newInstance = Instantiate(stonePrefab, pos[i], new Quaternion() { x=0, y=0, z=0 }, inter);
            newInstance.GetComponent<MeshRenderer>().enabled = false;
            newInstance.GetComponent<stone>().initNode(ref lines[i]);
        };
        return lines;
    }
    private Vector3[] findIntersections(Transform line) {

        Vector3 fwd = transform.TransformDirection(transform.forward);
        List<Vector3> pos = new List<Vector3>();
        List<RaycastHit> hitsAll = new List<RaycastHit>();

        hitsAll.AddRange(Physics.RaycastAll(line.position, fwd, 1000));
        hitsAll.AddRange(Physics.RaycastAll(line.position, fwd * -1, 1000));
        hitsAll.ForEach(el => {
            pos.Add(new Vector3() { x = line.position.x, y = line.position.y + 1f, z = el.transform.position.z });
        });
        pos.Add(new Vector3() { x = line.position.x, y = line.position.y + 1f, z = line.position.z });
        pos.Sort(delegate(Vector3 x, Vector3 y) {
            if (x.z == y.z)
                return 0;
            return x.z > y.z ? 1 : -1;
        });
        return pos.ToArray();
    }
}