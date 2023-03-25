package fitio

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecuteGit(args ...string) {
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

func PrintGitCommand(args ...string) {
	execMessage := fmt.Sprintf("execute command: git %v \n", strings.Join(args, " "))
	fmt.Fprint(os.Stderr, execMessage)
}
