package config

import (
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:              "config",
	TraverseChildren: true,
	Short:            "gitの設定に関する操作.",
}

func init() {
	ConfigCmd.AddCommand(ListCmd)
	ConfigCmd.AddCommand(SetCmd)
	ConfigCmd.AddCommand(UnsetCmd)
	ConfigCmd.AddCommand(EditCmd)
	ConfigCmd.AddCommand(CompletionCmd)

	ConfigCmd.PersistentFlags().BoolVar(&configFlag.local, "local", false, "ローカルリポジトリのコンフィグ設定")
	ConfigCmd.PersistentFlags().BoolVar(&configFlag.user, "user", true, "ユーザー指定のコンフィグ設定")
	ConfigCmd.PersistentFlags().BoolVar(&configFlag.system, "system", false, "システム全体のコンフィグ設定")
	ConfigCmd.MarkFlagsMutuallyExclusive("local", "user", "system")
}

var configFlag ConfigFlag

type ConfigFlag struct {
	local  bool
	user   bool
	system bool
}

func getScopeFlag() string {
	switch {
	case configFlag.local:
		return "--local"
	case configFlag.system:
		return "--system"
	default:
		return "--global"
	}
}
