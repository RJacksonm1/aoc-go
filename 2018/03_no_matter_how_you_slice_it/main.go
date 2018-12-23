package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"./garmentfactory"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input file")
		return
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fabric := garmentfactory.Fabric{}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		claim, err := garmentfactory.NewClaimFromDeclaration(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		fabric.AddClaim(claim)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fabric.Render()

	fmt.Printf("Fabric is at least %d x %d; %d square inches are in conflict\n",
		fabric.Width(),
		fabric.Height(),
		len(fabric.Conflicts()),
	)
}
