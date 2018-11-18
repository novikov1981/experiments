package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	voc := []string{"A", "C", "G", "T"}

	var dna string

	fmt.Print("Введите последовательность ДНК: ")
	fmt.Fscan(os.Stdin, &dna)

	lenDNA := len(dna)

	fmt.Printf("Ваша последовательность %s\n", dna)
	fmt.Printf("Длинна последовательности %d\n", lenDNA)

	dnaU := strings.ToUpper(dna)

	count := 0

	for _, o := range voc {
		c := strings.Count(dnaU, o)
		count += c
		fmt.Printf("Колличество d%s: %d\n", o, c)
	}

	pos := len(dnaU)
	var rdna string

	for len(rdna) < len(dnaU) {
		num := 1
		for _, r := range dnaU {
			s := string(r)
			if num == pos {
				rdna += s
				pos -= 1
				if len(rdna) == len(dnaU) {
					fmt.Println("Oбратная последовательность", rdna)
				}
			}
			num += 1
		}
	}

	if lenDNA > count {

		num := 1

		for _, r := range dnaU {
			count := 0
			s := string(r)
			for _, o := range voc {
				c := strings.Count(s, o)
				count += c
			}
			if count == 0 {
				fmt.Printf("Недопустимый символ №%d: %s\n", num, s)
				//fmt.Printf("Символ d%s: %d\n", num, s)
			}
			num += 1
		}

		fmt.Errorf("ОШИБКА последовательность содержит %d недопустимых символов", lenDNA-count)
		os.Exit(1)
	}

	b := "aaaa"

	work(b)

}

func work(s string) {

}
