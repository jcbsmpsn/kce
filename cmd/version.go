package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "kceversion",
	Short: "Print the version number of kce",
	Long:  `All software has versions. This is kce's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kce - kubectl expanded v0.1 -- HEAD")
	},
}
