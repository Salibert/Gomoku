using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class zoneCapture : MonoBehaviour {
	public Transform stonePrefab;
	private Transform[] listStone;
	public void SetZoneCapture(player player) {
		listStone = new Transform[10];
		Transform inter = transform.Find("stones");
		Vector3 pos = transform.position;
		float offsetX = pos.x - (transform.localScale.x / 2) + 2.25f;
		float offsetY = pos.y + (transform.localScale.y / 2) + 0.25f;
		for ( int i =0; i < 10; i++) {
			listStone[i] = Instantiate(stonePrefab, new Vector3(){x= offsetX+ (i * 1.5f), y=pos.y + offsetY, z=pos.z}, new Quaternion() { x=0, y=0, z=0 }, inter);
			stone script = listStone[i].GetComponent<stone>();
			script.SetMaterial(player.GetMaterial());
			Destroy(listStone[i].GetComponent<stone>());
			listStone[i].GetComponent<MeshRenderer>().enabled = false;
			listStone[i].GetComponent<Collider>().attachedRigidbody.useGravity = true;
		}
	}
	IEnumerator animeStone(int score) {
		Transform stone;
		for (int i = 0; i < score; i++) {
			stone = listStone[i];
			stone.GetComponent<MeshRenderer>().enabled = true;
			Vector3 up = stone.position;
        	up.y += 0.9f;
			stone.position = up;
			yield return new WaitForSeconds(0.25f);
		}
		yield return null;
    }

	public void AddStone(int score) {
		StartCoroutine(animeStone(score));
	}
}
