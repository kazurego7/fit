package stash

import (
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "保存されたスタッシュを一覧表示する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if listFlag.details == "" {
			gitSubCmd := []string{"reflog", "show", "--format=%C(03)%h%C(reset) %C(bold 1)%gD%C(reset) %C(bold 0)%s%C(reset)", "stash"}
			util.GitCommand(usecase.RootFlag, gitSubCmd)
		} else {
			gitSubCmd := []string{"stash", "show", "--stat", "--summary", "--patch", "--include-untracked", listFlag.details}
			util.GitCommand(usecase.RootFlag, gitSubCmd)
		}
	},
}

var listFlag struct {
	details string
}

func init() {
	ListCmd.Flags().StringVarP(&listFlag.details, "details", "d", "", "指定したスタッシュに含まれるファイルを表示する.")
}
