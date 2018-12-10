package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	voc := []string{"A", "C", "G", "T"} //Слайс

	var dna string

	fmt.Print("Введите последовательность ДНК: ")
	fmt.Fscan(os.Stdin, &dna) // чтение с консоли

	lenDNA := len(dna)           //подсчет колличества символов для строки "dna"
	dnaU := strings.ToUpper(dna) //Функция библиотеки string делает всю строку в верхнем регистре

	if lenDNA > 1 {
		measuring(voc, dnaU)
	}

	fmt.Printf("Ваша последовательность %s\n", dnaU) // %s об\n
	fmt.Printf("Длинна последовательности %d\n", lenDNA)

	//for _, o := range voc {
	//	c := strings.Count(dnaU, o)
	//	count += c
	//	fmt.Printf("Колличество d%s: %d\n", o, c)
	//}

	pos := len(dnaU)
	var rdna string

	for len(rdna) < len(dnaU) {
		num := 1
		for _, r := range dnaU {
			s := string(r)
			if num == pos {
				rdna += s
				pos -= 1
				if len(rdna) == len(dnaU) { ////
					fmt.Printf("Oбратная последовательность %s\n", rdna)
				}
			}
			num += 1
		}
	}
	count := 0

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
func measuring(voc []string, dnaU string) {
	count := 0

	for _, o := range voc {
		c := strings.Count(dnaU, string(o))
		count += c
		fmt.Printf("Колличество d%s: %d\n", o, c)
	}
}
func work(s string) {

}
