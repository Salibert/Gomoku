using UnityEngine.UI;
using UnityEngine;
using UnityEngine.SceneManagement;
using System;
using System.Collections.Generic;
using GomokuBuffer;
using TMPro;
public class mainMenu : MonoBehaviour {


	static public int modeGame;
	static public GomokuBuffer.ConfigRules config;
	public Toggle m_captureToggle;
	public Toggle m_freeThreeToggle;
	public Toggle m_firstPlayerToggle;
	public Toggle m_helperPlayerToggle;
	public GameObject Difficulty;
	private TextMeshProUGUI textDifficulty;
	private Dictionary<string, int> difficulty;
	private String[] nameDifficulty;
	private int indexDifficulty;
	void Awake() {
		config = new GomokuBuffer.ConfigRules(){
			IsActiveRuleCapture = true,
			IsActiveRuleFreeThree = true,
			IsActiveRuleWin = true,
		};
		nameDifficulty = new String[]{ "easy","medium","hard","very hard","master"};
		difficulty = new Dictionary<string, int>();
		int lvl = 1;
		for (int key=0; key < nameDifficulty.Length; key++) {
			lvl += 2;
			difficulty.Add(nameDifficulty[key], lvl);
		};
		m_captureToggle.isOn = config.IsActiveRuleCapture;
		m_captureToggle.onValueChanged.AddListener(delegate { CaptureValueChanged(m_captureToggle); });
		m_freeThreeToggle.isOn = config.IsActiveRuleFreeThree;
		m_freeThreeToggle.onValueChanged.AddListener(delegate { FreeThreeValueChanged(m_freeThreeToggle); });
		m_firstPlayerToggle.isOn = config.PlayerIndexIA > 2 ? true : false;
		m_helperPlayerToggle.isOn = config.IsActiveHelperPlayer;
		m_helperPlayerToggle.onValueChanged.AddListener(delegate { HelperPlayerValueChanged(m_helperPlayerToggle); });
		textDifficulty = Difficulty.GetComponent<TextMeshProUGUI>();
	}

	public void PlayGame1VS1() {
		modeGame = 2;
		config.PlayerIndexIA = 0;
		SceneManager.LoadScene(SceneManager.GetActiveScene().buildIndex + 1);
	}

	public void PlayGameIA() {
		modeGame = 1;
		if (m_firstPlayerToggle.isOn == true) {
			config.PlayerIndexIA = 1;
		} else {
			config.PlayerIndexIA = 2;
		}
		config.DepthIA = difficulty[textDifficulty.text];
		SceneManager.LoadScene(SceneManager.GetActiveScene().buildIndex + 1);
	}
	public void upDifficulty() {
		indexDifficulty = (indexDifficulty == nameDifficulty.Length - 1) ? 0 : indexDifficulty + 1;
		textDifficulty.text = nameDifficulty[indexDifficulty];
	}
	public void downDifficulty() {
		indexDifficulty = (indexDifficulty == 0) ? nameDifficulty.Length - 1: indexDifficulty - 1;
		textDifficulty.text = nameDifficulty[indexDifficulty];
	}
    void CaptureValueChanged(Toggle change)
    {
		config.IsActiveRuleCapture = change.isOn;
    }
	void FreeThreeValueChanged(Toggle change)
    {
		config.IsActiveRuleFreeThree = change.isOn;
    }
	void HelperPlayerValueChanged(Toggle change)
    {
		config.IsActiveHelperPlayer = change.isOn;
    }
	public void QuitGame() {
		Application.Quit();
	}
}
