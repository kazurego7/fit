/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kazurego7/fit/fit/gitexec"
	"github.com/spf13/cobra"
)

// remoteListCmd represents the remoteList command
var remoteListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"remote", "--verbose"}
		gitexec.Git(gitSubCmd...)
	},
}

func init() {
	remoteCmd.AddCommand(remoteListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
