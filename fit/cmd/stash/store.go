package stash

import (
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
		switch {
		case storeFlags.worktree && storeFlags.index || !storeFlags.worktree && !storeFlags.index:
			gitSubCmd = []string{"stash", "push", "--include-untracked"}
		case storeFlags.index:
			gitSubCmd = []string{"stash", "push", "--staged"}
		case storeFlags.worktree:
			gitSubCmd = []string{"stash", "push", "--include-untracked", "--keep-index"}
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
	worktree bool
	index    bool
	all      bool
}

func init() {
	StoreCmd.Flags().BoolVarP(&storeFlags.worktree, "worktree", "w", false, "stash only worktree")
	StoreCmd.Flags().BoolVarP(&storeFlags.index, "index", "i", false, "stash only index")
	StoreCmd.MarkFlagsMutuallyExclusive("worktree", "index")
}
