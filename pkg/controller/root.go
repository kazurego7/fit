package controller

import (
	"os"

	"fit/pkg/controller/branch"
	"fit/pkg/controller/change"
	"fit/pkg/controller/commit"
	"fit/pkg/controller/config"
	"fit/pkg/controller/conflict"
	"fit/pkg/controller/repository"
	"fit/pkg/controller/stash"
	"fit/pkg/controller/tag"
	"fit/pkg/usecase"

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
	RootCmd.AddCommand(stash.StashCmd)
	RootCmd.AddCommand(tag.TagCmd)
	RootCmd.AddCommand(repository.RepositoryCmd)
	RootCmd.AddCommand(config.ConfigCmd)

	RootCmd.PersistentFlags().BoolVar(&usecase.RootFlag.Dryrun, "dry-run", false, "データの変更なしに実行する.")
	RootCmd.PersistentFlags().BoolVar(&usecase.RootFlag.Debug, "debug", false, "fit内部で実行するgitコマンドを出力する.")

	cobra.EnableCommandSorting = false
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}
