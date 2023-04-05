package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GitCommand(dryrun bool, args ...string) int {
	if dryrun {
		return 0
	}
	extArgs := append([]string{"-c", "core.quotepath=false"}, args...)
	cmd := exec.Command("git", extArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return cmd.ProcessState.ExitCode()
}

func GitQuery(args ...string) ([]byte, int, error) {
	extArgs := append([]string{"-c", "core.quotepath=false"}, args...)
	cmd := exec.Command("git", extArgs...)
	out, err := cmd.Output()
	exitCode := cmd.ProcessState.ExitCode()
	return out, exitCode, err
}

func PrintGitCommand(dryrun bool, args ...string) {
	cmd := "git " + strings.Join(args, " ")
	if dryrun {
		fmt.Fprintln(os.Stderr, "dry-run: "+cmd)
	} else {
		fmt.Fprintln(os.Stderr, "command: "+cmd)
	}
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
