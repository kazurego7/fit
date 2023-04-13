package cmd

import (
	"os"

	"github.com/kazurego7/fit/fit/cmd/branch"
	"github.com/kazurego7/fit/fit/cmd/change"
	"github.com/kazurego7/fit/fit/cmd/conflict"
	"github.com/kazurego7/fit/fit/cmd/repository"
	"github.com/kazurego7/fit/fit/cmd/revision"
	"github.com/kazurego7/fit/fit/cmd/setting"
	"github.com/kazurego7/fit/fit/cmd/stash"
	"github.com/kazurego7/fit/fit/cmd/tag"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "fit",
	Short: "ユーザーフレンドリーな git CLI.",
	ValidArgsFunction: cobra.FixedCompletions(
		[]string{
			change.ChangeCmd.Name(),
			revision.RevisionCmd.Name(),
			conflict.ConflictCmd.Name(),
			branch.BranchCmd.Name(),
			stash.StashCmd.Name(),
			tag.TagCmd.Name(),
			repository.RepositoryCmd.Name(),
			setting.SettingCmd.Name(),
			"help",
			"completion",
		}, cobra.ShellCompDirectiveKeepOrder),
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(change.ChangeCmd)
	RootCmd.AddCommand(revision.RevisionCmd)
	RootCmd.AddCommand(conflict.ConflictCmd)
	RootCmd.AddCommand(branch.BranchCmd)
	RootCmd.AddCommand(stash.StashCmd)
	RootCmd.AddCommand(tag.TagCmd)
	RootCmd.AddCommand(repository.RepositoryCmd)
	RootCmd.AddCommand(setting.SettingCmd)

	RootCmd.PersistentFlags().BoolVar(&global.RootFlag.Dryrun, "dry-run", false, "データの変更なしに実行する.")
	RootCmd.PersistentFlags().BoolVar(&global.RootFlag.Debug, "debug", false, "fit内部で実行するgitコマンドを出力する.")

	cobra.EnableCommandSorting = false
}
