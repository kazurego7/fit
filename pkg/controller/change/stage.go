package change

import (
	"fmt"
	"os"

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

		// コンフリクト解消していないファイルがあればステージングしない
		if err := service.CheckConflictResolved(pathList); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		service.StageChange(pathList)
	},
	ValidArgs: service.GetUnstagingFileNameList(),
}
