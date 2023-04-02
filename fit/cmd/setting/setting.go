package setting

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SettingCmd = &cobra.Command{
	Use:              "setting",
	TraverseChildren: true,
	Short:            "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	SettingCmd.AddCommand(CompletionCmd)
	SettingCmd.AddCommand(ListCmd)
	SettingCmd.AddCommand(SetCmd)
	SettingCmd.AddCommand(UnsetCmd)
	SettingCmd.AddCommand(EditCmd)
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
