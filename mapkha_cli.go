package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
var cpuprofile string
var memprofile string

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
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("could not read input:", err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		fmt.Println(strings.Join(wordcut.Segment(scanner.Text()), "|"))
	}
}
