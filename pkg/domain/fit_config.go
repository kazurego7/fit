package domain

import "regexp"

// 公開された設定設定
var config = fitConfig{
	mainline: setting{
		name:         "fit.mainline",
		defaultValue: "main",
	},
	workflow: setting{
		name:         "fit.workflow",
		defaultValue: "gitlab",
	},
}

func FitConfig() fitConfig {
	return config
}

type fitConfig struct {
	mainline Setting
	workflow Setting
}

func (s fitConfig) Validate(configName string) bool {
	isFitSetting := regexp.MustCompile(`^fit\..+`).MatchString(configName)
	if !isFitSetting {
		return true
	}
	constNames := []string{
		FitConfig().Mainline().Value(),
	}
	for _, constName := range constNames {
		if configName == constName {
			return true
		}
	}
	return false
}

func (s fitConfig) Mainline() Setting {
	return s.mainline
}

func (s fitConfig) Workflow() Setting {
	return s.workflow
}
