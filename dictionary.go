package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Dictionary - структура словаря
type Dictionary struct {
	size int
	cap  int
	data []string
}

// NewDictionary - конструктор словаря (аналог Dictionary())
func NewDictionary() *Dictionary {
	return &Dictionary{
		size: 0,
		cap:  10,
		data: make([]string, 10),
	}
}

// addWord - добавление слова в словарь
func (d *Dictionary) addWord(str string) {
	if d.size == len(d.data) {
		// Увеличиваем емкость в 2 раза
		newData := make([]string, d.cap*2)
		copy(newData, d.data)
		d.data = newData
		d.cap *= 2
	}
	d.data[d.size] = str
	d.size++
}

// errors - подсчет ошибок в строке
func (d *Dictionary) errors(str string) int {
	result := 0
	words := strings.Fields(str)
	
	for _, word := range words {
		wordDown := strings.ToLower(word)
		
		// Если слово уже в нижнем регистре
		if word == wordDown {
			result++
			continue
		}

		haveWord := false
		for i := 0; i < d.size; i++ {
			dataDown := strings.ToLower(d.data[i])
			if wordDown == dataDown {
				haveWord = true
				break
			}
		}

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

// String - реализация вывода для Dictionary (аналог operator<<)
func (d *Dictionary) String() string {
	var result strings.Builder
	for i := 0; i < d.size; i++ {
		result.WriteString(d.data[i])
		result.WriteString(" ")
	}
	result.WriteString("\n")
	return result.String()
}

// toLowerCase - вспомогательная функция для преобразования строки к нижнему регистру
// (в Go есть встроенная strings.ToLower, но для полноты перевода)
func toLowerCase(str string) string {
	return strings.Map(unicode.ToLower, str)
}

func dictionary() {
	fmt.Println("Введите словарь (слова через пробел), затем слово 'stop', затем строку для проверки:")
	
	dict := NewDictionary()
	scanner := bufio.NewScanner(os.Stdin)
	
	if scanner.Scan() {
		input := scanner.Text()
		words := strings.Fields(input)
		
		// Флаг для определения, когда заканчивается словарь
		readingDict := true
		var textToCheck strings.Builder
		
		for _, word := range words {
			if word == "stop" {
				readingDict = false
				continue
			}
			
			if readingDict {
				dict.addWord(word)
			} else {
				if textToCheck.Len() > 0 {
					textToCheck.WriteString(" ")
				}
				textToCheck.WriteString(word)
			}
		}
		
		// Если после обработки всех слов textToCheck пуст, читаем новую строку
		checkText := textToCheck.String()
		if checkText == "" {
			fmt.Println("Введите строку для проверки:")
			if scanner.Scan() {
				checkText = scanner.Text()
			}
		}
		
		// Проверяем ошибки
		errorCount := dict.errors(checkText)
		
		fmt.Printf("Словарь: %s", dict.String())
		fmt.Printf("Проверяемая строка: %s\n", checkText)
		fmt.Printf("Количество ошибок: %d\n", errorCount)
	}
}
