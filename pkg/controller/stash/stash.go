package stash

import (
	"github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/infra"
	"github.com/spf13/cobra"
)

var StashCmd = &cobra.Command{
	Use:   "stash",
	Short: "スタッシュに関する操作.",
}

var git domain.Git = infra.NewGit()
var service = domain.NewService(git)

func init() {
	StashCmd.AddCommand(ListCmd)
	StashCmd.AddCommand(StoreCmd)
	StashCmd.AddCommand(RestoreCmd)
	StashCmd.AddCommand(DeleteCmd)
}
