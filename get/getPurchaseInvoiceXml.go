package get

import (
	"einvoice/utils"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

func GetPurchaseInvoiceXml(invoiceId string) {
	var body []byte
	var code int
	var status string

	url, err := url.JoinPath(utils.UrlPath, "purchase-invoice/xml")
	if err != nil {
		log.Println(err)
	}

	body, code, status = utils.GetQ(url, "invoiceId", invoiceId)

	if code == 200 {
		strXml := string(body)

		purchaseXmlFolder := filepath.Join(utils.FolderPath, "PurchaseXML")

		if _, err := os.Stat(purchaseXmlFolder); os.IsNotExist(err) {
			err := os.Mkdir(purchaseXmlFolder, 0777)
			if err != nil {
				log.Println(err)
			}
		}

		f, err := os.OpenFile(filepath.Join(purchaseXmlFolder, invoiceId+".xml"), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Println(err)
		}

		if _, err := f.WriteString(strXml); err != nil {
			log.Println(err)
		}
	} else {
		log.Fatalln(status)
		log.Fatalln(string(body))
	}
}
