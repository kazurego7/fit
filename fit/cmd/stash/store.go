package stash

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var StoreCmd = &cobra.Command{
	Use:   "store",
	Short: "ワークツリーやインデックスの変更をスタッシュとして保存する.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var gitSubCmd []string
		switch {
		case storeFlags.worktree && storeFlags.index || !storeFlags.worktree && !storeFlags.index:
			gitSubCmd = []string{"stash", "push", "--include-untracked"}
		case storeFlags.index:
			gitSubCmd = []string{"stash", "push", "--staged"}
		case storeFlags.worktree:
			gitSubCmd = []string{"stash", "push", "--include-untracked", "--keep-index"}
		}
		// メッセージがあれば追加
		if len(args) != 0 {
			gitSubCmd = append(gitSubCmd, args[0])
		}
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
		return nil
	},
}

var storeFlags struct {
	worktree bool
	index    bool
	all      bool
}

func init() {
	StoreCmd.Flags().BoolVarP(&storeFlags.worktree, "worktree", "w", false, "ワークツリーの変更だけをスタッシュとして保存する.")
	StoreCmd.Flags().BoolVarP(&storeFlags.index, "index", "i", false, "インデックスの変更だけをスタッシュとして保存する.")
	StoreCmd.MarkFlagsMutuallyExclusive("worktree", "index")
}
