package stash

import (
	"github.com/kazurego7/fit/fit/git"
	"github.com/spf13/cobra"
)

var StoreCmd = &cobra.Command{
	Use:   "store [<message>]",
	Short: "ワークツリーの変更をスタッシュとして保存する.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var message string
		if len(args) > 0 {
			message = args[0]
		}
		git.StashPushOnlyWorktree(message)
	},
}
