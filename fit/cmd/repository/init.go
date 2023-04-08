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
	_, exitCode, _ := util.GitQuery(gitSubCmd...)
	return exitCode == 0
}

func initGit() {
	gitSubCmd := []string{"init"}
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
}

func firstCommit() {
	gitSubCmd := []string{"commit", "--allow-empty", "-m", "first commit"}
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
}
