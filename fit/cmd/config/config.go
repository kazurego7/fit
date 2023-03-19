package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:              "config",
	TraverseChildren: true,
	Short:            "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	ConfigCmd.AddCommand(ListCmd)
	ConfigCmd.AddCommand(SetCmd)
	ConfigCmd.AddCommand(UnsetCmd)
	ConfigCmd.AddCommand(EditCmd)
	ConfigCmd.PersistentFlags().StringVar(&scopeFlag.arg, "scope", "user", `config scope from "repository", "user" or "local"`)
}

var scopeFlag ScopeFlag

type ScopeFlag struct {
	arg string
}

func (scopeFlag ScopeFlag) toGitFlag() (string, error) {
	switch scopeFlag.arg {
	case "repository":
		return "--local", nil
	case "user":
		return "--global", nil
	case "system":
		return "--system", nil
	default:
		return "", fmt.Errorf(`"%v" is invalid in "--scope" flag. use "repository", "user" or "system"`, scopeFlag.arg)
	}
}
