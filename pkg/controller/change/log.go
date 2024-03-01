package change

import (
	service "github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var LogCmd = &cobra.Command{
	Use:   "log <pathspec>",
	Short: "単一のファイルの履歴を表示する.",
	Args:  service.ExistsFiles(1),
	Run: func(cmd *cobra.Command, args []string) {
		var gitSubCmd []string
		if logFlag.details {
			gitSubCmd = []string{
				"log",
				"--abbrev-commit",
				"--decorate=no",
				"--date=format: %Y-%m-%d %H:%I:%S",
				"--format=format:%C(bold 2)next-details%C(reset)\n%C(03)%h%C(reset)   %C(bold 0)%s%C(reset) %C(04)%ad%C(reset)  %C(green)%<(16,trunc)%an%C(reset)",
				"--follow",
				"-p",
				args[0],
			}

		} else {
			gitSubCmd = []string{
				"log",
				"--abbrev-commit",
				"--decorate=no",
				"--date=format: %Y-%m-%d %H:%I:%S",
				"--format=format:%C(03)%h%C(reset)   %C(bold 0)%s%C(reset) %>|(140)%C(reset)  %C(04)%ad%C(reset)  %C(green)%<(16,trunc)%an%C(reset)",
				"--follow",
				args[0],
			}
		}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}

var logFlag struct {
	details bool
}

func init() {
	LogCmd.Flags().BoolVarP(&logFlag.details, "details", "d", false, "ファイルの変更の詳細を表示する.")
}
