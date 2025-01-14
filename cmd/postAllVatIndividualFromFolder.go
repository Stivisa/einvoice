package cmd

import (
	"einvoice/post"
	"einvoice/utils"

	"github.com/spf13/cobra"
)

// postAllVatIndividualFromFolderCmd represents the postAllVatIndividualFromFolder command
var postAllVatIndividualFromFolderCmd = &cobra.Command{
	Use:   "postAllVatIndividualFromFolder",
	Short: "Posts all VAT individual from a folder",
	Long: `This command posts all VAT individual from a folder and removes the JSON files after processing.
	Returned informations will append in file OdgovorPdv.txt (or Odgovor.txt if option is 1).
	File OdgovorPdv.txt format: DocumentNumber;individualVatId;"vat";"vat";ErrorMessage`,
	Run: func(cmd *cobra.Command, args []string) {
		option, _ := cmd.Flags().GetInt("option")
		post.PostAllVatIndividualFromFolder(utils.VatFolderPath, option)
	},
}

func init() {
	postAllVatIndividualFromFolderCmd.Flags().StringP("folder", "f", "", "Folder path")
	postAllVatIndividualFromFolderCmd.Flags().IntP("option", "o", 1, "Option (0 or 1)")
	rootCmd.AddCommand(postAllVatIndividualFromFolderCmd)
}
