package get

import (
	"einvoice/utils"
	"encoding/xml"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func GetSalesInvoiceXml(invoiceId string) {

	var body []byte
	var code int
	var status string

	url, err := url.JoinPath(utils.UrlPath, "sales-invoice/xml")
	if err != nil {
		log.Println(err)
	}

	body, code, status = utils.GetQ(url, "invoiceId", invoiceId)

	if code == 200 {
		strXml := string(body)

		salesXmlFolder := filepath.Join(utils.FolderPath, "SalesXML")

		if _, err := os.Stat(salesXmlFolder); os.IsNotExist(err) {
			err := os.Mkdir(salesXmlFolder, 0777)
			if err != nil {
				log.Println(err)
			}
		}

		f, err := os.OpenFile(filepath.Join(salesXmlFolder, invoiceId+".xml"), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Println(err)
		}

		if _, err := f.WriteString(strXml); err != nil {
			log.Println(err)
		}

		var InvoiceCustomDto1 utils.InvoiceCustomDto
		if err := xml.Unmarshal(body, &InvoiceCustomDto1); err != nil {
			log.Fatalln("Can not unmarshal JSON" + err.Error())
		}

		e := excelize.NewFile()

		e.SetCellValue("Sheet1", "A1", "SalesInvoiceID")
		e.SetCellValue("Sheet1", "A2", InvoiceCustomDto1.DocumentHeader.SalesInvoiceID)
		e.SetCellValue("Sheet1", "B1", "PurchaseInvoiceID")
		e.SetCellValue("Sheet1", "B2", InvoiceCustomDto1.DocumentHeader.PurchaseInvoiceID)
		e.SetCellValue("Sheet1", "C1", "CreationDate")
		e.SetCellValue("Sheet1", "C2", InvoiceCustomDto1.DocumentHeader.CreationDate)
		e.SetCellValue("Sheet1", "D1", "SendingDate")
		e.SetCellValue("Sheet1", "D2", InvoiceCustomDto1.DocumentHeader.SendingDate)

		e.SetCellValue("Sheet1", "E1", "BrojRacuna")
		e.SetCellValue("Sheet1", "E2", InvoiceCustomDto1.DocumentBody.Invoice.ID)
		e.SetCellValue("Sheet1", "F1", "DatumIzdavanjaRacuna")
		e.SetCellValue("Sheet1", "F2", InvoiceCustomDto1.DocumentBody.Invoice.IssueDate)
		e.SetCellValue("Sheet1", "G1", "DatumDospecaRacuna")
		e.SetCellValue("Sheet1", "G2", InvoiceCustomDto1.DocumentBody.Invoice.DueDate)
		e.SetCellValue("Sheet1", "H1", "TipDokumenta")
		e.SetCellValue("Sheet1", "H2", InvoiceCustomDto1.DocumentBody.Invoice.InvoiceTypeCode)

		if err := e.SaveAs(filepath.Join(salesXmlFolder, invoiceId+".xlsx")); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatalln(status)
		log.Fatalln(string(body))
	}
}
