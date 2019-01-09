using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class managerLights : MonoBehaviour {
	public Light[] listHouseLights;
	public float graduatingHouseLights;
	public float IntensityHouse;
	public Light[] listGroundLights;
	public float graduatingGroundLights;
	public float IntensityGround;

	IEnumerator GradualnessStartLight(Light[] lights, float intensity, float graduating) {
		for (float indexIntensity=0;  indexIntensity <= intensity; indexIntensity += graduating) {
			for (int i=0; i < lights.Length; i++ ) {
				lights[i].intensity = indexIntensity;
			}
			yield return new WaitForSeconds(0.5f);
		}
	}
	IEnumerator GradualnessDownLight(Light[] lights, float intensity, float graduating) {
		for (float indexIntensity=intensity;  indexIntensity >= 0f; indexIntensity -= graduating) {
			for (int i=0; i < lights.Length; i++ ) {
				lights[i].intensity = indexIntensity;
			}
			yield return new WaitForSeconds(0.5f);
		}
	}
	public void StartListHouseLights() {
		StartCoroutine(GradualnessStartLight(listHouseLights, IntensityHouse, IntensityHouse));
	}

	public void StartListGroundLights() {
		StartCoroutine(GradualnessStartLight(listGroundLights, IntensityGround, IntensityGround));
	}
	public void DownListHouseLights() {
		StartCoroutine(GradualnessDownLight(listHouseLights, IntensityHouse, IntensityHouse));
	}

	public void DownListGroundLights() {
		StartCoroutine(GradualnessDownLight(listGroundLights, IntensityGround, IntensityGround));
	}

	public void SwitchGroundToHouse() {
		StartCoroutine(GradualnessDownLight(listGroundLights, IntensityGround, IntensityGround));
		StartCoroutine(GradualnessStartLight(listHouseLights, IntensityHouse, IntensityHouse));
	}
	public void SwitchHouseToGround() {
		StartCoroutine(GradualnessDownLight(listHouseLights, IntensityHouse, IntensityHouse));
		StartCoroutine(GradualnessStartLight(listGroundLights, IntensityGround, IntensityGround));
	}

}
