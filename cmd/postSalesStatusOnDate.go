package cmd

import (
	"einvoice/post"
	"time"

	"github.com/spf13/cobra"
)

// postSalesStatusOnDateCmd represents the postSalesStatusOnDate command
var postSalesStatusOnDateCmd = &cobra.Command{
	Use:   "postSalesStatusOnDate",
	Short: "Fetches sales invoice status changes on a specific date",
	Long: `This command fetches sales invoice status changes on a specific date and returns some informations into a file StatusSales.txt.
	File format: SalesInvoiceID;Status;Comment;Date;`,
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
		post.PostSalesStatusOnDate(date)
	},
}

func init() {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	postSalesStatusOnDateCmd.Flags().StringP("date", "d", yesterday, "Date (YYYY-MM-DD)")
	rootCmd.AddCommand(postSalesStatusOnDateCmd)
}
