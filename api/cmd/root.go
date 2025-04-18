package cmd

import (
	"app/cmd/loadStock"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(loadStock.Cmd)
}

func Execute() bool {
	if len(os.Args) == 1 {
		return false
	}

	err := rootCmd.Execute()

	if err != nil {
		log.Fatal("Error al ejecutar el comando:", err)
	}

	return true
}
