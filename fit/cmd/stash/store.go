package stash

import (
	"github.com/spf13/cobra"
)

var StoreCmd = &cobra.Command{
	Use:   "store [<message>]",
	Short: "ワークツリー・インデックスの変更をスタッシュとして保存する.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var message string
		if len(args) > 0 {
			message = args[0]
		}
		switch {
		case storeFlags.worktree && storeFlags.index || !storeFlags.worktree && !storeFlags.index:
			stashPushAll(message)
		case storeFlags.index:
			stashPushOnlyWorktree(message)
		case storeFlags.worktree:
			stashPushOnlyIndex(message)
		}
	},
}

var storeFlags struct {
	worktree bool
	index    bool
}

func init() {
	StoreCmd.Flags().BoolVarP(&storeFlags.worktree, "worktree", "w", false, "ワークツリーの変更だけをスタッシュとして保存する.")
	StoreCmd.Flags().BoolVarP(&storeFlags.index, "index", "i", false, "インデックスの変更だけをスタッシュとして保存する.")
	StoreCmd.MarkFlagsMutuallyExclusive("worktree", "index")
}
