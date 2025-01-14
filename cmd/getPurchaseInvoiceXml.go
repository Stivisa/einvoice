package cmd

import (
	"einvoice/get"

	"github.com/spf13/cobra"
)

// getPurchaseInvoiceXmlCmd represents the getPurchaseInvoiceXml command
var getPurchaseInvoiceXmlCmd = &cobra.Command{
	Use:   "getPurchaseInvoiceXml",
	Short: "Fetches a purchase invoice XML and writes it to a file PurchaseXML/invoiceId.xml",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		get.GetPurchaseInvoiceXml(id)
	},
}

func init() {
	getPurchaseInvoiceXmlCmd.Flags().StringP("id", "i", "", "Invoice ID")
	getPurchaseInvoiceXmlCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(getPurchaseInvoiceXmlCmd)
}
