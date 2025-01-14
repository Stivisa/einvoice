package post

import (
	"einvoice/utils"
	"encoding/json"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

// if we skip status then returns all id, useless because we dont know status
func PostPurchaseInvoiceIds(statusInvoice string, dateFrom string, dateTo string) {

	var body []byte
	var code int
	var status string

	url, err := url.JoinPath(utils.UrlPath, "purchase-invoice/ids")
	if err != nil {
		log.Println(err)
	}

	body, code, status = utils.PostQ(url, "status", statusInvoice, "dateFrom", dateFrom, "dateTo", dateTo)

	if code == 200 {
		var PurchaseInvoicesDto1 utils.PurchaseInvoicesDto
		if err := json.Unmarshal(body, &PurchaseInvoicesDto1); err != nil {
			log.Fatalln("Can not unmarshal JSON" + err.Error())
		}
		log.Printf("%+v\n", PurchaseInvoicesDto1.PurchaseInvoiceIds)

		statustxt := statusInvoice + "Purchase.txt"
		f, err := os.OpenFile(filepath.Join(utils.FolderPath, statustxt), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Println(err)
		}

		for _, purchaseId := range PurchaseInvoicesDto1.PurchaseInvoiceIds {
			if _, err := f.WriteString(strconv.Itoa(purchaseId) + "\r\n"); err != nil {
				log.Println(err)
			}
		}
	} else {
		log.Fatalln(status)
		log.Fatalln(string(body))
	}
}
