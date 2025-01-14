package post

import (
	"bufio"
	"einvoice/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func PostSalesInvoiceUbl(file string) {

	var body []byte
	var code int
	var status string
	var requestId string
	var toCir string
	var MiniInvoiceDto1 utils.MiniInvoiceDto

	xmlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
	}
	bodyContent := string(xmlFile)
	if bodyContent == "" {
		log.Println("Xml is empty")
		return
	}

	//jer 7.konacna faktura ima extension prvo, pa ne treba citati prvi <cbc:ID>
	bodyContentAfter := utils.After(bodyContent, "<cbc:CustomizationID>")
	requestId = utils.Between(bodyContentAfter, "<cbc:ID>", "</cbc:ID>")
	//log.Println("reqId: " + requestId)

	fileName := filepath.Base(file)
	if strings.HasPrefix(fileName, "B") {
		toCir = "Yes"
	} else {
		toCir = "No"
	}
	log.Println("toCir: " + toCir)

	url, err := url.JoinPath(utils.UrlPath, "sales-invoice/ubl")
	if err != nil {
		log.Println(err)
	}

	var pdfError string
	pdfAttachmentPathXmlTag(&bodyContent, &pdfError)
	//pdfAttachmentFolder(file, &bodyContent)
	//log.Println("Body: " + bodyContent)

	f, err := os.OpenFile(filepath.Join(utils.FolderPath, "Odgovor.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if pdfError != "" {
		if _, err := f.WriteString(requestId + ";" + ";" + ";" + ";" + pdfError + "\r\n"); err != nil {
			log.Println(err)
		}
		return
	}

	body, code, status = utils.PostQXml(url, bodyContent, "requestId", requestId, "sendToCir", toCir)

	log.Println(status)

	if code == 200 {
		if err := json.Unmarshal(body, &MiniInvoiceDto1); err != nil {
			log.Println("Can not unmarshal JSON")
		}
		//log.Println(string(body))
		//log.Println(MiniInvoiceDto1)

		if _, err := f.WriteString(requestId + ";" + strconv.Itoa(MiniInvoiceDto1.InvoiceID) + ";" + strconv.Itoa(MiniInvoiceDto1.PurchaseInvoiceID) + ";" + strconv.Itoa(MiniInvoiceDto1.SalesInvoiceID) + ";" + "\r\n"); err != nil {
			log.Println(err)
		}

	} else {
		var ErrorInvoice1 utils.ErrorInvoice
		if err := json.Unmarshal(body, &ErrorInvoice1); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		log.Println(ErrorInvoice1)

		if _, err := f.WriteString(requestId + ";" + ";" + ";" + ";" + ErrorInvoice1.Message + "\r\n"); err != nil {
			log.Println(err)
		}
	}
}

// xml treba da sadrzi tag element <PRILOG>putanjaPriloga</PRILOG>
func pdfAttachmentPathXmlTag(bodyContent *string, pdfError *string) {
	maxPdf := 0
	var fullAttachmentString string
	bodyContentBeforeAsp := utils.Before(*bodyContent, "<cac:AccountingSupplierParty>")
	bodyContentAfterAsp := utils.After(*bodyContent, "<cac:AccountingSupplierParty>")
	for strings.Contains(bodyContentBeforeAsp, "PRILOG") {
		maxPdf++
		//'too many files' ce sef odgovoriti
		if maxPdf > 3 {
			log.Println("Maximum 3 pdf attachments! Each with a maximum size of 25MB.")
		}

		attachmentPath := utils.Between(bodyContentBeforeAsp, "<PRILOG>", "</PRILOG>")
		log.Println("Attachment file: " + attachmentPath)
		pdfName := filepath.Base(attachmentPath)
		pdfNameWithoutExt := pdfName[:len(pdfName)-len(filepath.Ext(pdfName))]

		pdfFile, err := os.Open(attachmentPath)
		if err != nil {
			//need to report this error (in case file doesn't exists anymore)
			//if not, invoice is valid and will be sent without attachments
			*pdfError = err.Error()
			log.Println(err.Error())
		}
		defer pdfFile.Close()

		stat, err := pdfFile.Stat()
		if err != nil {
			log.Println(err)
			return
		}
		bytes := stat.Size()
		megabytes := (int)(bytes / 1024 / 1024)
		log.Println("File size around: " + strconv.Itoa(megabytes) + "MB")
		if megabytes > 10 {
			//za veci fajl sef vrati 200 OK sa random html kodom, rezultat u odg bude 191;0;0;0; i faktura nije zavedena na sefu
			*pdfError = "Maximum size of one attachment file is 10MB"
			log.Println(*pdfError)
			return
		}

		pdfRreader := bufio.NewReader(pdfFile)
		pdfContent, _ := io.ReadAll(pdfRreader)
		pdfEncoded := base64.StdEncoding.EncodeToString(pdfContent)

		attachmentString := strings.Replace(utils.AttachmentPart, "NAZIV", pdfNameWithoutExt, 1)
		attachmentString = strings.Replace(attachmentString, "PRILOG", pdfEncoded, 1)
		fullAttachmentString += attachmentString

		bodyContentBeforeAtt := utils.Before(bodyContentBeforeAsp, "<PRILOG>")
		bodyContentAfterAtt := utils.After(bodyContentBeforeAsp, "</PRILOG>")
		bodyContentBeforeAsp = bodyContentBeforeAtt + bodyContentAfterAtt
		//log.Println(bodyContentBeforeAsp)
	}
	//log.Println("Attachment: " + fullAttachmentString)
	*bodyContent = bodyContentBeforeAsp + fullAttachmentString + "<cac:AccountingSupplierParty>" + bodyContentAfterAsp
}
