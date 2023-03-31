package change

import (
	"fmt"

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
		switch clearFlag.target {
		case "worktree":
			confirmBackup()
			restoreWorktree(args[0])
			clean(args[0])
		case "index":
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
		case "all":
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
		default:
			return fmt.Errorf(`"%v" is invalid in "--target" flag. use "worktree", "index" or "all"`, clearFlag.target)
		}
		return nil
	},
}

var clearFlag struct {
	target string
}

func init() {
	DeleteCmd.Flags().StringVarP(&clearFlag.target, "target", "t", "all", `clear target from "worktree", "index" or "all"`)
}
