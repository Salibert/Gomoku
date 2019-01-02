using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class player : MonoBehaviour {
        public Material material;
        public int index;
        public int score;

	public void SetScore(int newScore) {
		score = newScore;
	}

	public int GetScore() {
		return score;
	}

	public Material GetMaterial() {
		return material;
	}

	public int GetIndex() {
		return index;
	}
}
