using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class StoneOrbit : MonoBehaviour {
	public Transform stone1;
	public Transform stone2;
	public Transform Cam;

	void Start () {
		stone1.LookAt(Cam.position, Vector3.down);
		stone2.LookAt(Cam.position, Vector3.down);
	}

	public void SwitchAnimation(int target) {
			stone1.GetComponent<Animation>().enabled = true;
			stone2.GetComponent<Animation>().enabled = true;
	}
}
