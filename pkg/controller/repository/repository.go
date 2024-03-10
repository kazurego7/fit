package repository

import (
	"github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/infra"
	"github.com/spf13/cobra"
)

var RepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "ローカルリポジトリ・リモートリポジトリに関する操作.",
}

var git domain.Git = infra.NewGit()
var service = domain.NewService(git)

func init() {
	RepositoryCmd.AddCommand(InitCmd)
	RepositoryCmd.AddCommand(CloneCmd)
	RepositoryCmd.AddCommand(RemoteCmd)
	RepositoryCmd.AddCommand(ConnectCmd)
	RepositoryCmd.AddCommand(DisconnectCmd)
}
