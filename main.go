package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/sametdemirezen/minyr/yr"
)

var newFile *os.File
var err error

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Skriv convert for å forsette eller q for avsulette!")
	fmt.Print(" >>")
	for scanner.Scan() {
		fmt.Print(">> ")
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")

			if isFileExists("kjevik-temp- fahr-20220318-20230318.csv") {
				fmt.Println("Vil du generere filen på nytt!")
				fmt.Print(">> ")
				scanner.Scan()
				input = scanner.Text()
				if input == "j" {
					os.Remove("kjevik-temp- fahr-20220318-20230318.csv")
					fmt.Println("File is created!")

				} else if input == "n" {
					os.Exit(0)
				}
			}
			newFile, _ = os.Create("kjevik-temp- fahr-20220318-20230318.csv")

			source, _ := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
			defer source.Close()
			defer newFile.Close()
			checkNilError(err)

			lineScanner := bufio.NewScanner(source)
			for lineScanner.Scan() {
				lines := lineScanner.Text()
				line, _ := yr.CelsiusToFahrenheitLine(lines)
				newFile.WriteString(line)
				fmt.Fprintln(newFile, "")

				/*line = strings.Split(lines, ";")
				if strings.HasSuffix(line[3], "tur") {

				} else {
					celsius = line[3]
					fahrenheit, _ = yr.CelsiusToFahrenheitString(celsius)
					line[3] = fahrenheit

				}*/

			}
		} else {
			fmt.Println("Venligst velg convert, average eller exit:")

		}
	}

}

func isFileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !errors.Is(err, os.ErrNotExist)
}

func checkNilError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
