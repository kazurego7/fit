package domain

import "regexp"

// 公開された設定設定
var (
	FitSetting = fitSetting{mainlineType: "fit.mainline", mainleieDefaultValue: "main"}
)

type fitSetting struct {
	mainlineType         string
	mainleieDefaultValue string
}

func (s fitSetting) MainlineType() string {
	return s.mainlineType
}

func (s fitSetting) MainlineDefaultValue() string {
	return s.mainlineType
}

func ValidateFitSetting(configName string) bool {
	isFitSetting := regexp.MustCompile(`^fit\..+`).MatchString(configName)
	if !isFitSetting {
		return true
	}
	constNames := []string{
		FitSetting.MainlineDefaultValue(),
	}
	for _, constName := range constNames {
		if configName == constName {
			return true
		}
	}
	return false
}
