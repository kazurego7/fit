package stash

import (
	"errors"

	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
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
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var gitSubCmd []string
		switch storeFlags.target {
		case "all":
			gitSubCmd = []string{"stash", "push", "--include-untracked"}
		case "index":
			gitSubCmd = []string{"stash", "push", "--staged"}
		case "worktree":
			gitSubCmd = []string{"stash", "push", "--include-untracked", "--keep-index"}
		default:
			return errors.New("invalid target of stash files")
		}
		// メッセージがあれば追加
		if len(args) != 0 {
			gitSubCmd = append(gitSubCmd, args[0])
		}
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
		return nil
	},
}

var storeFlags struct {
	target string
}

func init() {
	StoreCmd.Flags().StringVarP(&storeFlags.target, "target", "t", "all", `select target of stash files from "index", "worktree", "all"`)
}
