package fileutil

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"unsafe"
)

// prefix to pull an value from a node to an edge
const fileOutputDir = "output"

// WriteNodes writes the gremlin structure to various individual files
func WriteNodes(fileData map[string][]map[string]interface{}) {

	if unsafe.Sizeof(fileData) == 0 {
		log.Fatal("Failed to create json extract data")
	}

	// Create output dir
	newpath := filepath.Join("./", fileOutputDir)
	os.MkdirAll(newpath, os.ModePerm)

	for name, fileData := range fileData {

		filename := name + ".json"
		filename = filepath.Join(newpath, filename)

		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		jsonBytes, _ := json.Marshal(fileData)
		if _, err := f.Write(jsonBytes); err != nil {
			log.Fatal(err)
		}
	}
}
