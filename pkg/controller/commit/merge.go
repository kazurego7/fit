package commit

import (
	"fmt"
	"os"

	"fit/pkg/infra/git"
	"fit/pkg/service"
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var MergeCmd = &cobra.Command{
	Use:   "merge <commit>",
	Short: "指定したブランチを現在のブランチにマージする.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		if git.ExistsUntrackedFiles([]string{":/"}) || git.ExistsWorktreeDiff([]string{":/"}) || git.ExistsIndexDiff([]string{":/"}) {
			message := "インデックス・ワークツリーにファイルの変更があるため、マージを中止しました\n" +
				"※ \"fit stash store\" でファイルの変更をスタッシュに保存するか、\"fit change delete\" でファイルの変更を破棄してください"
			fmt.Fprintln(os.Stderr, message)
			return
		}
		gitSubCmd := []string{"merge", "--no-ff", args[0]}
		util.GitCommand(usecase.RootFlag, gitSubCmd)
	},
}
