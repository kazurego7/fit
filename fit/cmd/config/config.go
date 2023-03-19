package config

import (
	"github.com/spf13/cobra"
)

var flags Flags

type Flags struct {
	local  bool
	global bool
	system bool
}

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
	ConfigCmd.PersistentFlags().BoolVar(&flags.local, "local", false, "local config")
	ConfigCmd.PersistentFlags().BoolVar(&flags.global, "global", false, "blobal config")
	ConfigCmd.PersistentFlags().BoolVar(&flags.system, "system", false, "system config")
	ConfigCmd.MarkFlagsMutuallyExclusive("local", "global", "system")
}

func getTargetScope(flags Flags) string {
	var targetScopeString string
	switch flags {
	case Flags{local: true, global: false, system: false}:
		targetScopeString = "--local"
	case Flags{local: false, global: true, system: false}:
		targetScopeString = "--global"
	case Flags{local: false, global: false, system: true}:
		targetScopeString = "--system"
	default:
		targetScopeString = "--local"
	}
	return targetScopeString
}
