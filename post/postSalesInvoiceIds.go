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
func PostSalesInvoiceIds(statusInvoice string, dateFrom string, dateTo string) {

	var body []byte
	var code int
	var status string

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
		log.Printf("%+v\n", SalesInvoicesDto1.SalesInvoiceIds)

		statustxt := statusInvoice + "Sales.txt"
		f, err := os.OpenFile(filepath.Join(utils.FolderPath, statustxt), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Println(err)
		}

		for _, salesId := range SalesInvoicesDto1.SalesInvoiceIds {
			if _, err := f.WriteString(strconv.Itoa(salesId) + "\r\n"); err != nil {
				log.Println(err)
			}
		}

	} else {
		log.Println(status)
		log.Fatalln(string(body))
	}
}
