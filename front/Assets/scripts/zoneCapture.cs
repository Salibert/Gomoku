using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class zoneCapture : MonoBehaviour {
	public Transform stonePrefab;
	private Transform[] listStone;
	// Use this for initialization
	void Start () {
		listStone = new Transform[10];
		Transform inter = transform.Find("stones");
		Vector3 pos = transform.position;
		for ( int i =0; i < 10; i++) {
			listStone[i] = Instantiate(stonePrefab, new Vector3(){x=pos.x + i, y=pos.y + 0.5f, z=pos.z}, new Quaternion() { x=0, y=0, z=0 }, inter);
		}
	}
	
	// Update is called once per frame
	void Update () {
		
	}
}
