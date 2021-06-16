package main

import (
	"fmt"
	"github.com/linkedin/goavro/v2"
	"os"
)

func main() {
	// First element in os.Args is always the program name,
	// So we need at least 2 arguments to have a file name argument.
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	dh, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	// data is the file content, you can use it
	ocfr, err := goavro.NewOCFReader(dh)
	if err != nil {
		fmt.Println("NewOCFReader failed.")
		panic(err)
	}
	fmt.Println(ocfr.Codec().Schema())
}
