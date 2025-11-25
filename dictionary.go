package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Dictionary struct {
	size int
	cap  int
	data []string
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		size: 0,
		cap:  10,
		data: make([]string, 10),
	}
}

func (d *Dictionary) AddWord(str string) {
	if d.size == d.cap {
		newCap := d.cap * 2
		newData := make([]string, newCap)
		copy(newData, d.data)
		d.data = newData
		d.cap = newCap
	}
	d.data[d.size] = str
	d.size++
}

func (d *Dictionary) Errors(str string) int {
	result := 0
	words := strings.Fields(str)

	for _, word := range words {
		wordLower := strings.ToLower(word)

		// Если слово полностью в нижнем регистре
		if word == wordLower {
			result++
			continue
		}

		// Проверка на несколько заглавных букв
		manyUpper := false
		upperCount := 0
		for _, ch := range word {
			if unicode.IsUpper(ch) {
				upperCount++
				if upperCount > 1 {
					manyUpper = true
					break
				}
			}
		}
		if manyUpper {
			result++
			continue
		}

		// Проверяем есть ли слово в словаре (без учета регистра)
		haveWord := false
		for i := 0; i < d.size; i++ {
			dictWordLower := strings.ToLower(d.data[i])
			if wordLower == dictWordLower {
				haveWord = true
				break
			}
		}

		// Проверяем точное совпадение
		noError := false
		for i := 0; i < d.size; i++ {
			if word == d.data[i] {
				noError = true
				break
			}
		}

		if haveWord && !noError {
			result++
		}
	}
	return result
}

func (d *Dictionary) String() string {
	return strings.Join(d.data[:d.size], " ") + "\n"
}

func dictionary() {
	fmt.Println("Введите словарь (слова через пробел), затем слово 'stop', затем строку для проверки:")

	dict := NewDictionary()
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	input := scanner.Text()

	words := strings.Fields(input)
	foundStop := false
	var textToCheck string

	for _, word := range words {
		if word == "stop" {
			foundStop = true
			break
		}
		dict.AddWord(word)
	}

	if foundStop {
		// Извлекаем оставшийся текст после 'stop'
		stopIndex := strings.Index(input, "stop")
		if stopIndex != -1 {
			textToCheck = strings.TrimSpace(input[stopIndex+4:])
		}
	}

	// Если текст для проверки пустой, читаем новую строку
	if textToCheck == "" {
		fmt.Println("Введите строку для проверки:")
		if !scanner.Scan() {
			return
		}
		textToCheck = scanner.Text()
	}

	// Проверяем ошибки
	errorCount := dict.Errors(textToCheck)

	fmt.Printf("Словарь: %s", dict.String())
	fmt.Printf("Проверяемая строка: %s\n", textToCheck)
	fmt.Printf("Количество ошибок: %d\n", errorCount)
}
