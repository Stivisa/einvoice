package get

import (
	"einvoice/utils"
	"encoding/json"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func GetUnitMeasures() {

	var body []byte
	var code int
	var status string

	url, err := url.JoinPath(utils.UrlPath, "get-unit-measures")
	if err != nil {
		log.Println(err)
	}

	body, code, status = utils.GetQ(url)

	if code == 200 {
		var Unit1 []utils.UnitMeasures
		if err := json.Unmarshal(body, &Unit1); err != nil {
			log.Println(err)
		}

		f, err := os.OpenFile(filepath.Join(utils.FolderPath, "JediniceMere.txt"), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Println(err)
		}

		for _, unit := range Unit1 {
			if _, err := f.WriteString(unit.Code + ";" + unit.Symbol + ";" + unit.NameEng + ";" + unit.NameSrbLtn + ";" + unit.NameSrbCyr + ";" + strconv.FormatBool(unit.IsOnShortList) + ";" + "\r\n"); err != nil {
				log.Println(err)
			}
		}

	} else {
		log.Println(status)
		log.Println(string(body))
	}
}
