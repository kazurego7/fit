package tag

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use:   "upload <tag>",
	Short: "リモートリポジトリにタグをアップロードする(警告：アップロードされたタグは削除できません).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"push", "origin", "tags/" + args[0], "--prune"}
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
