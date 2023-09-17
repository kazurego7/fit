package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"fit/pkg/usecase"
)

func GitCommand(globalFlag usecase.GlobalFlag, args []string) int {
	extArgs := append([]string{"-c", "core.quotepath=false"}, args...)

	if usecase.RootFlag.Debug {
		if globalFlag.Dryrun {
			fmt.Fprintln(os.Stderr, "dry-run: git "+strings.Join(extArgs, " "))
		} else {
			fmt.Fprintln(os.Stderr, "command: git "+strings.Join(extArgs, " "))
		}
	}

	if globalFlag.Dryrun {
		return 0
	}

	cmd := exec.Command("git", extArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return cmd.ProcessState.ExitCode()
}

func GitQuery(globalFlag usecase.GlobalFlag, args []string) ([]byte, int, error) {
	extArgs := append([]string{"-c", "core.quotepath=false"}, args...)

	if usecase.RootFlag.Debug {
		fmt.Fprintln(os.Stderr, "query: git "+strings.Join(extArgs, " "))
	}

	cmd := exec.Command("git", extArgs...)
	out, err := cmd.Output()
	exitCode := cmd.ProcessState.ExitCode()
	return out, exitCode, err
}

func InputYesOrNo(allwaysYes bool) (bool, error) {
	if allwaysYes {
		return true, nil
	}
	for {
		var ans string
		_, err := fmt.Scanf("%s\n", &ans)
		if err != nil {
			return false, err
		}
		switch ans {
		case "Yes", "Y", "yes", "y":
			return true, nil
		case "No", "N", "no", "n":
			return false, nil
		default:
			fmt.Print(`put "yes" or "no" : `)
			continue
		}
	}
}
