using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class cameraScript : MonoBehaviour {
	public float RotationSpeed = 5.0f;
	public float SmootherFactor = 0.5f;
	private Transform target;
	private Vector3 _cameraOfflset;
	void Start() {
		target = GameObject.Find("goban").transform;
		transform.LookAt(target.position);
		_cameraOfflset = transform.position - target.position;
	}
	void Update() {
		if (Input.GetMouseButton(1)) {
			if (Input.GetAxis("Mouse X")!= 0 || Input.GetAxis("Mouse Y") != 0) {
				
				Quaternion camTurnAngle = Quaternion.AngleAxis(Input.GetAxis("Mouse X") * RotationSpeed, Vector3.up);
				_cameraOfflset = camTurnAngle * _cameraOfflset;
				Vector3 newPos = target.position + _cameraOfflset;
				transform.position = Vector3.Slerp(transform.position, newPos, SmootherFactor);
				transform.LookAt(target.position);
			}
		}
	}
}