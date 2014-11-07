package main

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// type Token struct {
//     Type     TokenType
//     DataAtom atom.Atom
//     Data     string
//     Attr     []Attribute
// }
//
// type Attribute struct {
//     Namespace, Key, Val string
// }

func parseHtml(r io.Reader) {
	d := html.NewTokenizer(r)
	for {
		// token type
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			println("")
			return
		}
		token := d.Token()
		switch tokenType {
		case html.StartTagToken: // <tag>
			print("<" + token.Data)

			for _, attr := range token.Attr {
				print(" " + attr.Key + "=" + "\"" + attr.Val + "\"")
			}
			print(">")

		case html.TextToken: // text between start and end tag
			print(token.Data)
		case html.EndTagToken: // </tag>
			print("</" + token.Data + ">")
		case html.SelfClosingTagToken: // <tag/>
			print("<" + token.Data + "/>")

		}
	}

}

func main() {
	fmt.Println(os.Args)

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(b)

	parseHtml(r)

	println(string(b))

}
