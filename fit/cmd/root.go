package cmd

import (
	"os"

	"github.com/kazurego7/fit/fit/cmd/branch"
	"github.com/kazurego7/fit/fit/cmd/change"
	"github.com/kazurego7/fit/fit/cmd/conflict"
	"github.com/kazurego7/fit/fit/cmd/history"
	"github.com/kazurego7/fit/fit/cmd/repository"
	"github.com/kazurego7/fit/fit/cmd/setting"
	"github.com/kazurego7/fit/fit/cmd/stash"
	"github.com/kazurego7/fit/fit/cmd/tag"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "fit",
	Short: "単なるgitコマンドのファサード.",
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(branch.BranchCmd)
	RootCmd.AddCommand(setting.SettingCmd)
	RootCmd.AddCommand(history.HistoryCmd)
	RootCmd.AddCommand(change.ChangeCmd)
	RootCmd.AddCommand(conflict.ConflictCmd)
	RootCmd.AddCommand(repository.RepositoryCmd)
	RootCmd.AddCommand(stash.StashCmd)
	RootCmd.AddCommand(tag.TagCmd)

	RootCmd.PersistentFlags().BoolVarP(&global.Flags.Dryrun, "dry-run", "n", false, "実際にgitコマンドを実行しない.")
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}
