package cmd

import (
	"os"

	"github.com/kazurego7/fit/fit/cmd/branch"
	"github.com/kazurego7/fit/fit/cmd/conflict"
	"github.com/kazurego7/fit/fit/cmd/history"
	"github.com/kazurego7/fit/fit/cmd/index"
	"github.com/kazurego7/fit/fit/cmd/repository"
	"github.com/kazurego7/fit/fit/cmd/setting"
	"github.com/kazurego7/fit/fit/cmd/stash"
	"github.com/kazurego7/fit/fit/cmd/tag"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "fit",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	RootCmd.AddCommand(index.IndexCmd)
	RootCmd.AddCommand(conflict.ConflictCmd)
	RootCmd.AddCommand(repository.RepositoryCmd)
	RootCmd.AddCommand(stash.StashCmd)
	RootCmd.AddCommand(tag.TagCmd)

	RootCmd.PersistentFlags().BoolVarP(&global.Flags.Dryrun, "dry-run", "n", false, "don't actually run git command")
}
