package main

import (
	"strings"
	"testing"
)

func TestCountFromFile(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		wantLines int
		wantWords int
		wantBytes int
	}{
		{
			name:      "data.txt - обычный текст",
			filename:  "./resources/data.txt",
			wantLines: 1,
			wantWords: 5,
			wantBytes: 23,
		},
		{
			name:      "onlyNewLines.txt - только переводы строк",
			filename:  "./resources/onlyNewLines.txt",
			wantLines: 2,
			wantWords: 0,
			wantBytes: 2,
		},
		{
			name:      "symbols.txt - спецсимволы",
			filename:  "./resources/symbols.txt",
			wantLines: 1,
			wantWords: 5,
			wantBytes: 21,
		},
		{
			name:      "empty.txt - пустой файл",
			filename:  "./resources/empty.txt",
			wantLines: 0,
			wantWords: 0,
			wantBytes: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := countFromFile(tt.filename)
			if err != nil {
				t.Fatalf("ошибка countFromFile: %v", err)
			}

			if result.lines != tt.wantLines {
				t.Errorf("lines = %d, want %d", result.lines, tt.wantLines)
			}
			if result.words != tt.wantWords {
				t.Errorf("words = %d, want %d", result.words, tt.wantWords)
			}
			if result.bytes != tt.wantBytes {
				t.Errorf("bytes = %d, want %d", result.bytes, tt.wantBytes)
			}
		})
	}
}

func TestCountFromStdin(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantLines int
		wantWords int
		wantBytes int
	}{
		{
			name:      "пустой ввод",
			input:     "\n",
			wantLines: 1,
			wantWords: 0,
			wantBytes: 1,
		},
		{
			name:      "одна строка",
			input:     "hello world\n",
			wantLines: 1,
			wantWords: 2,
			wantBytes: 12,
		},
		{
			name:      "несколько строк",
			input:     "one\ntwo\nthree\n",
			wantLines: 3,
			wantWords: 3,
			wantBytes: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)

			result, err := countFromStdin(reader)
			if err != nil {
				t.Fatalf("ошибка countFromStdin: %v", err)
			}

			if result.lines != tt.wantLines {
				t.Errorf("lines = %d, want %d", result.lines, tt.wantLines)
			}
			if result.words != tt.wantWords {
				t.Errorf("words = %d, want %d", result.words, tt.wantWords)
			}
			if result.bytes != tt.wantBytes {
				t.Errorf("bytes = %d, want %d", result.bytes, tt.wantBytes)
			}
		})
	}
}
