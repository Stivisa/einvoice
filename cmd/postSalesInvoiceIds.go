package cmd

import (
	"einvoice/post"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// postSalesInvoiceIdsCmd represents the postSalesInvoiceIds command
var postSalesInvoiceIdsCmd = &cobra.Command{
	Use:   "postSalesInvoiceIds",
	Short: "Fetches sales by status, date from and date to and returns invoice IDs and writes them to a file StatusSales.txt",
	Run: func(cmd *cobra.Command, args []string) {
		statusInvoice, _ := cmd.Flags().GetString("status")
		dateFrom, _ := cmd.Flags().GetString("from")
		dateTo, _ := cmd.Flags().GetString("to")
		post.PostSalesInvoiceIds(statusInvoice, dateFrom, dateTo)
	},
}

func init() {
	postSalesInvoiceIdsCmd.Flags().StringP("status", "s", "Approved", "Status invoice")
	currentYear := time.Now().Year()
	postSalesInvoiceIdsCmd.Flags().StringP("from", "f", fmt.Sprintf("%d-01-01", currentYear), "Date from")
	postSalesInvoiceIdsCmd.Flags().StringP("to", "t", "", "Date to")
	rootCmd.AddCommand(postSalesInvoiceIdsCmd)
}
