package setting

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion script",
	Long: fmt.Sprintf(`To load completions:

Bash:

  $ source <(%[1]s %[2]s completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ %[1]s %[2]s completion --bash > /etc/bash_completion.d/%[1]s_%[2]s
  # macOS:
  $ %[1]s %[2]s completion --bash > $(brew --prefix)/etc/bash_completion.d/%[1]s_%[2]s

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ %[1]s %[2]s completion --zsh > "${fpath[1]}/_%[1]s_%[2]s"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ %[1]s %[2]s completion --fish | source

  # To load completions for each session, execute once:
  $ %[1]s %[2]s completion --fish > ~/.config/fish/completions/%[1]s_%[2]s.fish

PowerShell:

  PS> %[1]s %[2]s completion --powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> %[1]s %[2]s completion --powershell > %[1]s_%[2]s.ps1
  # and source this file from your PowerShell profile.
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
			return errors.New("flag not selected error")
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
