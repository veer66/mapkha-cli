package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	m "github.com/veer66/mapkha"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var dixPath string

func init() {
	flag.StringVar(&dixPath, "dix", "", "Dictionary path")
}

func main() {
	flag.Parse()

	var dict *m.Dict
	var e error
	if dixPath == "" {
		dict, e = m.LoadDefaultDict()
	} else {
		dict, e = m.LoadDict(dixPath)
	}
	check(e)
	wordcut := m.NewWordcut(dict)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(strings.Join(wordcut.Segment(scanner.Text()), "|"))
	}
}
