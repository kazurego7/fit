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
	Short: "ワークツリー・インデックスの変更に関する操作.",
}

func init() {
	ChangeCmd.AddCommand(ShowCmd)
	ChangeCmd.AddCommand(UnstageCmd)
	ChangeCmd.AddCommand(StageCmd)
	ChangeCmd.AddCommand(DeleteCmd)
	ChangeCmd.AddCommand(ListCmd)
}

func searchIndexList(diffFilter string, filenameList ...string) []string {
	if len(filenameList) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--relative", "--staged", "--no-renames", "--diff-filter=" + diffFilter, "--"}, filenameList...)
	out, _, _ := util.GitQuery(gitSubCmd...)
	return util.SplitLn(string(out))
}

func searchWorktreeList(diffFilter string, filenameList ...string) []string {
	if len(filenameList) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--relative", "--no-renames", "--diff-filter=" + diffFilter, "--"}, filenameList...)
	out, _, _ := util.GitQuery(gitSubCmd...)
	return util.SplitLn(string(out))
}

func removeIndex(filenameList ...string) int {
	gitSubCmd := append([]string{"rm", "--cache", "--"}, filenameList...)
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	return util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
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
	fmt.Print("ワークツリー・インデックスに対する上書き操作を含んでいます。\nファイルのバックアップをスタッシュに保存しますか？ [yes/no]: ")
	answer, err := util.InputYesOrNo(false)
	if err != nil {
		return
	}
	if answer {
		stash.Snap("fit 自動バックアップ")
	}
}
