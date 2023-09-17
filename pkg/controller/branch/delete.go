package branch

import (
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <branch name>",
	Short: "指定したブランチを削除する(ブランチの指すコミットは削除しない).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"branch", "--delete", "--force", args[0]}
		util.GitCommand(usecase.RootFlag, gitSubCmd)
	},
}
