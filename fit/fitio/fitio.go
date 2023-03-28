package fitio

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CommandGit(dryrun bool, args ...string) int {
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

func QueryGit(args ...string) ([]byte, error) {
	cmd := exec.Command("git", args...)
	return cmd.Output()
}

func PrintGitCommand(dryrun bool, args ...string) {
	cmd := "git " + strings.Join(args, " ")
	if dryrun {
		fmt.Fprintln(os.Stderr, "dry-run: "+cmd)
	} else {
		fmt.Fprintln(os.Stderr, "command: "+cmd)
	}
}
