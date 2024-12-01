/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Version represents the current version of snoop
const Version = "v2.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "snoop",
	Short: "Over the Network File Transfer tool.",
	Long: `
Navigate to the directory you want to share 
and run snoop serve to start the server.

That's it
	`,
	Version: Version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
