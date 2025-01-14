package get

import (
	"einvoice/utils"
	"encoding/json"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

func GetAllCompanies() {
	var body []byte
	var code int
	var status string

	url, err := url.JoinPath(utils.UrlPath, "getAllCompanies")
	if err != nil {
		log.Println(err)
	}

	body, code, status = utils.GetQ(url)

	if code == 200 {
		var Company1 []utils.Company
		if err := json.Unmarshal(body, &Company1); err != nil {
			log.Println(err)
		}

		f, err := os.OpenFile(filepath.Join(utils.FolderPath, "Firme.txt"), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Println(err)
		}

		for _, company := range Company1 {
			if _, err := f.WriteString(company.BugetCompanyNumber + ";" + company.RegistrationCode + ";" + company.VatRegistrationCode + ";" + company.Name + ";" + "\r\n"); err != nil {
				log.Println(err)
			}
		}

	} else {
		log.Println(status)
		log.Println(string(body))
	}
}
