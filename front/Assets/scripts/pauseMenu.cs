using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
public class pauseMenu : MonoBehaviour {

	public static bool GameIsPaused = false;
	public GameObject pauseMenuUI;

	void Awake() {
		pauseMenuUI.SetActive(false);
	}
	void Update () {
		if (Input.GetKeyDown(KeyCode.Escape)) {
			if (GameIsPaused) {
				Resume();
			} else {
				Pause();
			}
		}
	}


	public void Resume() {
		GameIsPaused = false;
		pauseMenuUI.SetActive(false);
		Time.timeScale = 1f;
	}

	public void Pause() {
		pauseMenuUI.SetActive(true);
		Time.timeScale = 0f;
		GameIsPaused = true;
	}
	public void LoadMenu() {
		SceneManager.LoadScene("Menu");
	}
	public void QuitGame() {
		Debug.Log("QUIT");
		Application.Quit();
	}
}
