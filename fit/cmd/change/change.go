package change

import (
	"fmt"

	"github.com/kazurego7/fit/fit/cmd/stash"
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var ChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	ChangeCmd.AddCommand(DiffCmd)
	ChangeCmd.AddCommand(UnstageCmd)
	ChangeCmd.AddCommand(StageCmd)
	ChangeCmd.AddCommand(ClearCmd)
	ChangeCmd.AddCommand(ListCmd)
}

func searchIndexList(diffFilter string, filenameList ...string) []string {
	if len(filenameList) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--relative", "--staged", "--diff-filter=" + diffFilter, "--"}, filenameList...)
	out, _, _ := util.GitQuery(gitSubCmd...)
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	return util.SplitLn(string(out))
}

func searchWorktreeList(diffFilter string, filenameList ...string) []string {
	if len(filenameList) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--relative", "--diff-filter=" + diffFilter, "--"}, filenameList...)
	out, _, _ := util.GitQuery(gitSubCmd...)
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	return util.SplitLn(string(out))
}

func restoreWorktree(filenameList ...string) int {
	gitSubCmd := append([]string{"restore", "--"}, filenameList...)
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	return util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
}

func restoreIndex(filenameList ...string) int {
	gitSubCmd := append([]string{"restore", "--staged", "--"}, filenameList...)
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	return util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
}

func clean(filenameList ...string) int {
	gitSubCmd := append([]string{"clean", "--force", "--"}, filenameList...)
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	return util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
}

func confirmBackup() {
	fmt.Print("Includes overwrite operations on worktree and index. \nDo you want to create backups? [yes/no]: ")
	answer := util.InputYesOrNo(false)
	if answer {
		stash.Snap("fit auto backup")
	}
}
