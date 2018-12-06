using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using GomokuBuffer;
public class mainMenu : MonoBehaviour {


	static public int modeGame;
	static public GomokuBuffer.ConfigRules config;
	
	void Awake() {
		config = new GomokuBuffer.ConfigRules(){
			IsActiveRuleCapture = true,
			IsActiveRuleFreeThree = true,
			IsActiveRuleAlignment = true,
		};
	}

	public void PlayGame1VS1() {
		modeGame = 2;
		SceneManager.LoadScene(SceneManager.GetActiveScene().buildIndex + 1);
	}

	public void PlayGameIA() {
		modeGame = 1;
		SceneManager.LoadScene(SceneManager.GetActiveScene().buildIndex + 1);
	}

	public void QuitGame() {
		Debug.Log("QUIT");
		Application.Quit();
	}
}
