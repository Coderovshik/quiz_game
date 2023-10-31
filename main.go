package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("CSV file is not provided")
	}
	csv := os.Args[1]
	f, err := os.Open(csv)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("File %s does not exist", csv)
		}

		log.Fatalf("Cannot open %s", csv)
	}

	problems := parseCSV(f)
	quiz(&problems)
}
