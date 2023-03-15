/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

	"github.com/kazurego7/fit/fit/gitexec"
	"github.com/spf13/cobra"
)

// branchCreateCmd represents the branchCreate command
var branchCreateCmd = &cobra.Command{
	Use:   "create",
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
		gitexec.Git("branch", args[0])
	},
}

func init() {
	branchCmd.AddCommand(branchCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
