package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	log.SetFlags(0)

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading from stdin: %v", err)
	}

	if err := clipboard.WriteAll(string(bytes)); err != nil {
		log.Fatalf("error writing to clipboard: %v", err)
	}
}
