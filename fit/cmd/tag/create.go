package tag

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <new tag name>",
	Short: "指定したコミットにタグを付ける.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"tag", args[0], createFlag.commit}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}

var createFlag struct {
	commit string
}

func init() {
	CreateCmd.Flags().StringVarP(&createFlag.commit, "commit", "c", "HEAD", "タグをつけるコミット")
}
