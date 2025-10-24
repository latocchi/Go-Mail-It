/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/latocchi/Go-Mail-It/internal/keyring"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var provider string

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup email provider to use",
	Run: func(cmd *cobra.Command, args []string) {
		var email, password string

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Your email: ")
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)

		fmt.Print("Your password/app password: ")
		bytePassword, _ := term.ReadPassword(int(os.Stdin.Fd()))
		password = strings.TrimSpace(string(bytePassword))

		if err := keyring.SaveCredentials(email, password); err != nil {
			panic(err)
		}

		fmt.Println("\nCredentials saved successfully!")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setupCmd.Flags().StringVarP(&provider, "provider", "p", "", "Email provider to use (e.g., 'google')")
	if err := setupCmd.MarkFlagRequired("provider"); err != nil {
		panic(err)
	}
}
