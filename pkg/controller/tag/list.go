package tag

import (
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "タグを一覧表示する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"tag", "--list", "--sort=-version:refname"}
		util.GitCommand(usecase.RootFlag, gitSubCmd)
	},
}
