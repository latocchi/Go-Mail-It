/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/latocchi/Go-Mail-It/internal/providers"
	"github.com/spf13/cobra"
)

var provider string

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup email provider to use",
	Run: func(cmd *cobra.Command, args []string) {
		provider = args[0]
		switch provider {
		case "google", "gmail":
			_, err := providers.SetupGoogle()
			if err != nil {
				fmt.Println("Error setting up Google provider:", err)
			}
		default:
			fmt.Println("Unsupported provider:", provider)
			// TODO: Change default so that the program exits with error code
		}
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

}
