package cmd

import (
	"os"

	"github.com/kazurego7/fit/fit/cmd/branch"
	"github.com/kazurego7/fit/fit/cmd/change"
	"github.com/kazurego7/fit/fit/cmd/commit"
	"github.com/kazurego7/fit/fit/cmd/config"
	"github.com/kazurego7/fit/fit/cmd/conflict"
	"github.com/kazurego7/fit/fit/cmd/repository"
	"github.com/kazurego7/fit/fit/cmd/stash"
	"github.com/kazurego7/fit/fit/cmd/tag"
	"github.com/kazurego7/fit/fit/global"
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

	RootCmd.PersistentFlags().BoolVar(&global.RootFlag.Dryrun, "dry-run", false, "データの変更なしに実行する.")
	RootCmd.PersistentFlags().BoolVar(&global.RootFlag.Debug, "debug", false, "fit内部で実行するgitコマンドを出力する.")

	cobra.EnableCommandSorting = false
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}
