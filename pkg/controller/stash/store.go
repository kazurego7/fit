package stash

import (
	service "github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/infra/git"

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
