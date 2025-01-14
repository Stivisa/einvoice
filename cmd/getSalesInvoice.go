package cmd

import (
	"einvoice/get"

	"github.com/spf13/cobra"
)

// getPurchaseInvoiceCmd represents the getPurchaseInvoice command
var getSalesInvoiceCmd = &cobra.Command{
	Use:   "getSalesInvoice",
	Short: "Get sales invoice by ID and prints to console",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		get.GetSalesInvoice(id)
	},
}

func init() {
	// Defining the 'id' flag correctly and making it required
	getSalesInvoiceCmd.Flags().StringP("id", "i", "", "Invoice ID")
	getSalesInvoiceCmd.MarkFlagRequired("id") // Make 'id' flag required
	rootCmd.AddCommand(getSalesInvoiceCmd)
}
