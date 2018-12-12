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

	fmt.Printf("Ваша последовательность %s\n", dnaU) // %s об\n
	fmt.Printf("Длинна последовательности %d\n", lenDNA)

	measuring(voc, dnaU)
	revers(dnaU)
	inadmissible(dnaU, voc)

	b := "aaaa"

	work(b)

}
func measuring(voc []string, dnaU string) { //Функция считает колличество каждого из амедитов в последовательности
	count := 0

	for _, o := range voc {
		c := strings.Count(dnaU, string(o))
		count += c
		fmt.Printf("Колличество d%s: %d\n", o, c)
	}
}
func revers(dnaU string) { //Функция делает реверсную последовательность
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
					fmt.Printf("Oбратная последовательность %s\n", rdna)
				}
			}
			num += 1
		}
	}

}
func inadmissible(dnaU string, voc []string) {
	count := 0
	lenDNA := len(dnaU)

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

}

func work(s string) {

}
