package repository

import (
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "ローカルリポジトリを初期化する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		initGit()
		if existsHEADCommit() {
			return
		}
		firstCommit()
	},
}

func existsHEADCommit() bool {
	gitSubCmd := []string{"rev-parse", "HEAD"}
	_, exitCode, _ := util.GitQuery(usecase.RootFlag, gitSubCmd...)
	return exitCode == 0
}

func initGit() {
	gitSubCmd := []string{"init"}
	util.GitCommand(usecase.RootFlag, gitSubCmd...)
}

func firstCommit() {
	gitSubCmd := []string{"commit", "--allow-empty", "-m", "first commit"}
	util.GitCommand(usecase.RootFlag, gitSubCmd...)
}
