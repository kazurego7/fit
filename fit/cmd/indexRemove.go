/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"kazurego7/fit/fit/gitexec"

	"github.com/spf13/cobra"
)

// indexRemoveCmd represents the indexRemove command
var indexRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"restore", "--staged"}
		allArgs := append(gitSubCmd, args...)
		gitexec.Git(allArgs...)
	},
}

func init() {
	indexCmd.AddCommand(indexRemoveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// indexRemoveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// indexRemoveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
