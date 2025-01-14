package post

import (
	"einvoice/utils"
	"encoding/json"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// called from PostAllVatIndividualFromFolder, write to Odgovor.txt with PostSalesInvoiceUbl, as program default behaviour
func PostVatRecordingIndividual(file string, option int) {

	var body []byte
	var code int
	var status string
	var requestId string
	var IndividualVatDto1 utils.IndividualVatDto
	var IndividualVatAddDto1 utils.IndividualVatAddDto
	var IndividualVatDtoV2 utils.IndividualVatDtoV2
	var IndividualVatAddDtoV2 utils.IndividualVatAddDtoV2
	var apiVersion int

	fileName := filepath.Base(file)
	if strings.HasPrefix(fileName, "I2") {
		apiVersion = 2
	} else {
		apiVersion = 1
	}
	log.Println("apiVersion: " + strconv.Itoa(apiVersion))

	jsonFile, err := os.ReadFile(file)
	if err != nil {
		log.Println(err)
	}
	bodyContent := string(jsonFile)
	if bodyContent == "" {
		log.Println("Json is empty")
		return
	}

	if apiVersion == 1 {
		if err := json.Unmarshal([]byte(bodyContent), &IndividualVatAddDto1); err != nil {
			log.Println("Can not unmarshal JSON")
			log.Println(err)
		}
		requestId = IndividualVatAddDto1.DocumentNumber
	} else {
		if err := json.Unmarshal([]byte(bodyContent), &IndividualVatAddDtoV2); err != nil {
			log.Println("Can not unmarshal JSON")
			log.Println(err)
		}
		requestId = IndividualVatAddDtoV2.DocumentNumber
	}

	var urlFull string
	if apiVersion == 1 {
		urlFull, err = url.JoinPath(utils.UrlPath, "vat-recording/individual")
		if err != nil {
			log.Println(err)
		}
	} else {
		urlFull, err = url.JoinPath(utils.UrlPathV2, "vat-recording/individual")
		if err != nil {
			log.Println(err)
		}
	}

	body, code, status = utils.PostQJson(urlFull, bodyContent)

	log.Println(status)

	var f *os.File
	if option == 1 {
		f, err = os.OpenFile(filepath.Join(utils.FolderPath, "Odgovor.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
	} else {
		f, err = os.OpenFile(filepath.Join(utils.VatFolderPath, "OdgovorPdv.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
	}

	if code == 200 {
		if apiVersion == 1 {
			if err := json.Unmarshal(body, &IndividualVatDto1); err != nil {
				log.Println("Can not unmarshal JSON:" + err.Error())
			}
			if _, err := f.WriteString(requestId + ";" + strconv.Itoa(IndividualVatDto1.IndividualVatID) + ";vat;vat;" + "\r\n"); err != nil {
				log.Println(err)
			}
		} else {
			if err := json.Unmarshal(body, &IndividualVatDtoV2); err != nil {
				log.Println("Can not unmarshal JSON:" + err.Error())
			}
			if _, err := f.WriteString(requestId + ";" + strconv.Itoa(IndividualVatDtoV2.IndividualVatID) + ";vat;vat;" + "\r\n"); err != nil {
				log.Println(err)
			}
		}

	} else {
		var ErrorInvoice1 utils.ErrorInvoice
		if err := json.Unmarshal(body, &ErrorInvoice1); err != nil {
			log.Println("Can not unmarshal JSON")
		}
		log.Println(ErrorInvoice1)

		if _, err := f.WriteString(requestId + ";;vat;vat;" + ErrorInvoice1.Message + "\r\n"); err != nil {
			log.Println(err)
		}
	}
	defer f.Close()
}
