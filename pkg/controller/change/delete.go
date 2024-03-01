package change

import (
	"fmt"
	"os"

	service "github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/infra/git"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <pathspec>…",
	Short: "ワークツリー・インデックスの変更を削除する.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case deleteFlag.worktree && deleteFlag.index || !deleteFlag.worktree && !deleteFlag.index:
			if !git.ExistsUntrackedFiles(args) && !git.ExistsWorktreeDiff(args) && !git.ExistsIndexDiff(args) {
				fmt.Fprintln(os.Stderr, "削除するファイルがありません")
				return
			}
			service.BackupDelete(args)
			service.DeleteAll(args)
		case deleteFlag.worktree:
			if !git.ExistsUntrackedFiles(args) && !git.ExistsWorktreeDiff(args) {
				fmt.Fprintln(os.Stderr, "削除するファイルがありません")
				return
			}
			service.BackupDelete(args)
			service.DeleteWorktree(args)
		case deleteFlag.index:
			if !git.ExistsIndexDiff(args) {
				fmt.Fprintln(os.Stderr, "削除するファイルがありません")
				return
			}
			service.BackupDelete(args)
			service.DeleteIndex(args)
		}
	},
}

var deleteFlag struct {
	worktree bool
	index    bool
}

func init() {
	DeleteCmd.Flags().BoolVarP(&deleteFlag.worktree, "worktree", "w", false, "ワークツリーの変更を削除する.")
	DeleteCmd.Flags().BoolVarP(&deleteFlag.index, "index", "i", false, "インデックスの変更を削除する.")
	DeleteCmd.MarkFlagsMutuallyExclusive("worktree", "index")
}
