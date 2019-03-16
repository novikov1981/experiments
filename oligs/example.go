package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"log"
	"os"
	"strings"
)

type Synthesis struct {
	//           these are tags, provide additional information for the field for some library
	Uuid      string `db:"uuid"`
	Name      string `db:"name"`
	Content   string `db:"content"`
	CreatedAt string `db:"created_at"`
}

func main() {
	var synthesisName = flag.String("name", "generic", "the name of the sysnthesis under interest")
	var filePath = flag.String("path", "", "the file path with the oligs to be parsed and saved")
	var mode = flag.String("mode", "validate", "the mode to run the application: validate, save, search")
	var oligPattern = flag.String("oligPattern", "", "the pattern of the olig's name to search for in the database")

	flag.Parse()

	fmt.Printf("synthesis '%s' from file '%s' running in mode '%s', search pattern '%s' (if search mode)\n", *synthesisName, *filePath, *mode, *oligPattern)

	oligs := readOligsFromFile(*filePath)

	printOligs(oligs)

	for i, o := range oligs {
		var oSeq = ""
		parts := strings.Split(o, ",")
		if len(parts) > 1 {
			oSeq = strings.Split(o, ",")[1]
		}
		fmt.Printf("Olig %d, sequence '%s'\n", i, oSeq)
	}
	voc := []string{"A", "C", "G", "T", "R", "Y", "K", "M", "S", "W", "B", "D", "H", "V", "N"}
	inadmissible(oligs, voc)
	measuring(voc, oligs)
}

func readOligsFromFile(filePath string) []string {
	//name:= namesinthes()
	oligs := make([]string, 0)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	utf8Reader := transform.NewReader(file, charmap.Windows1251.NewDecoder())
	scanner := bufio.NewScanner(utf8Reader)
	for scanner.Scan() {
		oligs = append(oligs, scanner.Text())
	}
	return oligs
	//sqllight(name, str)
}

func printOligs(oo []string) {
	for i, o := range oo {
		fmt.Printf("Olig %d: %s\n", i+1, o)
	}
}
func measuring(voc []string, oligs []string) { //Функция считает колличество каждого из амедитов в последовательности
	var dA, dC, dG, dT float32 = 0, 0, 0, 0
	for i, o := range oligs {
		var oSeq = ""
		parts := strings.Split(o, ",")
		if len(parts) > 1 {
			oSeq = strings.Split(o, ",")[1]
		}

		dnaU := strings.ToUpper(oSeq)
		count := 0
		if count < len(dnaU) {
			for _, o := range voc {
				c := strings.Count(dnaU, string(o))
				if c > 0 {
					cf := float32(c)
					switch o {

					case "A":
						dA += cf
					case "C":
						dC += cf
					case "G":
						dG += cf
					case "T":
						dT += cf
					case "R":
						dA += cf / 2
						dG += cf / 2
					case "Y":
						dC += cf / 2
						dT += cf / 2
					case "K":
						dG += cf / 2
						dT += cf / 2
					case "M":
						dA += cf / 2
						dC += cf / 2
					case "S":
						dG += cf / 2
						dC += cf / 2
					case "W":
						dA += cf / 2
						dT += cf / 2
					case "B":
						dC += cf / 3
						dG += cf / 3
						dT += cf / 3
					case "D":
						dA += cf / 3
						dG += cf / 3
						dT += cf / 3
					case "H":
						dA += cf / 3
						dC += cf / 3
						dG += cf / 3
					case "V":
						dA += cf / 3
						dC += cf / 3
						dG += cf / 3
					case "N":
						dA += cf / 4
						dC += cf / 4
						dG += cf / 4
						dT += cf / 4
					}
				}
				count += c
			}
			if i == 95 {
				fmt.Print("\nКолличество dA:", dA)
				fmt.Print("\nКолличество dC:", dC)
				fmt.Print("\nКолличество dG:", dG)
				fmt.Print("\nКолличество dT:", dT)
			}

		}
	}
}
func inadmissible(oligs []string, voc []string) {
	for i, o := range oligs {
		var oSeq = ""
		parts := strings.Split(o, ",")
		if len(parts) > 1 {
			oSeq = strings.Split(o, ",")[1]
		}
		dnaU:=strings.ToUpper(oSeq)
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
fmt.Printf("В строке №%d недопустимый символ №%d: %s\n", i+1,num, s)
//fmt.Printf("Символ d%s: %d\n", num, s)
}
num += 1
}

//fmt.Errorf("ОШИБКА последовательность содержит %d недопустимых символов", lenDNA-count)
//os.Exit(1)
}
}
}


//
//	fmt.Println(time.Now().Format("02-01-2006"))
//	oligs := readOligsFromFile(`F:\path\to\seq162_2018-11-01.txt`)
//	file, err := os.Open(`F:\path\to\seq162_2018-11-01.txt`)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	var base [96][3]string
//	i := 0
//	j := 0
//	for scanner.Scan() {
//		str := scanner.Text()
//			for _, o := range str {
//			stro := string(o)
//			if stro == "," {
//				j += 1
//
//			} else {
//				base[i][j] = base[i][j] + stro
//			}
//		}
//
//		fmt.Printf(" %s ", base[i][0])
//		fmt.Printf(" %s ", base[i][1])
//		fmt.Printf(" %s\n ", base[i][2])
//		j = 0
//		i += 1
//	}
//
//	if err := scanner.Err(); err != nil {
//		log.Fatal(err)
//	}
//
//	voc := []string{"A", "C", "G", "T", "R", "Y", "K", "M", "S", "W", "B", "D", "H", "V", "N"} //Слайс
//
//	var dna string
//
//	//fmt.Print("Введите последовательность ДНК: ")
//	//fmt.Fscan(os.Stdin, &dna) // чтение с консоли
//
//	dna = base[3][1]
//	dnaU := strings.ToUpper(dna) //Функция библиотеки string делает всю строку в верхнем регистре
//
//	lening(dnaU)
//
//	choice := 0
//
//	fmt.Printf("1 - подсчитать колличество амедитов в последовательности  \n2 - получить обратную последовательность \n ")
//	fmt.Fscan(os.Stdin, &choice)
//	switch choice {
//	case 1:
//		measuring(voc, base)
//	case 2:
//		revers(dnaU)
//	}
//	inadmissible(dnaU, voc)
//
//	b := "aaaa"
//
//	work(b)
//
//}

//
//func namesinthes() string {
//	var n string
//
//	fmt.Print("Введите название синтеза: ")
//	fmt.Fscan(os.Stdin, &n) // чтение с консоли
//	return n
//}
//
//func sqllight(name string, str string) {
//
//	database, err := sqlx.Open("sqlite3", "./synthesis1.db")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	statement, err := database.Prepare(`
//CREATE TABLE IF NOT EXISTS synthesis (
//	uuid TEXT PRIMARY KEY,
//	name TEXT,
//	content TEXT,
//	created_at TEXT
//);`)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	_, err = statement.Exec()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//err = insertSnt(database, Synthesis{"","synthesis cool!","96 rows","date"})
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	err = insertSnt(database, Synthesis{"", name, str, "date 555555"})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	ss := []Synthesis{}
//	err = database.Select(&ss, "SELECT uuid, name, content, created_at FROM synthesis;")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, s := range ss {
//		fmt.Printf("%+v\n", s)
//	}
//}
//
//func insertSnt(database *sqlx.DB, snt Synthesis) error {
//	statement, err := database.Prepare(`
//	INSERT INTO synthesis (uuid, name, content, created_at) VALUES (?, ?, ?, ?);
//	`)
//	if err != nil {
//		log.Fatal(err)
//	}
//	u := genuuid()
//
//	//u := uuid.NewV4()
//	if err != nil {
//		log.Fatal(err)
//	}
//	_, err = statement.Exec(fmt.Sprintf("%s", u), snt.Name, snt.Content, snt.CreatedAt)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return err
//}
//
//func genuuid() uuid.UUID {
//	u := uuid.NewV4()
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	return u
//}
//
//func lening(dnaU string) { //подсчет колличества символов для строки "dna"
//	lenDNA := len(dnaU)
//
//	fmt.Printf("Ваша последовательность %s\n", dnaU)     // %s выводимая переменная типа string \n новая строка
//	fmt.Printf("Длинна последовательности %d\n", lenDNA) // %d тип данных int
//
//}
//
//func measuring(voc []string, base [96][3]string) { //Функция считает колличество каждого из амедитов в последовательности
//	j := 1
//	var dA, dC, dG, dT float32 = 0, 0, 0, 0
//	for i := 0; i <= 95; i += 1 {
//		if base[i][j] == " " {
//			break
//		}
//		dnaU := strings.ToUpper(base[i][j])
//		count := 0
//		if count < len(dnaU) {
//			for _, o := range voc {
//				c := strings.Count(dnaU, string(o))
//				if c > 0 {
//					cf := float32(c)
//					switch o {
//
//					case "A":
//						dA += cf
//					case "C":
//						dC += cf
//					case "G":
//						dG += cf
//					case "T":
//						dT += cf
//					case "R":
//						dA += cf / 2
//						dG += cf / 2
//					case "Y":
//						dC += cf / 2
//						dT += cf / 2
//					case "K":
//						dG += cf / 2
//						dT += cf / 2
//					case "M":
//						dA += cf / 2
//						dC += cf / 2
//					case "S":
//						dG += cf / 2
//						dC += cf / 2
//					case "W":
//						dA += cf / 2
//						dT += cf / 2
//					case "B":
//						dC += cf / 3
//						dG += cf / 3
//						dT += cf / 3
//					case "D":
//						dA += cf / 3
//						dG += cf / 3
//						dT += cf / 3
//					case "H":
//						dA += cf / 3
//						dC += cf / 3
//						dG += cf / 3
//					case "V":
//						dA += cf / 3
//						dC += cf / 3
//						dG += cf / 3
//					case "N":
//						dA += cf / 4
//						dC += cf / 4
//						dG += cf / 4
//						dT += cf / 4
//					}
//				}
//				count += c
//			}
//			if i == 95 {
//				fmt.Print("\nКолличество dA:", dA)
//				fmt.Print("\nКолличество dC:", dC)
//				fmt.Print("\nКолличество dG:", dG)
//				fmt.Print("\nКолличество dT:", dT)
//			}
//
//		}
//	}
//}
//func revers(dnaU string) { //Функция делает реверсную последовательность
//	pos := len(dnaU)
//	var rdna string
//
//	for len(rdna) < len(dnaU) {
//		num := 1
//		for _, r := range dnaU {
//			s := string(r)
//			if num == pos {
//				rdna += s
//				pos -= 1
//				if len(rdna) == len(dnaU) {
//					fmt.Printf("Oбратная последовательность %s\n", rdna)
//				}
//			}
//			num += 1
//		}
//	}
//
//}
//func inadmissible(dnaU string, voc []string) {
//	count := 0
//	lenDNA := len(dnaU)
//
//	if lenDNA > count {
//
//		num := 1
//
//		for _, r := range dnaU {
//			count := 0
//			s := string(r)
//			for _, o := range voc {
//				c := strings.Count(s, o)
//				count += c
//			}
//			if count == 0 {
//				fmt.Printf("Недопустимый символ №%d: %s\n", num, s)
//				//fmt.Printf("Символ d%s: %d\n", num, s)
//			}
//			num += 1
//		}
//
//		fmt.Errorf("ОШИБКА последовательность содержит %d недопустимых символов", lenDNA-count)
//		os.Exit(1)
//	}
//
//}
//
//func work(s string) {
//
//}
