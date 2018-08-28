package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	inputfile := flag.String("file", "", "Grab File")

	keyfile := flag.String("key", "", "Key File")

	flag.Parse()

	// Hop out if either file is void
	if len(*inputfile) > 0 && len(*keyfile) > 0 {

		fileData, err := ioutil.ReadFile(*inputfile)

		if err != nil {
			fmt.Println("Can't read file:", os.Args[1])
			panic(err)
		}

		keyData, err := ioutil.ReadFile(*keyfile)

		if err != nil {
			fmt.Println("Can't read file:", os.Args[2])
			panic(err)
		}

		_ = Translate(fileData, keyData)
	}
}
