package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input file")
		return
	}

	freqChangesData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	var freqChangeDataStr = string(freqChangesData)
	freqChangeDataStr = strings.Replace(freqChangeDataStr, ",", "\n", -1)

	strChanges := strings.Fields(freqChangeDataStr)
	var freqChanges []int

	for _, change := range strChanges {
		change, err := strconv.Atoi(change)
		if err != nil {
			panic(err)
		}

		freqChanges = append(freqChanges, change)
	}

	fmt.Println(Calibrate(freqChanges))
}
