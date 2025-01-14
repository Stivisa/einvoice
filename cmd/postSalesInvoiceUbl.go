package cmd

import (
	"einvoice/post"

	"github.com/spf13/cobra"
)

// postSalesInvoiceUblCmd represents the postSalesInvoiceUbl command
var postSalesInvoiceUblCmd = &cobra.Command{
	Use:   "postSalesInvoiceUbl",
	Short: "Posts a sales invoice UBL",
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		post.PostSalesInvoiceUbl(file)
	},
}

func init() {
	postSalesInvoiceUblCmd.Flags().StringP("file", "f", "", "File path")
	getSalesInvoiceXmlCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(postSalesInvoiceUblCmd)
}
