package fitio

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
	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return cmd.ProcessState.ExitCode()
}

func GitQuery(args ...string) ([]byte, int, error) {
	cmd := exec.Command("git", args...)
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
