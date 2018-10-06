package cmd

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"

	"github.com/joshvanl/gopem/pkg/gen"
)

const (
	FlagSize   = "size"
	FlagType   = "type"
	FlagPrefix = "prefix"
)

var flags gen.Flags

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate fresh x509 pem keys and certs",
}

var genKeyCmd = &cobra.Command{
	Use:   "key",
	Short: "generate fresh private keys",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			Must(errors.New("must provide at least one private key file"))
		}

		var result *multierror.Error
		for _, a := range args {
			flags.Path = a
			g := gen.NewKeyGen(&flags)
			if err := g.Gen(); err != nil {
				result = multierror.Append(result, err)
			} else {
				fmt.Printf("generated key: %s\n", a)
			}
		}

		Must(result.ErrorOrNil())
	},
}

var genCACmd = &cobra.Command{
	Use:   "ca",
	Short: "generate fresh full x509 root CA",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	genCmd.PersistentFlags().IntVarP(&flags.Size, FlagSize, "s", 2048, "bit size of generated keys")
	genCmd.PersistentFlags().StringVarP(&flags.Type, FlagType, "t", "RSA", "key algorithm")
	genCmd.PersistentFlags().StringVarP(&flags.Prefix, FlagPrefix, "p", "", "prefix to all file names")
	genCmd.AddCommand(genKeyCmd)
	genCmd.AddCommand(genCACmd)
	RootCmd.AddCommand(genCmd)
}
