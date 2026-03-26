package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
)

type Result struct {
	lines int
	words int
	bytes int
}

func countFromReader(reader io.Reader) (Result, error) {
	result := Result{}
	inWord := false

	r := bufio.NewReader(reader)

	for {
		ch, size, err := r.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return result, fmt.Errorf("ошибка чтения: %w", err)
		}

		result.bytes += size

		switch {
		case ch == '\n':
			result.lines++
			if inWord {
				result.words++
				inWord = false
			}
		case unicode.IsSpace(ch):
			if inWord {
				result.words++
				inWord = false
			}
		default:
			inWord = true
		}
	}

	if inWord {
		result.words++
	}

	return result, nil
}

func countFromFile(filename string) (Result, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Result{}, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ошибка закрытия файла: %v\n", closeErr)
		}
	}()

	return countFromReader(file)
}

func printResult(result Result, filename string) {
	_, _ = fmt.Printf("\t%d\t%d\t%d %s\n", result.lines, result.words, result.bytes, filename)
}

func main() {
	if len(os.Args) > 1 {
		result, err := countFromFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		printResult(result, os.Args[1])
	} else {
		result, err := countFromReader(os.Stdin)
		if err != nil {
			panic(err)
		}
		printResult(result, "")
	}
}
