package main

import (
	"bufio"
	"fmt"
	"os"
)

/*Алгоритм:
1. Определяем специальные символы, которые будут являться началом и концом комм.
	# - для однострочного
	{} - для многострочного
То есть, если встречается символ комментария, то поледующие символы до конца комментария мы не запоминаем
	Для однострочного комментария концом является символ переноса строки - \n
	Для многострочного специальный символ, в нашем случае - }
Также нужно учесть вложенность комментариев - для этого будем вести два счетчика:
	SingleCounter = 0 - для однострочного
	MultipleCounter = 0 - для многострочного

2. Бегаем посимвольно по файлу и проверяем наши условия заданные выше
*/

func main() {
	SingleCounter := 0
	MultipleCounter := 0

	file, err := os.Open("parse_comment/comment.txt")
	if err != nil {
		fmt.Println("Error open file")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}

		if char == '{' {
			MultipleCounter++
		} else if char == '}' {
			MultipleCounter--
		} else if char == '#' {
			SingleCounter++
		} else if char == '\n' && SingleCounter != 0 {
			SingleCounter--
		} else if MultipleCounter == 0 && SingleCounter == 0 {
			fmt.Print(string(char))
		}
	}

}
