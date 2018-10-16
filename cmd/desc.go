package cmd

import (
	"errors"
	//"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"

	"github.com/joshvanl/gopem/pkg/desc"
)

var descCmd = &cobra.Command{
	Use:     "describe [files..]",
	Short:   "describe any pem type file contents",
	Aliases: []string{"desc", "des", "d"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			Must(errors.New("provide at least one pem file to describe"))
		}

		var result *multierror.Error
		for _, a := range args {
			d, err := desc.New(a)
			if err != nil {
				result = multierror.Append(result, err)
				continue
			}

			if err := d.Describe(); err != nil {
				result = multierror.Append(result, err)
			}
		}

		Must(result.ErrorOrNil())
	},
}

func init() {
	RootCmd.AddCommand(descCmd)
}
