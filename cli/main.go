package main

import (
	"clean/cmd/cli/moduler"

	"github.com/spf13/cobra"
)

func main() {
	// first appends
	rootCmd.AddCommand(moduler.Moduler)

	// then run
	cobra.CheckErr(rootCmd.Execute())
}

var rootCmd = &cobra.Command{
	Use:   "lahij-cli",
	Short: "lahij cli tools !",
	Long:  "This CLI helps you generate controller files automatically using Go.",
}
