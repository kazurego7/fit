/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"kazurego7/fit/fit/gitexec"

	"github.com/spf13/cobra"
)

// remoteAddCmd represents the remoteAdd command
var remoteAddCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"remote", "add", "origin", args[0]}
		gitexec.Git(gitSubCmd...)
	},
}

func init() {
	remoteCmd.AddCommand(remoteAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
