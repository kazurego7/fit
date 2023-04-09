package revision

import (
	"fmt"
	"os"

	"github.com/kazurego7/fit/fit/git"
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var MergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "指定したブランチを現在のブランチにマージする.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repositoryPath := git.GetRepositoryPath()
		if git.ExistsUntrackedFiles(repositoryPath) || git.ExistsWorktreeDiff(repositoryPath) || git.ExistsIndexDiff(repositoryPath) {
			message := "インデックス・ワークツリーにファイルの変更があるため、マージを中止しました\n" +
				"※ \"fit stash store\" でファイルの変更をスタッシュに保存するか、\"fit change delete\" でファイルの変更を破棄してください"
			fmt.Fprintln(os.Stderr, message)
			return
		}
		gitSubCmd := []string{"merge", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
