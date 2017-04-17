package main

import (
	"fmt"
	"os"

	"github.com/jcbsmpsn/kce/cmd"
	"github.com/jcbsmpsn/kce/config"
	"github.com/jcbsmpsn/kce/delegate"
)

func main() {
	if err := config.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	cmd.InitializeAliases()

	if _, _, err := cmd.RootCmd.Find(os.Args[1:]); err == nil {
		if err := cmd.RootCmd.Execute(); err != nil {
			os.Exit(1)
		}
	} else {
		if err := delegate.Execute(os.Args); err != nil {
			os.Exit(1)
		}
	}
}
