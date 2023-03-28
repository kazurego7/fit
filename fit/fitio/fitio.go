package fitio

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecuteGit(dryrun bool, args ...string) {
	if dryrun {
		return
	}
	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	defer cmd.Wait()
}

func ExecuteGitOutput(args ...string) ([]byte, error) {
	cmd := exec.Command("git", args...)
	return cmd.Output()
}

func PrintGitCommand(dryrun bool, args ...string) {
	cmd := "git " + strings.Join(args, " ")
	if dryrun {
		fmt.Fprint(os.Stderr, "dry-run: "+cmd)
	} else {
		fmt.Fprint(os.Stderr, "command: "+cmd)
	}
}
