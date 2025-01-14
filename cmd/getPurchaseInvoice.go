package cmd

import (
	"einvoice/get"

	"github.com/spf13/cobra"
)

// getPurchaseInvoiceCmd represents the getPurchaseInvoice command
var getPurchaseInvoiceCmd = &cobra.Command{
	Use:   "getPurchaseInvoice",
	Short: "Get purchase invoice by ID and prints to console",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		get.GetPurchaseInvoice(id)
	},
}

func init() {
	// Defining the 'id' flag correctly and making it required
	getPurchaseInvoiceCmd.Flags().StringP("id", "i", "", "Invoice ID")
	getPurchaseInvoiceCmd.MarkFlagRequired("id") // Make 'id' flag required
	rootCmd.AddCommand(getPurchaseInvoiceCmd)
}
