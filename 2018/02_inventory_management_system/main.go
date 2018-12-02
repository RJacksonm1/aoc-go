package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input file")
		return
	}

	boxIdsData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	boxIds := strings.Fields(string(boxIdsData))
	var boxes []Box

	for _, id := range boxIds {
		boxes = append(boxes, Box{id: id})
	}

	fabricBoxA, fabricBoxB := FindPrototypeFabrics(boxes)
	var fabricBoxCommonLetters string

	for i, r := range fabricBoxA.id {
		if fabricBoxA.id[i] == fabricBoxB.id[i] {
			fabricBoxCommonLetters += string(r)
		}
	}

	fmt.Println(fabricBoxCommonLetters)
}
