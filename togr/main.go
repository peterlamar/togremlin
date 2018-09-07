package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/peterlamar/togremlin/fileutil"
	"github.com/peterlamar/togremlin/gremlin"
)

type commandLineOptions struct {
	SourceFile *string
	KeyFile    *string
}

func newCommandLineOptions() *commandLineOptions {
	cmdLineOpts := commandLineOptions{}
	cmdLineOpts.SourceFile = flag.String("source", "", "Filename to retrieve data from")
	cmdLineOpts.KeyFile = flag.String("key", "", "Filename to retreive graph key information from")

	flag.Usage = cmdLineOpts.Usage
	flag.Parse()

	return &cmdLineOpts
}

func (clo *commandLineOptions) GetSourceFile() string {
	return *clo.SourceFile
}
func (clo *commandLineOptions) GetKeyFile() string {
	return *clo.KeyFile
}

func (clo *commandLineOptions) Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "  %s [-key pathtokeyfile] [-source pathtosourcefile] \n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
	flag.PrintDefaults()
}

func main() {
	cmdLineOpts := newCommandLineOptions()

	if cmdLineOpts.GetSourceFile() == "" {
		cmdLineOpts.Usage()
		return
	}

	fileData, err := ioutil.ReadFile(cmdLineOpts.GetSourceFile())
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	if cmdLineOpts.GetKeyFile() == "" {

		graphData := gremlin.Translate(fileData)

		// Commit data to disk in multiple files
		fileutil.WriteNodes(graphData)
	} else {
		gremlinKeys, err := ioutil.ReadFile(cmdLineOpts.GetKeyFile())
		if err != nil {
			fmt.Println("Can't read file:", os.Args[1])
			panic(err)
		}

		graphData := gremlin.TranslateWithKey(fileData, gremlinKeys)

		// Commit data to disk in multiple files
		fileutil.WriteNodes(graphData)
	}
}
