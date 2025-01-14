package cmd

import (
	"einvoice/post"

	"github.com/spf13/cobra"
)

// postSalesInvoiceIdsXmlManuallyCmd represents the postSalesInvoiceIdsXmlManually command
var postSalesInvoiceIdsXmlManuallyCmd = &cobra.Command{
	Use:   "postSalesInvoiceIdsXmlManually",
	Short: "Finds invoice names based on SalesId",
	Long: `This command finds invoice names based on sales invoice IDs from RucnoUnos.txt and returns into a file OdgovorRucno.txt.
	File format OdgovorRucno.txt:  SalesId;PurchaseId;DocumentId(document name);
	File format RucnoUnos.txt: id\n id\n id...`,
	Run: func(cmd *cobra.Command, args []string) {
		post.PostSalesInvoiceIdsXmlManually()
	},
}

func init() {
	rootCmd.AddCommand(postSalesInvoiceIdsXmlManuallyCmd)
}
