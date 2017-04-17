package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kce",
	Short: "kce (kubectl expanded) provides convenient aliases, defaults and option expansion for kubectl",
	Long:  `kce (kubectl expanded) provides convenient aliases, defaults and option expansion for kubectl`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`kce is a command line program that replaces kubectl for controlling a
Kubernetes cluster. By default it behaves like kubectl but offers options for
customizing and extending the behavior of kubectl.
`)

		maxLength := 0
		for _, cmd := range cmd.Commands() {
			if len(cmd.Name()) > maxLength {
				maxLength = len(cmd.Name())
			}
		}

		fmtString := fmt.Sprintf("  %%-%ds %%s\n", maxLength+5)
		commands := cmd.Commands()
		sort.Sort(ByName(commands))
		for _, cmd := range commands {
			fmt.Printf(fmtString, cmd.Name(), cmd.Short)
		}

		fmt.Println(`
Use "kce <command> --help" for more information about a given command.
`)
	},
}

// ByName implements sort.Interface for []cobra.Command based on
// the Name field.
type ByName []*cobra.Command

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name() < a[j].Name() }
