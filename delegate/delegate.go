package delegate

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Execute will find a 'kubectl' instance on the path and execute it with args
// as the command line arguments. args[0] corresponds the name the executable
// is invoked as, so it corresponds to $0 in a bash script.
func Execute(args []string) error {
	binary, lookErr := exec.LookPath("kubectl")
	if lookErr != nil {
		return lookErr
	}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		return fmt.Errorf("Could not execute the 'kubectl' command: %v\n", execErr)
	}
	return nil
}
