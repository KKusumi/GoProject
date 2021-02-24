package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"unicode/utf8"
)

func NewRuneScanner(r io.Reader) * RuneScanner {
	return &RuneScanner{r: r}
}

func main() {
	s:= NewRuneScanner(strings.NewReader("HelloWorld!"))
	for {
		r, err := s.Scan()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%c\n", r)
	}
}