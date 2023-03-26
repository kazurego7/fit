package stash

import (
	"errors"

	"github.com/kazurego7/fit/fit/fitio"
	"github.com/spf13/cobra"
)

var StoreCmd = &cobra.Command{
	Use:   "store",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var gitSubCmd []string
		switch storeFlags.kind {
		case "all":
			gitSubCmd = []string{"stash", "push", "--include-untracked"}
		case "index":
			gitSubCmd = []string{"stash", "push", "--staged"}
		case "worktree":
			gitSubCmd = []string{"stash", "push", "--include-untracked", "--keep-index"}
		default:
			return errors.New("invalid kind of stash files")
		}
		fitio.PrintGitCommand(gitSubCmd...)
		fitio.ExecuteGit(gitSubCmd...)
		return nil
	},
}

var storeFlags struct {
	kind string
}

func init() {
	StoreCmd.Flags().StringVarP(&storeFlags.kind, "kind", "k", "all", `select kind of stash files from "index", "worktree", "all"`)
}
