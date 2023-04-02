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
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case deleteFlag.worktree:
			confirmBackup()
			deleteWorktree(args...)
		case deleteFlag.index:
			confirmBackup()
			deleteIndex(args...)
		case deleteFlag.all:
			confirmBackup()
			deleteAll(args...)
		}
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

func deleteWorktree(args ...string) {
	unergedList := searchWorktreeList("U", args[0])
	for i := range unergedList {
		unergedList[i] = ":!" + unergedList[i]
	}
	restoreList := searchWorktreeList("", append(unergedList, args[0])...)
	if len(restoreList) != 0 {
		exitCode := restoreWorktree(restoreList...)
		if exitCode != 0 {
			return
		}
	}
	addedList := searchWorktreeList("A", args[0])
	if len(addedList) != 0 {
		exitCode := removeIndex(addedList...)
		if exitCode != 0 {
			return
		}
	}
	clean(args[0])
}

func deleteIndex(args ...string) {
	indexList := searchIndexList("", args[0])
	worktreeList := searchWorktreeList("", args[0])
	indexOnlyList := util.Difference(indexList, worktreeList)
	restoreList := searchIndexList("a", indexOnlyList...)
	cleanList := searchIndexList("A", indexOnlyList...)
	if len(indexList) != 0 {
		exitCode := restoreIndex(indexList...)
		if exitCode != 0 {
			return
		}
	}
	if len(restoreList) != 0 {
		exitCode := restoreWorktree(restoreList...)
		if exitCode != 0 {
			return
		}
	}
	if len(cleanList) != 0 {
		exitCode := clean(cleanList...)
		if exitCode != 0 {
			return
		}
	}
}

func deleteAll(args ...string) {
	indexList := searchIndexList("", args[0])
	if len(indexList) != 0 {
		exitCode := restoreIndex(indexList...)
		if exitCode != 0 {
			return
		}
	}
	addedList := searchWorktreeList("A", args[0])
	if len(addedList) != 0 {
		exitCode := removeIndex(addedList...)
		if exitCode != 0 {
			return
		}
	}
	restoreList := searchWorktreeList("a", args[0])
	if len(restoreList) != 0 {
		exitCode := restoreWorktree(restoreList...)
		if exitCode != 0 {
			return
		}
	}
	exitCode := clean(args[0])
	if exitCode != 0 {
		return
	}
}
