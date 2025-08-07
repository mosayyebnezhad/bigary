package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var rootCmd = &cobra.Command{
	Use:   "tinker",
	Short: "Tinker CLI",
}

var cmdHash = &cobra.Command{
	Use:   "hash [string]",
	Short: "Hash a string",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		hashed, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("failed to hash string: %v", err)
		}
		fmt.Println(string(hashed))
	},
}

func main() {
	rootCmd.AddCommand(cmdHash)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}