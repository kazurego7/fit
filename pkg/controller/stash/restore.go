package stash

import (
	"fit/pkg/infra"
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore [<stash number> | <stash id>]",
	Short: "保存されたスタッシュをワークツリー・インデックスに復元する.",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1), infra.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {

		var stashcommit string
		if len(args) == 0 {
			stashcommit = `stash@{0}`
		} else {
			stashcommit = args[0]
		}

		exitCode := applyKeepIndex(stashcommit)
		if exitCode != 0 {
			apply(stashcommit)
		}
	},
}

func applyKeepIndex(stashcommit string) int {
	gitSubCmd := []string{"stash", "apply", "--quiet", "--index", stashcommit}
	return util.GitCommand(usecase.RootFlag, gitSubCmd...)
}

func apply(stashcommit string) int {
	gitSubCmd := []string{"stash", "apply", "--quiet", stashcommit}
	return util.GitCommand(usecase.RootFlag, gitSubCmd...)
}
