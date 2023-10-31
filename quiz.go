package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func logScore(score, total int) {
	fmt.Printf("\nYou scored %d out of %d", score, total)
}

func quiz(problems []problem, limit int) {
	correctCount := 0
	r := bufio.NewReader(os.Stdin)

running:
	for idx, p := range problems {
		fmt.Printf("Problem #%d: %s = ", idx+1, p.formula)

		timeoutch := make(chan struct{})
		answch := make(chan string)

		go func() {
			time.Sleep(time.Duration(limit) * time.Second)
			close(timeoutch)
		}()

		go func() {
			answ, _ := r.ReadString('\n')
			answch <- answ
		}()

		select {
		case answ := <-answch:
			answ = strings.TrimSuffix(answ, "\r\n")
			if p.solution == answ {
				correctCount++
			}
			close(answch)
		case <-timeoutch:
			break running
		}
	}

	logScore(correctCount, len(problems))
}
