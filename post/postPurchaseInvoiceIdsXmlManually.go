package post

import (
	"bufio"
	"einvoice/utils"
	"encoding/xml"
	"log"
	"os"
	"path/filepath"
)

func PostSalesInvoiceIdsXmlManually() {

	file, err := os.Open(filepath.Join(utils.FolderPath, "RucnoUnos.txt"))
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	f, err := os.OpenFile(filepath.Join(utils.FolderPath, "OdgovorRucno.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		item := scanner.Text()
		//log.Println(item)

		xmlBody := SalesInvoiceXml(item)

		if len(xmlBody) == 0 {
			//log.Println("Byte array is empty.")
			if _, err := f.WriteString(
				item + ";" + ";" + ";" +
					"\r\n"); err != nil {
				log.Println(err)
			}
		} else {
			var icDto utils.InvoiceCustomDto
			if err := xml.Unmarshal(xmlBody, &icDto); err != nil {
				log.Println("Can not unmarshal JSON" + err.Error())
			}

			//log.Println(icDto.DocumentHeader.SalesInvoiceID + ";" + icDto.DocumentHeader.PurchaseInvoiceID + ";" + icDto.DocumentBody.Invoice.ID + ";")

			if _, err := f.WriteString(
				icDto.DocumentHeader.SalesInvoiceID + ";" + icDto.DocumentHeader.PurchaseInvoiceID + ";" + icDto.DocumentBody.Invoice.ID + ";" +
					"\r\n"); err != nil {
				log.Println(err)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}
}
