package repository

import (
	"github.com/spf13/cobra"
)

var RepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "ローカルリポジトリ・リモートリポジトリに関する操作.",
}

func init() {
	RepositoryCmd.AddCommand(AsyncCmd)
	RepositoryCmd.AddCommand(CloneCmd)
	RepositoryCmd.AddCommand(InitCmd)
	RepositoryCmd.AddCommand(RemoteCmd)
	RepositoryCmd.AddCommand(UnsyncCmd)
}
