package commit

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "コミットを一覧表示する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if listFlag.details == "" {
			git.FetchPrune()

			gitSubCmd := []string{
				"-c",
				"core.pager=less -FRXS",
				"log",
				"--branches",
				"--tags",
				"--remotes",
				"--graph",
				"--abbrev-commit",
				"--decorate=no",
				"--date=format: %Y-%m-%d %H:%I:%S",
				"--format=format: %C(03)%>|(26)%h%C(reset)   %C(bold 1)%d%C(reset) %C(bold 0)%s%C(reset) %>|(140)%C(reset)  %C(04)%ad%C(reset)  %C(green)%<(16,trunc)%an%C(reset)",
			}
			if listFlag.orphan {
				gitSubCmd = append(gitSubCmd, "--reflog")
			}
			if listFlag.stash {
				gitSubCmd = append(gitSubCmd, "--all")
			}
			util.GitCommand(global.RootFlag, gitSubCmd)
		} else {
			gitSubCmd := []string{"show", "--stat", "--summary", "--patch", listFlag.details}
			util.GitCommand(global.RootFlag, gitSubCmd)
		}
	},
}

var listFlag struct {
	orphan  bool
	stash   bool
	details string
}

func init() {
	ListCmd.Flags().BoolVarP(&listFlag.orphan, "orphan", "o", false, "ブランチ・タグの付いていない孤独なコミットもログに表示する(スタッシュも表示される).")
	ListCmd.Flags().BoolVarP(&listFlag.stash, "stash", "s", false, "スタッシュもログに表示する.")
	ListCmd.Flags().StringVarP(&listFlag.details, "details", "d", "", "指定したコミットに含まれるファイルを表示する.")
}
