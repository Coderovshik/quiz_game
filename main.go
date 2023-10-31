package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	csv := flag.String("csv", "problems.csv",
		"a csv file in the format of 'question,answer' (default \"problems.csv\")")
	limit := flag.Int("limit", 30,
		"the time limit for the quiz in seconds (default 30)")
	flag.Parse()

	f, err := os.Open(*csv)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("File %s does not exist", *csv)
		}

		log.Fatalf("Cannot open %s", *csv)
	}

	problems := parseCSV(f)
	quiz(problems, *limit)
}
