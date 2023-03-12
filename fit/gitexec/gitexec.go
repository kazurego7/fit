package gitexec

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Git(args ...string) {
	execMessage := fmt.Sprintf("execute command: git %v \n", strings.Join(args, " "))
	fmt.Fprint(os.Stdout, execMessage)
	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	defer cmd.Wait()
}
