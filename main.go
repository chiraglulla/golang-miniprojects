package main

import (
	"os"
	"fmt"
	"strings"

	"golang-miniprojects/parser"
)


func main() {
	exampleHtml, err := os.ReadFile("ex4.html")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	r := strings.NewReader(string(exampleHtml))
	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
