package branch

import (
	"fmt"

	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var deleteConfirmMessage = `do you delete "` + args[0] + `" branch ? [yes/no] : `
		if !deleteFlag.yes {
			fmt.Print(deleteConfirmMessage)
		Prompt:
			for {
				var ans string
				fmt.Scanf("%s\n", &ans)
				switch ans {
				case "Yes", "Y", "yes", "y":
					break Prompt
				case "No", "N", "no", "n":
					fmt.Println("cancel delete branch")
					return
				default:
					fmt.Print(`put "yes" or "no". ` + deleteConfirmMessage)
					continue
				}
			}
		}
		gitSubCmd := []string{"branch", "--delete", "--force", args[0]}
		fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		fitio.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
}

var deleteFlag struct {
	yes bool
}

func init() {
	DeleteCmd.Flags().BoolVarP(&deleteFlag.yes, "yes", "y", false, "delete a branch without prompt confirmation")
}
