package change

import (
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch {
		case deleteFlag.worktree:
			confirmBackup()
			restoreWorktree(args[0])
			clean(args[0])
		case deleteFlag.index:
			confirmBackup()
			indexList := searchIndexList("", args[0])
			worktreeList := searchWorktreeList("", args[0])
			indexOnlyList := util.Difference(indexList, worktreeList)
			restoreList := searchIndexList("a", indexOnlyList...)
			cleanList := searchIndexList("A", indexOnlyList...)
			if len(indexList) != 0 {
				restoreIndex(indexList...)
			}
			if len(restoreList) != 0 {
				restoreWorktree(restoreList...)
			}
			if len(cleanList) != 0 {
				clean(cleanList...)
			}
		case deleteFlag.all:
			confirmBackup()
			indexList := searchIndexList("", args[0])
			if len(indexList) != 0 {
				restoreIndex(indexList...)
			}
			restoreList := searchWorktreeList("a", args[0])
			if len(restoreList) != 0 {
				restoreWorktree(restoreList...)
			}
			clean(args[0])
		}
		return nil
	},
}

var deleteFlag struct {
	worktree bool
	index    bool
	all      bool
}

func init() {
	DeleteCmd.Flags().BoolVarP(&deleteFlag.worktree, "worktree", "w", false, "delete changes worktree to index")
	DeleteCmd.Flags().BoolVarP(&deleteFlag.index, "index", "i", false, "delete changes index to HEAD")
	DeleteCmd.Flags().BoolVarP(&deleteFlag.all, "all", "a", true, "delete changes index and worktree")
	DeleteCmd.MarkFlagsMutuallyExclusive("worktree", "index", "all")
}
