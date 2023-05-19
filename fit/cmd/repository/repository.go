package repository

import (
	"github.com/spf13/cobra"
)

var RepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "ローカルリポジトリ・リモートリポジトリに関する操作.",
}

func init() {
	RepositoryCmd.AddCommand(FetchCmd)
	RepositoryCmd.AddCommand(InitCmd)
	RepositoryCmd.AddCommand(CloneCmd)
	RepositoryCmd.AddCommand(RemoteCmd)
	RepositoryCmd.AddCommand(ConnectCmd)
	RepositoryCmd.AddCommand(DisconnectCmd)
}
