package change

import (
	"errors"
	"fmt"
	"os"

	"github.com/kazurego7/fit/fit/cmd/stash"
	"github.com/kazurego7/fit/fit/git"
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var ChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "ワークツリー・インデックスの変更に関する操作.",
}

func init() {
	ChangeCmd.AddCommand(UnstageCmd)
	ChangeCmd.AddCommand(StageCmd)
	ChangeCmd.AddCommand(DeleteCmd)
	ChangeCmd.AddCommand(ListCmd)
	ChangeCmd.AddCommand(LogCmd)
	ChangeCmd.AddCommand(RestoreCmd)
}

func removeIndex(filenameList ...string) int {
	gitSubCmd := append([]string{"rm", "--cache", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd...)
}

func restoreWorktree(filenameList ...string) int {
	gitSubCmd := append([]string{"restore", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd...)
}

func restoreIndex(filenameList ...string) int {
	gitSubCmd := append([]string{"restore", "--staged", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd...)
}

func clean(filenameList ...string) int {
	gitSubCmd := append([]string{"clean", "--force", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd...)
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

func existsFiles(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(n)(cmd, args); err != nil {
			return err
		}
		for i := 0; i < n; i++ {
			if f, err := os.Stat(args[i]); os.IsNotExist(err) || f.IsDir() {
				return errors.New("ファイルが存在しない、または対象がファイルではありません")
			}
		}
		return nil
	}
}

func existsWorktreeChanges() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		overwriteList := git.SearchWorktreeList("", args...)
		if len(overwriteList) != 0 {
			return errors.New("復元するファイルに変更があります.変更を削除するか、ステージングを行ってください")
		}
		return nil
	}

}
