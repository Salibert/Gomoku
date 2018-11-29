using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class cameraScript : MonoBehaviour {
	public float RotationSpeedX = 2.0f;
	public float RotationSpeedY = 0.5f;
	public float SmootherFactor = 0.5f;
	private Transform target;
	private Vector3 _cameraOfflset;
	public float minFov = 30f;
	public float maxFov = 90f;
	public float sensitivity = 10f;
	public Vector3 tmp;
	public Vector3 tmp1;
	void Start() {
		target = GameObject.Find("goban").transform;
		transform.LookAt(target.position);
		_cameraOfflset = transform.position - target.position;
	}

	private void HandleMouseButton() {
		if (Input.GetMouseButton(1) && Input.GetAxis("Mouse X")!= 0) {
			_cameraOfflset = Quaternion.AngleAxis(Input.GetAxis("Mouse X") * RotationSpeedX, Vector3.up)* _cameraOfflset;
			Vector3 newPos = target.position + _cameraOfflset;
			transform.position = Vector3.Slerp(transform.position, newPos, SmootherFactor);
			transform.LookAt(target.position);
		}
	}

	private void HandleKeyBoard() {
		 if (Input.GetKey(KeyCode.UpArrow)){
			if (transform.position.y < 30f) {
				_cameraOfflset = Quaternion.AngleAxis(RotationSpeedY, Vector3.right) * _cameraOfflset;
				Vector3 newPos = target.position + _cameraOfflset;
				transform.position = Vector3.Slerp(transform.position, newPos, SmootherFactor);
				transform.LookAt(target.position);
			}
		} else if (Input.GetKey(KeyCode.DownArrow)) {
			if (transform.position.y > 10f) {
				_cameraOfflset = Quaternion.AngleAxis(RotationSpeedY * -1, Vector3.right) * _cameraOfflset;
				Vector3 newPos = target.position + _cameraOfflset;
				transform.position = Vector3.Slerp(transform.position, newPos, SmootherFactor);
				transform.LookAt(target.position);
			}
		}
	}

	private void HandleScrollWhell() {
		if (Input.GetAxis("Mouse ScrollWheel") != 0f ) {
			transform.LookAt(target);
			float fov = Camera.main.fieldOfView;
			fov += Input.GetAxis("Mouse ScrollWheel") * sensitivity; 
			fov = Mathf.Clamp(fov, minFov, maxFov);
			Camera.main.fieldOfView = fov;
		}
	}
	void Update() {
		HandleMouseButton();
		HandleKeyBoard();
		HandleScrollWhell();
	}
}