package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/andig/gosml"
)

const (
	PRINTRANGE = 46
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	files, err := filepath.Glob("/Users/andig/htdocs/libsml-testing/*.bin")
	check(err)

	for _, f := range files {
		fmt.Println(f)

		f, err := os.Open(f)
		check(err)

		r := bufio.NewReader(f)

		for {
			buf, err := sml.TransportRead(r)
			if err == io.EOF {
				break
			}
			check(err)

			// parse without escape sequence/ begin/end marker
			messages, err := sml.FileParse(buf[8 : len(buf)-16])

			for _, msg := range messages {
				sml.PrintMessage(msg)
			}

			if err != nil {
				fmt.Printf("%+v\n", err)
				goto nextfile
			}
		}
	nextfile:
		// return
	}
}
