package post

import (
	"einvoice/utils"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// goes through all sales invoices change and return last change/status for each invoice
func PostSalesStatusOnDate(date string) {

	var SalesInvoiceStatusChangeDto1 utils.SalesInvoiceStatusChangeDto = PostSalesInvoiceChanges(date)

	var changes = map[int]utils.Status{}

	for _, v := range SalesInvoiceStatusChangeDto1 {
		if val, ok := changes[v.SalesInvoiceID]; ok {
			// Found , we want status from last event
			if v.EventID > val.EventId {
				changes[v.SalesInvoiceID] = utils.Status{EventId: v.EventID, Status: v.NewInvoiceStatus, Comment: v.Comment, Date: v.Date}
			}
			continue
		}
		changes[v.SalesInvoiceID] = utils.Status{EventId: v.EventID, Status: v.NewInvoiceStatus, Comment: v.Comment, Date: v.Date}
	}

	//log.Printf("%+v\n", changes)

	f, err := os.OpenFile(filepath.Join(utils.FolderPath, "StatusSales.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	//log.Printf("%+v\n", changes)
	for key, value := range changes {
		if _, err := f.WriteString(strconv.Itoa(key) + ";" + value.Status + ";" + value.Comment + ";" + value.Date + ";" + "\r\n"); err != nil {
			log.Println(err)
		}
	}

	defer f.Close()
}
