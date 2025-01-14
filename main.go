package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"einvoice/cmd"

	"einvoice/utils"

	"github.com/go-ini/ini"
)

func setupLogging(currentDirectory string) {
	errorFile, err := os.OpenFile(filepath.Join(currentDirectory, "Error.txt"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)
	}
	mw := io.MultiWriter(os.Stdout, errorFile)
	log.SetOutput(mw)
}

func walkThroughFiles(currentDirectory string, foundIniFile *bool) error {
	return filepath.Walk(currentDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalln(err.Error())
		}

		// Restrict to not go in subfolders
		maxDepth := 0
		if info.IsDir() && strings.Count(path, string(os.PathSeparator)) > strings.Count(currentDirectory, string(os.PathSeparator))+maxDepth {
			return filepath.SkipDir
		}

		// Check if file is an .ini file
		if filepath.Ext(path) == ".ini" {
			*foundIniFile = true
			return processIniFile(path)
		}
		return nil
	})
}

func processIniFile(path string) error {
	log.Println(path)

	// Load .ini file
	cfgs, err := ini.Load(path)
	if err != nil {
		log.Fatalln(err)
	}

	// Set up configuration
	utils.ApiKey = cfgs.Section("eFaktura").Key("ApiKey").Value()
	if len(utils.ApiKey) == 0 {
		log.Fatalln("U podešavanjima fali ApiKey polje!")
	}
	utils.FolderPath, _ = filepath.Abs(cfgs.Section("eFaktura").Key("FolderPath").Value())
	if len(cfgs.Section("eFaktura").Key("FolderPath").Value()) == 0 {
		log.Fatalln("U podešavanjima fali FolderPath polje!")
	}
	utils.VatFolderPath, _ = filepath.Abs(cfgs.Section("eFaktura").Key("VatFolderPath").Value())
	if len(cfgs.Section("eFaktura").Key("VatFolderPath").Value()) == 0 {
		log.Fatalln("U podešavanjima fali VatFolderPath polje!")
	}

	return nil
}

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	setupLogging(currentDirectory)

	foundIniFile := false
	err = walkThroughFiles(currentDirectory, &foundIniFile)
	if err != nil {
		log.Fatal(err)
	}

	if !foundIniFile {
		log.Fatalln("No .ini file!")
	}

	cmd.Execute()
}
