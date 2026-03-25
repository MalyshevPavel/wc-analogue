package main

import (
	"bufio"
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

func countFromStdin(reader io.Reader) (Result, error) {
	result := Result{}
	inWord := false

	r := bufio.NewReader(reader)

	for {
		ch, size, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
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

	return countFromStdin(file)
}

func exitOnError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
	os.Exit(1)
}

func printResult(result Result, filename string) {
	if filename != "" {
		_, _ = fmt.Printf("\t%d\t%d\t%d %s\n", result.lines, result.words, result.bytes, filename)
	} else {
		_, _ = fmt.Printf("\t%d\t%d\t%d\n", result.lines, result.words, result.bytes)
	}
}

func main() {
	if len(os.Args) > 1 {
		result, err := countFromFile(os.Args[1])
		if err != nil {
			exitOnError(err)
		}
		printResult(result, os.Args[1])
	} else {
		result, err := countFromStdin(os.Stdin)
		if err != nil {
			exitOnError(err)
		}
		printResult(result, "")
	}
}
