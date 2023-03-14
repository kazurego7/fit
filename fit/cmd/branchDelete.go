/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"kazurego7/fit/fit/gitexec"

	"github.com/spf13/cobra"
)

// branchDeleteCmd represents the branchDelete command
var branchDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"branch", "--delete", args[0]}
		gitexec.Git(gitSubCmd...)
	},
}

func init() {
	branchCmd.AddCommand(branchDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
