package get

import (
	"einvoice/utils"
	"encoding/json"
	"log"
	"net/url"
)

func GetPurchaseInvoice(invoiceId string) {
	var body []byte
	var code int
	var status string

	url, err := url.JoinPath(utils.UrlPath, "purchase-invoice")
	if err != nil {
		log.Println(err)
	}

	body, code, status = utils.GetQ(url, "invoiceId", invoiceId)

	if code == 200 {
		var SimpleInvoiceDto1 utils.SimpleInvoiceDto
		if err := json.Unmarshal(body, &SimpleInvoiceDto1); err != nil {
			log.Fatalln("Can not unmarshal JSON" + err.Error())
		}
		log.Printf("%+v\n", SimpleInvoiceDto1)
	} else {
		log.Fatalln(status)
		log.Fatalln(string(body))
	}
}
