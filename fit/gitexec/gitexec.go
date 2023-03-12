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
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Fprint(os.Stdout, string(out))
}
