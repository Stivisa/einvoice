package cmd

import (
	"einvoice/post"
	"einvoice/utils"

	"github.com/spf13/cobra"
)

// postAllSalesInvoiceFromFolderCmd represents the postAllSalesInvoiceFromFolder command
var postAllSalesInvoiceFromFolderCmd = &cobra.Command{
	Use:   "postAllSalesInvoiceFromFolder",
	Short: "Posts all sales invoices from a folder.",
	Long: `This command posts all sales invoices from a folder FolderPath from .ini and removes the XML files after processing. 
	Returned informations will append in file Odgovor.txt.
	File format: DocumentNumber;InvoiceID;PurchaseInvoiceID;SalesInvoiceID;ErrorMessage`,
	Run: func(cmd *cobra.Command, args []string) {
		folder, _ := cmd.Flags().GetString("folder")
		if folder == "" {
			folder = utils.FolderPath
		}
		post.PostAllSalesInvoiceFromFolder(folder)
	},
}

func init() {
	postAllSalesInvoiceFromFolderCmd.Flags().StringP("folder", "f", "", "Folder path")
	rootCmd.AddCommand(postAllSalesInvoiceFromFolderCmd)
}
