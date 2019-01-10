using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
public class pauseMenu : MonoBehaviour {

	public static bool GameIsPaused = false;
	public GameObject pauseMenuUI;

	void Awake() {
		pauseMenuUI.SetActive(false);
		Time.timeScale = 1f;
		GameIsPaused = false;
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
	IEnumerator LoadAsync() {
		AsyncOperation load = SceneManager.LoadSceneAsync("Menu");
		yield return load;
		SceneManager.UnloadSceneAsync("Game");
	}
	public void LoadMenu() {
		StartCoroutine(LoadAsync());
		goban.GM.GetCDGame(true);
	}
	public void QuitGame() {
		Application.Quit();
	}
}
