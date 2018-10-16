package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gopem",
	Short: "gopem is a utility tool for generating and signing PKI certs among other things. (openssl sucks)",
}

func Execute(args []string) {
	RootCmd.SetArgs(args)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}

func Must(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
