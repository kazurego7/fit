package change

import (
	"fmt"
	"os"

	"github.com/kazurego7/fit/pkg/infra/git"
	"github.com/kazurego7/fit/pkg/service"

	"github.com/spf13/cobra"
)

var StageCmd = &cobra.Command{
	Use:   "stage <pathspec>…",
	Short: "ワークツリーのファイルの変更をインデックスにステージングする.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		// コンフリクト解消していないファイルがあればステージングしない
		if err := service.CheckConflictResolved(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		service.StageChange(args)
	},
	ValidArgs: append(git.SearchUntrackedFiles([]string{":/"}), git.SearchWorktreeList("u", []string{":/"})...),
}
