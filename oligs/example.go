package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open(`F:\path\to\file.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var base [96][5]string
	i := 0
	j := 0
	for scanner.Scan() {
		base[i][j] = scanner.Text()
		fmt.Println(base[i][j])
		i += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	voc := []string{"A", "C", "G", "T"} //Слайс

	var dna string

	fmt.Print("Введите последовательность ДНК: ")
	fmt.Fscan(os.Stdin, &dna) // чтение с консоли

	dnaU := strings.ToUpper(dna) //Функция библиотеки string делает всю строку в верхнем регистре

	lening(dnaU)

	choice := 0

	fmt.Printf("1 - подсчитать колличество амедитов в последовательности  \n2 - получить обратную последовательность \n ")
	fmt.Fscan(os.Stdin, &choice)
	switch choice {
	case 1:
		measuring(voc, dnaU)
	case 2:
		revers(dnaU)
	}
	inadmissible(dnaU, voc)

	b := "aaaa"

	work(b)

}

func lening(dnaU string) { //подсчет колличества символов для строки "dna"
	lenDNA := len(dnaU)

	fmt.Printf("Ваша последовательность %s\n", dnaU)     // %s выводимая переменная типа string \n новая строка
	fmt.Printf("Длинна последовательности %d\n", lenDNA) // %d тип данных int

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
