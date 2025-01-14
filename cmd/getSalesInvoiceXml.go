package cmd

import (
	"einvoice/get"

	"github.com/spf13/cobra"
)

// getSalesInvoiceXmlCmd represents the getSalesInvoiceXml command
var getSalesInvoiceXmlCmd = &cobra.Command{
	Use:   "getSalesInvoiceXml",
	Short: "Fetches a sales invoice XML and writes it to a file SalesXML/invoiceId.xml",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		get.GetSalesInvoiceXml(id)
	},
}

func init() {
	getSalesInvoiceXmlCmd.Flags().StringP("id", "i", "", "Invoice ID")
	getSalesInvoiceXmlCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(getSalesInvoiceXmlCmd)
}
