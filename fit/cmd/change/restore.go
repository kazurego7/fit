package change

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore <revision>",
	Short: "リビジョンに含まれるファイルをワークツリーに復元する.",
	Args:  cobra.MatchAll(existsFiles(1), existsWorktreeChanges()),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"restore", "--source", RestoreFlag.revision, args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}

var RestoreFlag struct {
	revision string
	before   bool
	after    bool
}

func init() {
	RestoreCmd.Flags().StringVarP(&RestoreFlag.revision, "revision", "r", "HEAD", "gitリビジョンを指定する.")
}
