package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {

	var encodeURL bool
	flag.BoolVar(&encodeURL, "e", false, "Encodes the given URL/String")
	flag.Parse()

	paramStart := 1
	if encodeURL {
		paramStart = 2
	}

	if len(os.Args) <= paramStart {
		log.Fatal("Expected at least one parameter")
		flag.Usage()
		return
	}

	params := strings.Join(os.Args[paramStart:], " ")

	if encodeURL {
		fmt.Println(encode(params))
		return
	}

	ret, err := decode(params)
	if err != nil {
		log.Fatal(err)
		flag.Usage()
		return
	}

	fmt.Println(ret)
}

func decode(fullurl string) (string, error) {

	return url.PathUnescape(fullurl)

}

func encode(fullurl string) string {

	return url.QueryEscape(fullurl)

}
