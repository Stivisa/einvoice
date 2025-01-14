package post

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// iterates through folder and for each json call PostVatRecordingIndividual for option 0 (OdgovorPd.txt) or (Odgovor.txt) for option 1 as part of program default behaviour
func PostAllVatIndividualFromFolder(folder string, option int) {
	iterateVatIndividual(folder, option)
}

func iterateVatIndividual(path string, option int) {
	sepCount := strings.Count(path, string(os.PathSeparator))
	//log.Println("sepCount " + strconv.Itoa(sepCount))
	var warn = 1
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		//restrict to not go in subfolders
		maxDepth := 0
		if info.IsDir() && strings.Count(path, string(os.PathSeparator)) > sepCount+maxDepth {
			//log.Println("skip", path)
			return filepath.SkipDir
		}
		//log.Println("Visited : " + path)

		if filepath.Ext(path) == ".json" {
			warn = 0
			log.Println("File Name: " + path)

			PostVatRecordingIndividual(path, option)

			err := os.Remove(path)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
		return nil
	})
	if warn == 1 {
		log.Println("No json file in folder!")
	}
}
