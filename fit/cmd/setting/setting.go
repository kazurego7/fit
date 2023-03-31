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
	SettingCmd.PersistentFlags().StringVar(&scopeFlag.arg, "scope", "user", `config scope from "local", "user" or "system"`)
}

var scopeFlag ScopeFlag

type ScopeFlag struct {
	arg string
}

func (scopeFlag ScopeFlag) toGitFlag() (string, error) {
	switch scopeFlag.arg {
	case "local":
		return "--local", nil
	case "user":
		return "--global", nil
	case "system":
		return "--system", nil
	default:
		return "", fmt.Errorf(`"%v" is invalid in "--scope" flag. use "local", "user" or "system"`, scopeFlag.arg)
	}
}
