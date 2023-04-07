package setting

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: "自動補完スクリプトを生成する.",
	Long: fmt.Sprintf(`自動補完スクリプトを設定する方法:

Bash:
  # Linux:
  $ %[1]s %[2]s completion --bash > /etc/bash_completion.d/%[1]s_%[2]s
  # macOS:
  $ %[1]s %[2]s completion --bash > $(brew --prefix)/etc/bash_completion.d/%[1]s_%[2]s

Zsh:
  $ %[1]s %[2]s completion --zsh > "${fpath[1]}/_%[1]s_%[2]s"
  # この設定を有効にするためには、新しいシェルを起動する必要があります。

fish:
  $ %[1]s %[2]s completion --fish > ~/.config/fish/completions/%[1]s_%[2]s.fish

PowerShell:
  PS> %[1]s %[2]s completion --powershell > %[1]s_%[2]s.ps1
  PS> echo '. "%[1]s_%[2]s.ps1"' >> $PROFILE
`, "fit", "setting"),
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch {
		case completionFlag.bash:
			cmd.Root().GenBashCompletion(os.Stdout)
		case completionFlag.zsh:
			cmd.Root().GenZshCompletion(os.Stdout)
		case completionFlag.fish:
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case completionFlag.powershell:
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			return errors.New("エラー：フラグが選択されていません.")
		}
		return nil
	},
}

var completionFlag struct {
	bash       bool
	zsh        bool
	fish       bool
	powershell bool
}

func init() {
	CompletionCmd.Flags().BoolVar(&completionFlag.bash, "bash", false, "output bash completion")
	CompletionCmd.Flags().BoolVar(&completionFlag.bash, "zsh", false, "output zsh completion")
	CompletionCmd.Flags().BoolVar(&completionFlag.bash, "fish", false, "output fish completion")
	CompletionCmd.Flags().BoolVar(&completionFlag.bash, "powershell", false, "output powershell completion")
	CompletionCmd.MarkFlagsMutuallyExclusive("bash", "zsh", "fish", "powershell")
}
