package tag

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <new tag name>",
	Short: "指定したコミットにタグを付ける.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"tag", args[0], createFlag.commit}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		gitSubCmd := []string{"tag", "--list", "--sort=-refname"}
		output, _, err := util.GitQuery(global.RootFlag, gitSubCmd)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return util.SplitLn(string(output)), cobra.ShellCompDirectiveNoFileComp
	},
}

var createFlag struct {
	commit string
}

func init() {
	CreateCmd.Flags().StringVarP(&createFlag.commit, "commit", "c", "HEAD", "タグをつけるコミット")
}
