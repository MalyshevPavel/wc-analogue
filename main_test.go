package main

import (
	"fmt"
	"testing"
)

var errInputMissing = fmt.Errorf("ошибка: укажите либо имя файла, либо строку для анализа")

func TestWordCounterTable(t *testing.T) {

	table := []struct {
		inputFile          string
		inputString        string
		expectedLinesCount int
		expectedWordsCount int
		expectedBytesCount int
		err                error
	}{
		{"", " ", 1, 0, 1, nil},
		{"", "one two three\nfour five", 2, 5, 23, nil},
		{"", "\n\n\n", 4, 0, 3, nil},
		{"./resources/empty.txt", "", 1, 0, 0, nil},
		{"./resources/data.txt", "", 2, 5, 23, nil},
		{"./resources/onlyNewLines.txt", "", 2, 0, 2, nil},
		{"./resources/symbols.txt", "", 2, 5, 21, nil},
		{"", "", -1, -1, -1, errInputMissing},
	}

	for _, data := range table {
		actualLinesCount, actualWordsCount, actualBytesCount, err := wordCounter(data.inputFile, data.inputString)

		if data.err != nil {
			if err == nil {
				t.Errorf("Expected error %q, but got nil", data.err)
				continue
			}
			if err.Error() != data.err.Error() {
				t.Errorf("Expected error %q, but got %q", data.err, err)
			}
			continue
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				continue
			}
		}

		if actualLinesCount != data.expectedLinesCount {
			fmt.Println("Result for input:", data.inputFile, data.inputString)
			t.Errorf("Expected %d lines, got %d", data.expectedLinesCount, actualLinesCount)
		}

		if actualWordsCount != data.expectedWordsCount {
			fmt.Println("Result for input:", data.inputFile, data.inputString)
			t.Errorf("Expected %d words, got %d", data.expectedWordsCount, actualWordsCount)
		}

		if actualBytesCount != data.expectedBytesCount {
			fmt.Println("Result for input:", data.inputFile, data.inputString)
			t.Errorf("Expected %d bytes, got %d", data.expectedBytesCount, actualBytesCount)
		}
	}
}
