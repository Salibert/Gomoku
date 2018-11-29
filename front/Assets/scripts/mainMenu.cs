using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
public class mainMenu : MonoBehaviour {

	static public int modeGame;
	public void PlayGame1VS1() {
		modeGame = 2;
		SceneManager.LoadScene(SceneManager.GetActiveScene().buildIndex + 1);
	}

	public void PlayGameAI() {
		modeGame = 1;
		SceneManager.LoadScene(SceneManager.GetActiveScene().buildIndex + 1);
	}

	public void QuitGame() {
		Debug.Log("QUIT");
		Application.Quit();
	}
}
