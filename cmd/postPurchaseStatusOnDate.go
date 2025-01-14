package cmd

import (
	"einvoice/post"
	"time"

	"github.com/spf13/cobra"
)

// postPurchaseStatusOnDateCmd represents the postPurchaseStatusOnDate command
var postPurchaseStatusOnDateCmd = &cobra.Command{
	Use:   "postPurchaseStatusOnDate",
	Short: "Fetches purchase invoice status changes on a specific date",
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
		post.PostPurchaseStatusOnDate(date)
	},
}

func init() {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	postPurchaseStatusOnDateCmd.Flags().StringP("date", "d", yesterday, "Date (YYYY-MM-DD)")
	rootCmd.AddCommand(postPurchaseStatusOnDateCmd)
}
