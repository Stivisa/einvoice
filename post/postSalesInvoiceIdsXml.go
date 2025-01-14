package post

import (
	"einvoice/utils"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func PostSalesInvoiceIdsXml(statusInvoice string, dateFrom string, dateTo string) {

	var body []byte
	var code int
	var status string
	var approvedIds []int

	url, err := url.JoinPath(utils.UrlPath, "sales-invoice/ids")
	if err != nil {
		log.Println(err)
	}

	body, code, status = utils.PostQ(url, "status", statusInvoice, "dateFrom", dateFrom, "dateTo", dateTo)

	if code == 200 {
		var SalesInvoicesDto1 utils.SalesInvoicesDto
		if err := json.Unmarshal(body, &SalesInvoicesDto1); err != nil {
			log.Fatalln("Can not unmarshal JSON" + err.Error())
		}
		approvedIds = SalesInvoicesDto1.SalesInvoiceIds
		log.Printf("%+v\n", approvedIds)
	} else {
		log.Println(status)
		log.Fatalln(string(body))
	}

	f, err := os.OpenFile(filepath.Join(utils.FolderPath, "OdgovorProdaja.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	for _, item := range approvedIds {
		//log.Println(item)
		xmlBody := SalesInvoiceXml(strconv.Itoa(item))

		var icDto utils.InvoiceCustomDto
		if err := xml.Unmarshal(xmlBody, &icDto); err != nil {
			log.Fatalln("Can not unmarshal JSON" + err.Error())
		}

		var taxSubtotalsFlattened string
		for i := 0; i < 3; i++ {
			if i < len(icDto.DocumentBody.Invoice.TaxTotal.TaxSubtotal) {
				subtotal := icDto.DocumentBody.Invoice.TaxTotal.TaxSubtotal[i]
				taxSubtotalsFlattened +=
					subtotal.TaxableAmount + ";" +
						subtotal.TaxAmount + ";" +
						subtotal.TaxCategory.ID + ";" +
						subtotal.TaxCategory.Percent + ";"
			} else {
				taxSubtotalsFlattened += ";;;;"
			}
		}

		if _, err := f.WriteString(
			icDto.DocumentHeader.SalesInvoiceID + ";" +
				icDto.DocumentHeader.PurchaseInvoiceID + ";" +
				icDto.DocumentBody.Invoice.ID + ";" +
				icDto.DocumentBody.Invoice.IssueDate + ";" +
				icDto.DocumentBody.Invoice.DueDate + ";" +
				icDto.DocumentBody.Invoice.InvoiceTypeCode + ";" +
				icDto.DocumentBody.Invoice.DocumentCurrencyCode + ";" +
				icDto.DocumentBody.Invoice.InvoicePeriod.DescriptionCode + ";" +
				icDto.DocumentBody.Invoice.AccountingSupplierParty.Party.EndpointID + ";" +
				icDto.DocumentBody.Invoice.AccountingSupplierParty.Party.PartyLegalEntity.CompanyID + ";" +
				icDto.DocumentBody.Invoice.AccountingSupplierParty.Party.PartyName.Name + ";" +
				icDto.DocumentBody.Invoice.Delivery.ActualDeliveryDate + ";" +
				icDto.DocumentBody.Invoice.PaymentMeans.PaymentID + ";" +
				icDto.DocumentBody.Invoice.TaxTotal.TaxAmount + ";" +
				taxSubtotalsFlattened +
				icDto.DocumentBody.Invoice.LegalMonetaryTotal.LineExtensionAmount + ";" +
				icDto.DocumentBody.Invoice.LegalMonetaryTotal.TaxExclusiveAmount + ";" +
				icDto.DocumentBody.Invoice.LegalMonetaryTotal.TaxInclusiveAmount + ";" +
				icDto.DocumentBody.Invoice.LegalMonetaryTotal.AllowanceTotalAmount + ";" +
				icDto.DocumentBody.Invoice.LegalMonetaryTotal.PrepaidAmount + ";" +
				icDto.DocumentBody.Invoice.LegalMonetaryTotal.PayableAmount + ";" +
				"\r\n"); err != nil {
			log.Println(err)
		}

		filePath := filepath.Join(utils.FolderPath, "OdgovorProdajaStavke.txt")
		fInvoiceLine, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer fInvoiceLine.Close()

		for _, line := range icDto.DocumentBody.Invoice.InvoiceLine {
			if _, err := fInvoiceLine.WriteString(
				icDto.DocumentHeader.SalesInvoiceID + ";" +
					icDto.DocumentHeader.PurchaseInvoiceID + ";" +
					line.ID + ";" +
					line.InvoicedQuantity + ";" +
					line.LineExtensionAmount + ";" +
					line.AllowanceCharge.ChargeIndicator + ";" +
					line.AllowanceCharge.MultiplierFactorNumeric + ";" +
					line.AllowanceCharge.Amount + ";" +
					line.Item.Name + ";" +
					line.Item.SellersItemIdentification.ID + ";" +
					line.Item.ClassifiedTaxCategory.ID + ";" +
					line.Item.ClassifiedTaxCategory.Percent + ";" +
					line.Price.PriceAmount + ";" +
					"\r\n",
			); err != nil {
				log.Println("Error writing invoice line data:", err)
			}
		}
	}
}

func SalesInvoiceXml(invoiceId string) []byte {

	var body []byte
	var code int
	var status string

	url, err := url.JoinPath(utils.UrlPath, "sales-invoice/xml")
	if err != nil {
		log.Println(err)
	}

	body, code, status = utils.GetQ(url, "invoiceId", invoiceId)

	if code == 200 {
		return body
	} else {
		log.Println(status)
		log.Println(string(body))
		return nil
	}
}
