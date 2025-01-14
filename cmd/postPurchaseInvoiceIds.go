package cmd

import (
	"einvoice/post"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// postPurchaseInvoiceIdsCmd represents the postPurchaseInvoiceIds command
var postPurchaseInvoiceIdsCmd = &cobra.Command{
	Use:   "postPurchaseInvoiceIds",
	Short: "Fetches purchase by status, date from and date to and returns invoice IDs and writes them to a file StatusPurchase.txt",
	Run: func(cmd *cobra.Command, args []string) {
		statusInvoice, _ := cmd.Flags().GetString("status")
		dateFrom, _ := cmd.Flags().GetString("from")
		dateTo, _ := cmd.Flags().GetString("to")
		post.PostPurchaseInvoiceIds(statusInvoice, dateFrom, dateTo)
	},
}

func init() {
	postPurchaseInvoiceIdsCmd.Flags().StringP("status", "s", "Approved", "Status invoice")
	currentYear := time.Now().Year()
	postPurchaseInvoiceIdsCmd.Flags().StringP("from", "f", fmt.Sprintf("%d-01-01", currentYear), "Date from")
	postPurchaseInvoiceIdsCmd.Flags().StringP("to", "t", "", "Date to")
	rootCmd.AddCommand(postPurchaseInvoiceIdsCmd)
}
