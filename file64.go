package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	MidLineFormat = "%s\n"
	EndLineFormat = "%s\n"
)

func main() {
	quotes := flag.Bool("quotes", false, "Print content as Go string literal")
	flag.Parse()

	if *quotes {
		MidLineFormat = "\"%s\"+\n"
		EndLineFormat = "\"%s\"\n"
	}

	filename := flag.Arg(0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	src, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	dest := base64.StdEncoding.EncodeToString(src)
	for ; len(dest) > 80; dest = dest[80:] {
		fmt.Printf(MidLineFormat, dest[:80])
	}
	fmt.Printf(EndLineFormat, dest)
}
