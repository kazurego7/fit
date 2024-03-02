package domain

// 公開されたコンフィグ設定
var (
	FitConfigConstant = fitConfigConstant{mainlineType: "fit.mainline", mainleieDefaultValue: "main"}
)

type fitConfigConstant struct {
	mainlineType         string
	mainleieDefaultValue string
}

func (config fitConfigConstant) MainlineType() string {
	return config.mainlineType
}

func (config fitConfigConstant) MainlineDefaultValue() string {
	return config.mainlineType
}
