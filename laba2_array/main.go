package main

import (
	"bufio"
	"fmt"
	"os"
	. "project_dict/laba2_array/array"
	"strconv"
	"unicode"
)

func main() {
	var containerArrays map[string]Array = make(map[string]Array)

	var fileName string
	fmt.Print("Введите название файла: ")
	if _, err := fmt.Scan(&fileName); err != nil {
		fmt.Printf("Error scan Stdin: %v", err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error open file: %v", err)
		return
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		var commandSplit []string
		var temp string

		for _, char := range line {
			if char == ' ' || char == ',' || char == ';' || char == '(' || char == ')' {
				if temp != "" {
					commandSplit = append(commandSplit, temp)
					temp = ""
				}
			} else {
				temp += string(unicode.ToLower(char))
			}
		}

		searchCommand(commandSplit, containerArrays)
	}

}

func searchCommand(command []string, arrays map[string]Array) {
	nameCommand := command[0]

	if nameCommand == "load" {
		if err := LoadArray(command[1], command[2], arrays); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else if nameCommand == "save" {
		if err := SaveArray(command[1], command[2], arrays); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else if nameCommand == "rand" {
		count, _ := strconv.Atoi(command[2])
		lb, _ := strconv.Atoi(command[3])
		rb, _ := strconv.Atoi(command[4])

		if err := RandArray(command[1], count, lb, rb, arrays); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else if nameCommand == "concat" {
		if err := ConcatArray(command[1], command[2], arrays); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else if nameCommand == "free" {
		if err := FreeArray(command[1], arrays); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else if nameCommand == "remove" {
		index, _ := strconv.Atoi(command[2])
		count, _ := strconv.Atoi(command[3])

		if err := RemoveArray(command[1], index, count, arrays); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else if nameCommand == "copy" {
		lb, _ := strconv.Atoi(command[2])
		rb, _ := strconv.Atoi(command[3])

		if err := CopyArray(command[1], command[4], lb, rb, arrays); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else if nameCommand == "sort" {
		if err := SortArray(command[1], arrays); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else if nameCommand == "shuffle" {
		if err := ShuffleArray(command[1], arrays); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else if nameCommand == "stats" {
		if err := StatsArray(command[1], arrays); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else if nameCommand == "print" {
		if len(command) == 3 {
			if err := PrintArray(command[1], command[2], arrays); err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			lb, _ := strconv.Atoi(command[2])
			rb, _ := strconv.Atoi(command[3])

			if err := PrintRangeArray(command[1], lb, rb, arrays); err != nil {
				fmt.Printf("Error: %v", err)
			}
		}
	} else {
		fmt.Println("Invalid command!!!")
	}
}
