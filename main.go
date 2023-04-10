package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/sametdemirezen/funtemps/conv"
	"github.com/sametdemirezen/minyr/yr"
)

var newFile *os.File

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(" >>")
	for scanner.Scan() {
		fmt.Println("Skriv convert eller average for å forsette, q for avsulette på programmet!")
		fmt.Print(">> ")
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
			fmt.Println("====================================================================")
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")
			fmt.Println("=======================================================================")
			if isFileExists("kjevik-temp- fahr-20220318-20230318.csv") {
				fmt.Println("Vil du generere filen på nytt!")
				fmt.Print(">> ")
				scanner.Scan()
				input = scanner.Text()
				if input == "j" {
					os.Remove("kjevik-temp- fahr-20220318-20230318.csv")
					fmt.Println("Filen er opprettet på nytt!")
				} else if input == "n" {
					os.Exit(0)
				}
			}
			newFile, _ = os.Create("kjevik-temp- fahr-20220318-20230318.csv")
			source, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
			if err != nil {
				log.Fatal(err)
			}
			defer source.Close()
			defer newFile.Close()

			lineScanner := bufio.NewScanner(source)
			for lineScanner.Scan() {
				lines := lineScanner.Text()
				line, _ := yr.NewLines(lines)
				newFile.WriteString(line)
				fmt.Fprintln(newFile, "")
			}
		} else if input == "average" {
			fmt.Println("Vennligst skriv for å se gjennomsnittstemperatur f eller c!")
			scanner.Scan()
			input = scanner.Text()
			if input == "f" {
				celsius := yr.AverageTempratureCelsius()
				fahr := conv.CelsiusToFahrenheit(celsius)
				fmt.Printf("Gjennomsnittstemperatur for hele perioden er "+"%.2f"+" fahrenheit.", fahr)
				fmt.Println("")
			} else if input == "c" {
				fmt.Printf("Gjennomsnittstemperatur for hele perioden er "+"%.2f"+" celsius grader.", yr.AverageTempratureCelsius())
				fmt.Println("")
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
