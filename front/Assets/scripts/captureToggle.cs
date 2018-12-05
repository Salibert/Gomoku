using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class captureToggle : MonoBehaviour {
    public Toggle m_Toggle;
    void Start()
    {
        m_Toggle.isOn = mainMenu.config.IsActiveRuleCapture;
        m_Toggle = GetComponent<Toggle>();
        m_Toggle.onValueChanged.AddListener(delegate {
                ToggleValueChanged(m_Toggle);
            });
    }

    void ToggleValueChanged(Toggle change)
    {
		mainMenu.config.IsActiveRuleCapture = m_Toggle.isOn;
    }
}
