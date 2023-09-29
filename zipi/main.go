package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func init() {
	log.SetPrefix("Zip: ")
}
func main() {
	log.Println("Program started")
	zipiFolder := "archive.zip"
	file := "db.json"
	ziping(zipiFolder, file)
	log.Println("the file has being zipped")
}
func ziping(zipiFolder string, file string) {
	archive, err := os.Create(zipiFolder)
	if err != nil {
		log.Panic("error creating zip file")
	}
	defer archive.Close()
	// open a file to zip
	file1, err := os.Open(file)
	if err != nil {
		log.Panic("error opening the file")
	}
	defer file1.Close()
	// add file to archive
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()
	w, err := zipWriter.Create(file)
	if err != nil {
		log.Panic("error creating zip file in the archive")
	}
	if _, err := io.Copy(w, file1); err != nil {
		log.Panic("error writting the file to the zip folder")
	}
}
