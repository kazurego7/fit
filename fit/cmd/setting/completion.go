package setting

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: "コマンド補完スクリプトを出力する.",
	Long: fmt.Sprintf(`コマンド補完を設定する方法:

Bash: Linux
  $ %[1]s %[2]s %[3]s --bash > /etc/bash_completion.d/%[1]s_%[3]s

Bash: macOS
  $ %[1]s %[2]s %[3]s --bash > $(brew --prefix)/etc/bash_completion.d/%[1]s_%[3]s

Zsh:
  $ %[1]s %[2]s %[3]s --zsh > "${fpath[1]}/_%[1]s_%[3]s"
  ※この機能を有効にするためには、シェルを再起動する必要があります。

fish:
  $ %[1]s %[2]s %[3]s --fish > ~/.config/fish/completions/%[1]s_%[3]s.fish

PowerShell:
  PS> %[1]s %[2]s %[3]s --powershell > "$(Split-Path -Path $PROFILE)/%[1]s_%[3]s.ps1"
  PS> echo '. "$(Split-Path -Path $PROFILE)/%[1]s_%[3]s.ps1"' >> $PROFILE
  ※この機能を有効にするためには、シェルを再起動する必要があります。
`, "fit", "setting", "completion"),
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
			return errors.New("エラー：フラグが選択されていません")
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
	CompletionCmd.Flags().BoolVar(&completionFlag.zsh, "zsh", false, "output zsh completion")
	CompletionCmd.Flags().BoolVar(&completionFlag.fish, "fish", false, "output fish completion")
	CompletionCmd.Flags().BoolVar(&completionFlag.powershell, "powershell", false, "output powershell completion")
	CompletionCmd.MarkFlagsMutuallyExclusive("bash", "zsh", "fish", "powershell")
}
