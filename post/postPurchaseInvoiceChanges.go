package post

import (
	"einvoice/utils"
	"encoding/json"
	"log"
	"net/url"
)

// called from PostPurchaseStatusOnDate, this return all changes during given date, i put yesterday as default
func PostPurchaseInvoiceChanges(date string) utils.PurchaseInvoiceStatusChangeDto {

	var body []byte
	var code int
	var status string

	url, err := url.JoinPath(utils.UrlPath, "purchase-invoice/changes")
	if err != nil {
		log.Println(err)
	}

	body, code, status = utils.PostQ(url, "date", date)

	if code == 200 {
		var PurchaseInvoiceStatusChangeDto1 utils.PurchaseInvoiceStatusChangeDto
		if err := json.Unmarshal(body, &PurchaseInvoiceStatusChangeDto1); err != nil {
			log.Fatalln("Can not unmarshal JSON" + err.Error())
		}
		log.Printf("%+v\n", PurchaseInvoiceStatusChangeDto1)
		return PurchaseInvoiceStatusChangeDto1
	} else {
		log.Fatalln(status)
		log.Fatalln(string(body))
	}
	return nil
}
