package conflict

import (
	"fmt"
	"os"

	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var ResolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "ファイルのマージコンフリクトを解消し、ステージングする.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkCmd := []string{"-c", "core.whitespace=-trailing-space,-space-before-tab,-indent-with-non-tab,-tab-in-indent,-cr-at-eol", "diff", "--check", args[0]}
		out, _, _ := util.GitQuery(checkCmd...)
		if string(out) != "" {
			fmt.Fprintln(os.Stderr, "コンフリクトマーカーが残っています.コンフリクトマーカーを取り除いてください.")
			return
		} else {
			gitSubCmd := []string{"add", args[0]}
			util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
			util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
			return
		}

	},
}
