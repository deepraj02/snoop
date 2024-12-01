/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"deepraj02/snoop/internal/server"
	"log"

	"github.com/spf13/cobra"
)

// / `port` : takes an optional port number to serve on.
var port string

var exclude []string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server on the network.",
	Long: `
Shares the content of your ` + "`" + `pwd` + "`" + ` on the network for other clients to download.

	`,
	Run: runServe,
}

func init() {

	/// serves on port 8080 by default
	/// can be changed using the -p / --port  flag
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to serve on")

	serveCmd.Flags().StringArrayVarP(&exclude, "exclude", "e", []string{}, "Type Files to Exclude from serving")

	rootCmd.AddCommand(serveCmd)
}

func runServe(cmd *cobra.Command, args []string) {
	srv := server.Spawn(port, exclude)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
