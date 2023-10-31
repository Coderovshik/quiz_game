package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func quiz(problems *[]problem) {
	correctCount := 0
	r := bufio.NewReader(os.Stdin)

	for idx, p := range *problems {
		fmt.Printf("Problem #%d: %s = ", idx+1, p.formula)
		answ, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("failed to read answer")
			continue
		}

		answ = strings.TrimSuffix(answ, "\r\n")
		if p.solution == answ {
			correctCount++
		}
	}

	fmt.Printf("You scored %d out of %d", correctCount, len(*problems))
}
