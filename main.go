package main

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	n, err := html.Parse(strings.NewReader(string(b)))
	if err != nil {
		log.Fatalf("Parse error: %s", err)
	}
	var buf bytes.Buffer
	if err := html.Render(&buf, n); err != nil {
		log.Fatalf("Render error: %s", err)
	}
	fmt.Println(buf.String())

}
