using System.Collections;
using UnityEngine;

public class intersections : MonoBehaviour
{
    public Transform intersectionPrefab;
    public void createIntersection(int len, gameMaster.nodes[] lines) {
        for (int i = 0; i < len; i++)
            Instantiate(intersectionPrefab,
                new Vector3() { x=lines[i].x, y=transform.position.y, z=lines[i].y },
                new Quaternion() { x=0, y=0, z=0 });
    }
}
