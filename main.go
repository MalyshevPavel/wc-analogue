package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func wordCounter(fileName string, inputString string) (int, int, int, error) {
	var linesCounter, wordsCounter, bytesCounter = -1, -1, -1

	if len(fileName) == 0 && len(inputString) == 0 {
		return -1, -1, -1, fmt.Errorf("ошибка: укажите либо имя файла, либо строку для анализа")
	}

	fileInfo, err := os.Stat(fileName)
	if err == nil && !fileInfo.IsDir() {
		data, err := os.ReadFile(fileName)
		if err != nil {
			return -1, -1, -1, err
		}
		bytesCounter = len(data)

		linesCounter = bytes.Count(data, []byte{'\n'})

		// файл пуст, но все равно содержит 1 строку
		if len(data) == 0 {
			linesCounter++
		}

		// Если файл не пуст и не заканчивается на '\n', добавляем 1 строку по аналогии с командой wc
		if len(data) > 0 && data[len(data)-1] != '\n' {
			linesCounter++
		}

		wordsCounter = len(strings.Fields(string(data)))

		fmt.Println("\t", linesCounter, "\t", wordsCounter, "\t", bytesCounter, fileName)
		return linesCounter, wordsCounter, bytesCounter, nil
	}

	// Преобразуем строку с \n в реальные переносы строк
	inputString = strings.ReplaceAll(inputString, `\n`, "\n")

	bytesCounter = len([]byte(inputString))

	linesCounter = strings.Count(inputString, "\n") + 1

	wordsCounter = len(strings.Fields(inputString))

	fmt.Println("\t", linesCounter, "\t", wordsCounter, "\t", bytesCounter)
	return linesCounter, wordsCounter, bytesCounter, nil
}

func main() {

}
