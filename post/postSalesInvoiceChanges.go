package post

import (
	"einvoice/utils"
	"encoding/json"
	"log"
	"net/url"
)

// called from PostSalesStatusOnDate, this return all changes during given date, i put yesterday as default
func PostSalesInvoiceChanges(date string) utils.SalesInvoiceStatusChangeDto {

	var body []byte
	var code int
	var status string

	url, err := url.JoinPath(utils.UrlPath, "sales-invoice/changes")
	if err != nil {
		log.Println(err)
	}

	body, code, status = utils.PostQ(url, "date", date)

	if code == 200 {
		var SalesInvoiceStatusChangeDto1 utils.SalesInvoiceStatusChangeDto
		if err := json.Unmarshal(body, &SalesInvoiceStatusChangeDto1); err != nil {
			log.Fatalln("Can not unmarshal JSON" + err.Error())
		}
		//log.Printf("%+v\n", SalesInvoiceStatusChangeDto1)
		return SalesInvoiceStatusChangeDto1
	} else {
		log.Fatalln(status)
		log.Fatalln(string(body))
	}
	return nil
}
