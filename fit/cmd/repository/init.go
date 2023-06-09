package repository

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
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
	_, exitCode, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	return exitCode == 0
}

func initGit() {
	gitSubCmd := []string{"init"}
	util.GitCommand(global.RootFlag, gitSubCmd...)
}

func firstCommit() {
	gitSubCmd := []string{"commit", "--allow-empty", "-m", "first commit"}
	util.GitCommand(global.RootFlag, gitSubCmd...)
}
