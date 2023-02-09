// Package gomcoponentizer contain the implementation of the main command
package gomcoponentizer

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/leonardoce/gomcoponentizer/internal/logging"
)

// Cmd is the "gomcoponentizer" command
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "gomcoponentizer",
		Args: cobra.ExactArgs(1),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx := logging.IntoContext(
				cmd.Context(),
				viper.GetBool("debug"))
			cmd.SetContext(ctx)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fileName := args[0]
			node, err := parse(fileName)
			if err != nil {
				return err
			}

			fmt.Printf(`
			package hello

			func hello() {
			return `)
			spit(node, os.Stdout)
			fmt.Print("nil\n}")
			return nil
		},
	}

	cmd.PersistentFlags().Bool(
		"debug",
		true,
		"Enable debugging mode",
	)
	_ = viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))

	viper.SetEnvPrefix("gomcoponentizer")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	return cmd
}
