package repository

import (
	"github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/infra/gitImpl"
	"github.com/spf13/cobra"
)

var RepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "ローカルリポジトリ・リモートリポジトリに関する操作.",
}

var git = gitImpl.Git{}
var service = domain.NewService(git)

func init() {
	RepositoryCmd.AddCommand(InitCmd)
	RepositoryCmd.AddCommand(CloneCmd)
	RepositoryCmd.AddCommand(RemoteCmd)
	RepositoryCmd.AddCommand(ConnectCmd)
	RepositoryCmd.AddCommand(DisconnectCmd)
}
