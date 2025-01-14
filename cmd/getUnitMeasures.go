package cmd

import (
	"einvoice/get"
	"fmt"

	"github.com/spf13/cobra"
)

// getAllCompaniesCmd represents the getAllCompanies command
var getUnitMeasuresCmd = &cobra.Command{
	Use:   "getUnitMeasures",
	Short: "Fetches all unit measures and writes them to JediniceMere.txt",
	Long: `This command fetches a list of unitmeasures via the API and writes their details to JediniceMere.txt. 
		Format of the file: Code;Symbol;NameEng;NameSrbLtn;NameSrbCyr;IsOnShortList;`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching all unit measures...")
		get.GetUnitMeasures()
	},
}

func init() {
	rootCmd.AddCommand(getUnitMeasuresCmd)
}
