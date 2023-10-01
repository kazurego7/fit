package change

import (
	"fmt"
	"os"

	"github.com/kazurego7/fit/pkg/infra/git"
	"github.com/kazurego7/fit/pkg/service"

	"github.com/spf13/cobra"
)

var StageCmd = &cobra.Command{
	Use:   "stage <filename>…",
	Short: "ワークツリーのファイルの変更をインデックスにステージングする.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		// ファイル名からあいまい検索のパスを作成
		pathList := service.AddFuzzyParentPath(args)
		if stageFlag.details {
			git.DiffWorktree(pathList)
		} else {
			// コンフリクト解消していないファイルがあればステージングしない
			if err := service.CheckConflictResolved(pathList); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			service.StageChange(pathList)
		}
	},
	ValidArgs: service.GetUnstagingFileNameList(),
}

var stageFlag struct {
	details bool
}

func init() {
	StageCmd.Flags().BoolVarP(&stageFlag.details, "details", "d", false, "ファイルの変更の詳細を表示する")
}
