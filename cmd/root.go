package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "einvoice",
	Short: "einvoice application for managing online invoices",
	Long:  `einvoice helps in interacting and processing e-invoices. Main purpose is for integrating with your ERP solution.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
