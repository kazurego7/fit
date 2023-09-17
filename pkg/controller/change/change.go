package change

import (
	"errors"
	"os"

	"fit/pkg/infra"
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var ChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "ワークツリー・インデックスの変更に関する操作.",
}

func init() {
	ChangeCmd.AddCommand(ListCmd)
	ChangeCmd.AddCommand(StageCmd)
	ChangeCmd.AddCommand(UnstageCmd)
	ChangeCmd.AddCommand(DeleteCmd)
	ChangeCmd.AddCommand(LogCmd)
	ChangeCmd.AddCommand(RestoreCmd)
}

func removeIndex(filenameList ...string) int {
	gitSubCmd := append([]string{"rm", "--cache", "--"}, filenameList...)
	return util.GitCommand(usecase.RootFlag, gitSubCmd...)
}

func restoreWorktree(filenameList ...string) int {
	gitSubCmd := append([]string{"restore", "--"}, filenameList...)
	return util.GitCommand(usecase.RootFlag, gitSubCmd...)
}

func restoreIndex(filenameList ...string) int {
	gitSubCmd := append([]string{"restore", "--staged", "--"}, filenameList...)
	return util.GitCommand(usecase.RootFlag, gitSubCmd...)
}

func clean(filenameList ...string) int {
	gitSubCmd := append([]string{"clean", "--force", "--"}, filenameList...)
	return util.GitCommand(usecase.RootFlag, gitSubCmd...)
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
		overwriteList := infra.SearchWorktreeList("", args...)
		if len(overwriteList) != 0 {
			return errors.New("復元するファイルに変更があります.\n" +
				"\"fit change delete\" で変更を削除するか、\"fit change stage\" でステージングを行ってください")
		}
		return nil
	}

}
