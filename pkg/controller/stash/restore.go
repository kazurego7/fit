package stash

import (
	"github.com/kazurego7/fit/pkg/infra/git"
	"github.com/kazurego7/fit/pkg/service"

	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore [<stash number> | <stash id>]",
	Short: "保存されたスタッシュをワークツリー・インデックスに復元する.",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1), service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {

		var stashcommit string
		if len(args) == 0 {
			stashcommit = `stash@{0}`
		} else {
			stashcommit = args[0]
		}

		exitCode := git.ApplyKeepIndex(stashcommit)
		if exitCode != 0 {
			git.Apply(stashcommit)
		}
	},
}
