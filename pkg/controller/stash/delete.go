package stash

import (
	"errors"

	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete (<stash> | --all)",
	Short: "保存されたスタッシュを削除する.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var gitSubCmd []string
		if deleteFlag.all {
			gitSubCmd = []string{"stash", "clear"}
		} else {
			if len(args) == 0 {
				return errors.New("require 1 argument")
			}
			gitSubCmd = []string{"stash", "drop", args[0]}
		}
		util.GitCommand(global.RootFlag, gitSubCmd)
		return nil
	},
}

var deleteFlag struct {
	all bool
}

func init() {
	DeleteCmd.Flags().BoolVarP(&deleteFlag.all, "all", "a", false, "全スタッシュを削除する.")
}
