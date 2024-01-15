package cmd

import (
	"fmt"

	server2 "github.com/glaydsonp/go-hexagonal/adapters/web/server"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use: "http",
	Run: func(cmd *cobra.Command, args []string) {
		server := server2.MakeNewWebserver()
		server.Service = &productService
		fmt.Println("Webserver has been started")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
