package setting

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SettingCmd = &cobra.Command{
	Use:              "setting",
	TraverseChildren: true,
	Short:            "gitの設定に関する操作.",
}

func init() {
	SettingCmd.AddCommand(ListCmd)
	SettingCmd.AddCommand(SetCmd)
	SettingCmd.AddCommand(UnsetCmd)
	SettingCmd.AddCommand(EditCmd)
	SettingCmd.AddCommand(CompletionCmd)

	SettingCmd.PersistentFlags().BoolVar(&settingFlag.local, "local", false, "config scope local")
	SettingCmd.PersistentFlags().BoolVar(&settingFlag.user, "user", true, "config scope user")
	SettingCmd.PersistentFlags().BoolVar(&settingFlag.system, "system", false, "config scope system")
	SettingCmd.MarkFlagsMutuallyExclusive("local", "user", "system")
}

var settingFlag SettingFlag

type SettingFlag struct {
	local  bool
	user   bool
	system bool
}

func getScopeFlag() string {
	fmt.Println(settingFlag)
	switch {
	case settingFlag.local:
		return "--local"
	case settingFlag.system:
		return "--system"
	default:
		return "--global"
	}
}
