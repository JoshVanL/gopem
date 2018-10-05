package cmd

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"

	"github.com/joshvanl/gopem/pkg/gen"
)

const (
	FlagKeySize = "key-size"
)

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

		size, err := cmd.PersistentFlags().GetInt(FlagKeySize)
		Must(err)

		var gens []*gen.KeyGen
		for _, a := range args {
			gens = append(gens, &gen.KeyGen{Path: a, Size: size, Type: "RSA"})
		}

		var result error
		for _, g := range gens {
			if err := g.Gen(); err != nil {
				result = multierror.Append(result, err)
			} else {
				fmt.Printf("generated key: %s\n", g.Path)
			}
		}
		Must(result)
	},
}

var genCACmd = &cobra.Command{
	Use:   "ca",
	Short: "generate fresh full x509 root CA",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	genKeyCmd.PersistentFlags().IntP(FlagKeySize, "s", 2048, "bit size of generated keys")
	genCmd.AddCommand(genKeyCmd)
	genCmd.AddCommand(genCACmd)
	RootCmd.AddCommand(genCmd)
}
