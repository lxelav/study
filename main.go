package main

import (
	"big_boss/classes"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func searchCommand(args []string, pools *classes.AllPools) error {
	nameCommand := args[0]

	if nameCommand == "ADD_POOL" {
		pools.AddPool(args[1])
	} else if nameCommand == "REMOVE_POOL" {
		pools.PopPool(args[1])
	} else if nameCommand == "ADD_SCHEMA" {
		pool, err := pools.GetPool(args[1])
		if err != nil {
			return err
		}

		pool.AddSchema(args[2])
	} else if nameCommand == "REMOVE_SCHEMA" {
		pool, err := pools.GetPool(args[1])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		pool.PopSchema(args[2])
	} else if nameCommand == "ADD_COLLECTION" {
		pool, err := pools.GetPool(args[1])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		schema, err := pool.GetSchema(args[2])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		if err = schema.AddContainer(args[3], args[4]); err != nil {
			//fmt.Println("Error: ", err)
			return err
		}
	} else if nameCommand == "REMOVE_COLLECTION" {
		pool, err := pools.GetPool(args[1])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		schema, err := pool.GetSchema(args[2])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		schema.PopContainer(args[3])
	} else if nameCommand == "ADD_RECORD" {
		pool, err := pools.GetPool(args[1])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		schema, err := pool.GetSchema(args[2])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		container, err := schema.GetContainer(args[3])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		if err := container.Insert(args[4], args[5]); err != nil {
			//fmt.Println("Error: ", err)
			return err
		}
	} else if nameCommand == "READ_RECORD" {
		pool, err := pools.GetPool(args[1])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		schema, err := pool.GetSchema(args[2])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		container, err := schema.GetContainer(args[3])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		result, err := container.Get(args[4])
		if err != nil {
			//fmt.Println("Error: ", err)
			return err
		}

		fmt.Printf("key: %v, value: %v", args[4], result)
	} else if nameCommand == "READ_RECORDS_RANGE" {
		pool, err := pools.GetPool(args[1])
		if err != nil {
			return err
		}

		schema, err := pool.GetSchema(args[2])
		if err != nil {
			return err
		}

		container, err := schema.GetContainer(args[3])
		if err != nil {
			return err
		}

		result, err := container.GetRange(args[4], args[5])
		if err != nil {
			return err
		}

		fmt.Printf("key_min: %v, key_max: %v; result - %v", args[4], args[5], result)
	} else if nameCommand == "UODATE_RECORD" {
		pool, err := pools.GetPool(args[1])
		if err != nil {
			return err
		}

		schema, err := pool.GetSchema(args[2])
		if err != nil {
			return err
		}

		container, err := schema.GetContainer(args[3])
		if err != nil {
			return err
		}

		if err := container.Update(args[4], args[5]); err != nil {
			return err
		}
	} else if nameCommand == "DELETE_RECOOROD" {
		pool, err := pools.GetPool(args[1])
		if err != nil {
			return err
		}

		schema, err := pool.GetSchema(args[2])
		if err != nil {
			return err
		}

		container, err := schema.GetContainer(args[3])
		if err != nil {
			return err
		}

		if err := container.Remove(args[4]); err != nil {
			fmt.Println("Error: ", err)
		}
	} else {
		fmt.Println("Ввведена неправильная команда")
	}

	return nil
}

func main() {
	pools := classes.InitAllPools()

	fmt.Print("Добро пожаловать в программу! Введите help для ознакомления с командами - ")
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}

	answer = strings.TrimRight(answer, "\n\r")
	if answer == "help" {
		file, err := os.OpenFile("help.txt", os.O_RDONLY, 0666)
		if err != nil {
			fmt.Printf("Error open file: %v", err)
			return
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := file.Close(); err != nil {
			fmt.Println("Error: ", err)
			return
		}

		for {
			fmt.Print("Введите команду или file, если хотите ввести файл с командами(q чтобы закончить): ")

			str, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка ввода: ", err)
			}

			str = strings.TrimRight(str, "\n\r")
			if str == "q" {
				break
			} else if str == "file" {
				fmt.Print("Ввведите название файла в формате <имя файла>.txt - ")

				var fileName string
				if _, err := fmt.Scan(&fileName); err != nil {
					fmt.Println("Error: ", err)
				}

				file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
				if err != nil {
					fmt.Printf("Error open file: %v", err)
					return
				}

				var line string
				scanner_file := bufio.NewScanner(file)
				for scanner_file.Scan() {
					line = scanner_file.Text()
					if err := searchCommand(strings.Split(line, " "), pools); err != nil {
						fmt.Println("Ошибка в поиске комманды - ", err)
					}
				}

				if err := file.Close(); err != nil {
					fmt.Println("Error: ", err)
					return
				}
			} else {
				if err := searchCommand(strings.Split(str, " "), pools); err != nil {
					fmt.Println("Ошибка в поиске комманды - ", err)
				}
			}
		}
	} else {
		fmt.Println("Вы ввели фигню")
	}

	pools.AddPool("tochka proverki")
}
