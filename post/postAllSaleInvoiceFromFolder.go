package post

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// iterates through folder and for each xml call PostSalesInvoiceUbl
func PostAllSalesInvoiceFromFolder(folder string) {
	iterate(folder)
}

func iterate(currentDirectory string) {
	sepCount := strings.Count(currentDirectory, string(os.PathSeparator))
	//log.Println("sepCount " + strconv.Itoa(sepCount))
	var warn = 1
	filepath.Walk(currentDirectory, func(path string, info os.FileInfo, err error) error {
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

		if filepath.Ext(path) == ".xml" {
			warn = 0
			log.Println("File Name: " + path)
			PostSalesInvoiceUbl(path)
			err := os.Remove(path)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
		return nil
	})
	if warn == 1 {
		log.Println("No xml file in folder!")
	}
}
