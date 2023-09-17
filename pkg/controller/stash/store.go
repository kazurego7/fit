package stash

import (
	"fit/pkg/infra/git"
	"fit/pkg/service"

	"github.com/spf13/cobra"
)

var StoreCmd = &cobra.Command{
	Use:   "store [<message>]",
	Short: "ワークツリーの変更をスタッシュとして保存する.",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1), service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		var message string
		if len(args) > 0 {
			message = args[0]
		}
		git.StashPushOnlyWorktree(message)
	},
}
