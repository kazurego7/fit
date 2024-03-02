package repository

import (
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "ローカルリポジトリを初期化する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		git.InitRepository()
		git.SetConfigDefaultMainline()
		if git.ExistsHEADCommit() {
			return
		}
		git.FirstCommit()
	},
}
