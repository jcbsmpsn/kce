package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/jcbsmpsn/kce/config"
	"github.com/jcbsmpsn/kce/delegate"
	"github.com/spf13/cobra"
)

func InitializeAliases() {
	for name, action := range config.Config.Alias {
		RootCmd.AddCommand(&cobra.Command{
			Use:   name,
			Short: "User defined alias",
			Long: `User defined alias that expands to

` + action,
			Run: selectAction(action),
		})
	}
}

func selectAction(action string) func(*cobra.Command, []string) {
	if strings.HasPrefix(action, "bash:") {
		return bashScriptAction(action)
	} else {
		return commandLineExpansionAction(action)
	}
}

func bashScriptAction(action string) func(cmd *cobra.Command, args []string) {
	action = strings.TrimPrefix(action, "bash:")
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to find 'bash'. %v\n", lookErr)
		return nil
	}
	return func(cmd *cobra.Command, args []string) {
		var execArgs []string
		execArgs = append(execArgs, "bash", "-c", action)
		// The following appends set of the list of command line arguments as
		// they will appear to the script as it runs. That is to say, these
		// values will be the values of $0, $1, $2 ... for the bash script that
		// executes.
		execArgs = append(execArgs, "kce-bash-alias")
		execArgs = append(execArgs, args...)
		env := os.Environ()
		execErr := syscall.Exec(binary, execArgs, env)
		if execErr != nil {
			fmt.Fprintf(os.Stderr, "Unable to run %q. %v\n", binary, execErr)
			return
		}
	}
}

func commandLineExpansionAction(action string) func(cmd *cobra.Command, args []string) {
	actionArgs := strings.Split(strings.Trim(action, " "), " ")
	return func(cmd *cobra.Command, args []string) {
		execArgs := []string{"kubectl"}
		execArgs = append(execArgs, actionArgs...)
		execArgs = append(execArgs, args...)
		fmt.Printf("Going to execute %v\n", execArgs)
		if err := delegate.Execute(execArgs); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to run 'kubectl'. %v\n", err)
		}
	}
}
