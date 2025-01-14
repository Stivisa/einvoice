package post

import (
	"einvoice/utils"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// goes through all purchase invoices change and return last change/status for each invoice
func PostPurchaseStatusOnDate(date string) {

	var PurchaseInvoiceStatusChangeDto1 utils.PurchaseInvoiceStatusChangeDto = PostPurchaseInvoiceChanges(date)

	var changes = map[int]utils.Status{}

	for _, v := range PurchaseInvoiceStatusChangeDto1 {
		if val, ok := changes[v.PurchaseInvoiceID]; ok {
			// Found , we wont status from last event
			if v.EventID > val.EventId {
				changes[v.PurchaseInvoiceID] = utils.Status{EventId: v.EventID, Status: v.NewInvoiceStatus}
			}
			continue
		}
		changes[v.PurchaseInvoiceID] = utils.Status{EventId: v.EventID, Status: v.NewInvoiceStatus}
	}
	//log.Printf("%+v\n", changes)
	f, err := os.OpenFile(filepath.Join(utils.FolderPath, "StatusPurchase.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	//log.Printf("%+v\n", changes)
	for key, value := range changes {
		if _, err := f.WriteString(strconv.Itoa(key) + ";" + value.Status + "\r\n"); err != nil {
			log.Println(err)
		}
	}

	defer f.Close()
}
