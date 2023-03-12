/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"kazurego7/fit/fit/gitexec"

	"github.com/spf13/cobra"
)

// historyCommitCmd represents the historyCommit command
var historyCommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("args error")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"commit", "-m", args[0]}
		gitexec.Git(gitSubCmd...)
	},
}

func init() {
	historyCmd.AddCommand(historyCommitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// historyCommitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// historyCommitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
