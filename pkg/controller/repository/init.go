package repository

import (
	"fit/pkg/infra/git"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "ローカルリポジトリを初期化する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		git.InitGit()
		if git.ExistsHEADCommit() {
			return
		}
		git.FirstCommit()
	},
}
