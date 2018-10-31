using UnityEngine;
using System.Collections;
using System.Collections.Generic;
public class goban : MonoBehaviour
{
    public Transform stonePrefab;
    public static gameMaster.node[][] board;
    private Transform inter;
    void Start()
    {
        board = new gameMaster.node[19][];
        Transform line = null;
        Transform lines = transform.Find("lines").transform;
        inter = transform.Find("stones");
        for (int i = 0; i < lines.childCount; i++)
        {
            line = lines.GetChild(i).transform;
            if (line.rotation.y == 0) {
                board[i] = findIntersections(line);
                createIntersection(ref board[i], i);
            }
        }
    }

    private void createIntersection(ref gameMaster.node[] lines, int x) {
        Transform newInstance;
        for (int i = 0; i < lines.Length; i++){
            lines[i].posArray.x = x;
            lines[i].posArray.y = i;
            newInstance = Instantiate(stonePrefab, lines[i].position, new Quaternion() { x=0, y=0, z=0 }, inter);
            newInstance.GetComponent<MeshRenderer>().enabled = false;
            newInstance.GetComponent<stone>().initNode(ref lines[i]);
        };
    }
    private gameMaster.node[] findIntersections(Transform line) {

        Vector3 fwd = transform.TransformDirection(transform.forward);
        List<gameMaster.node> nodes = new List<gameMaster.node>();
        List<RaycastHit> hitsAll = new List<RaycastHit>();

        hitsAll.AddRange(Physics.RaycastAll(line.position, fwd, 1000));
        hitsAll.AddRange(Physics.RaycastAll(line.position, fwd * -1, 1000));
        hitsAll.ForEach(el => {
            nodes.Add(new gameMaster.node() {
                position = new Vector3() { x = line.position.x, y = line.position.y + 1f, z = el.transform.position.z }
            });
        });
        nodes.Add(new gameMaster.node() {
            position = new Vector3() { x = line.position.x, y = line.position.y + 1f, z = line.position.z }
        });
        nodes.Sort(delegate(gameMaster.node x, gameMaster.node y) {
            if (x.position.z == y.position.z)
                return 0;
            return x.position.z > y.position.z ? 1 : -1;
        });
        return nodes.ToArray();
    }
}
