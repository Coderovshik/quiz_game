package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type problem struct {
	formula  string
	solution string
}

func parseCSV(f *os.File) (problems []problem) {
	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		problems = append(problems, problem{
			formula:  record[0],
			solution: record[1],
		})
	}

	return
}
