package controller

import (
	"os"

	"github.com/kazurego7/fit/pkg/controller/branch"
	"github.com/kazurego7/fit/pkg/controller/change"
	"github.com/kazurego7/fit/pkg/controller/commit"
	"github.com/kazurego7/fit/pkg/controller/config"
	"github.com/kazurego7/fit/pkg/controller/conflict"
	"github.com/kazurego7/fit/pkg/controller/repository"
	"github.com/kazurego7/fit/pkg/controller/stash"
	"github.com/kazurego7/fit/pkg/controller/tag"
	"github.com/kazurego7/fit/pkg/global"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "fit",
	Short: "ユーザーフレンドリーな git CLI.",
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(change.ChangeCmd)
	RootCmd.AddCommand(commit.CommitCmd)
	RootCmd.AddCommand(conflict.ConflictCmd)
	RootCmd.AddCommand(branch.BranchCmd)
	RootCmd.AddCommand(tag.TagCmd)
	RootCmd.AddCommand(stash.StashCmd)
	RootCmd.AddCommand(repository.RepositoryCmd)
	RootCmd.AddCommand(config.ConfigCmd)

	RootCmd.PersistentFlags().BoolVar(&global.RootFlag.Dryrun, "dry-run", false, "データの変更なしに実行する.")
	RootCmd.PersistentFlags().BoolVar(&global.RootFlag.Debug, "debug", false, "fit内部で実行するgitコマンドを出力する.")

	cobra.EnableCommandSorting = false
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}
