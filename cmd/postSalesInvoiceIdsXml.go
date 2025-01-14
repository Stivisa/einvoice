package cmd

import (
	"einvoice/post"
	"time"

	"github.com/spf13/cobra"
)

// postSalesInvoiceIdsXmlCmd represents the postSalesInvoiceIdsXml command
var postSalesInvoiceIdsXmlCmd = &cobra.Command{
	Use:   "postSalesInvoiceIdsXml",
	Short: "Get some informations about sales invoice.",
	Long: `This command fetches sales invoices by status,
		 date from and date to and returns some informations into a file OdgovorProdaja.txt and OdgovorProdajaStavke.txt
		File format OdgovorProdaja.txt: SalesInvoiceID;PurchaseInvoiceID;BrojFakture;DatumIzdavanja;DatumDospeca;
		TipFakture;ValutaDokumenta;FakturaPeriodOpisKod;PosiljaocPIB;PosiljaocMaticniBroj;PosiljaocNazivFirme;
		DatumIsporuke;ModelPozivNaBroj;TaxTotalIznos;TaxOsnovica1;TaxIznos1;TaxOznaka1;TaxPosto1;
		TaxOsnovica2;TaxIznos2;TaxOznaka2;TaxPosto2;TaxOsnovica3;TaxIznos3;TaxOznaka3;TaxPosto3;
		VrednostStavki;VrednostOsnoviceUkupno;IznosRacuna;DatiRabat;AvansnoPlaceno;IznosZaPlacanje;
		File format OdgovorProdajaStavke.txt: SalesInvoiceID;PurchaseInvoiceID;
		Oznaka;Kolicina;UkupanIznos;TeretIndikator;TeretProcenat;TeretIznos;Naziv;OznakaProdavca;TaxOznaka;TaxPosto;Iznos;`,
	Run: func(cmd *cobra.Command, args []string) {
		statusInvoice, _ := cmd.Flags().GetString("status")
		dateFrom, _ := cmd.Flags().GetString("from")
		dateTo, _ := cmd.Flags().GetString("to")
		post.PostSalesInvoiceIdsXml(statusInvoice, dateFrom, dateTo)
	},
}

func init() {
	postSalesInvoiceIdsXmlCmd.Flags().StringP("status", "s", "Approved", "Status invoice")
	dateFrom := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	postSalesInvoiceIdsXmlCmd.Flags().StringP("from", "f", dateFrom, "Date from")
	dateTo := time.Now().Format("2006-01-02")
	postSalesInvoiceIdsXmlCmd.Flags().StringP("to", "t", dateTo, "Date to")
	rootCmd.AddCommand(postSalesInvoiceIdsXmlCmd)
}
