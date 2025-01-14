package cmd

import (
	"einvoice/get"
	"fmt"

	"github.com/spf13/cobra"
)

// getAllCompaniesCmd represents the getAllCompanies command
var getAllCompaniesCmd = &cobra.Command{
	Use:   "getAllCompanies",
	Short: "Fetches all companies and writes them to Firme.txt",
	Long: `This command fetches a list of companies via the API and writes their details to Firme.txt. 
		Format of the file: BugetCompanyNumber;RegistrationCode;VatRegistrationCode;Name;`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching all companies...")
		get.GetAllCompanies()
	},
}

func init() {
	rootCmd.AddCommand(getAllCompaniesCmd)
}
