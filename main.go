package main

import "flag"

func main() {

	inputfile := flag.String("file", "", "Grab File")
	keyfile := flag.String("key", "", "Key File")

	flag.Parse()
	if len(*inputfile) > 0 && len(*keyfile) > 0 {
	}
}
