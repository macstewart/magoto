package cmd

import (
	"magoto/server"
	"os"

	"github.com/spf13/cobra"
)

const DEFAULT_PORT = 80

var rootCmd = &cobra.Command{
	Use: "magoto",
	Run: run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	var port = cmd.Flags().IntP("port", "p", DEFAULT_PORT, "The port to use for the server")
	server.Start(*port)
}
