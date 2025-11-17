package main

import (
	"fmt"
)

// MPair - аналог pair<char, char>
type MPair struct {
	first  byte
	second byte
}

// MorfTable - структура для проверки изоморфизма
type MorfTable struct {
	table []MPair
	size  int
	cap   int
}

// NewMorfTable - конструктор таблицы
func NewMorfTable() *MorfTable {
	return &MorfTable{
		table: make([]MPair, 10),
		size:  0,
		cap:   10,
	}
}

// SearchSimbolMorf - поиск символа в морфологической таблице
func (m *MorfTable) SearchSimbolMorf(str1 string, str2 string, symbol byte) bool {
	// Сначала проверяем, есть ли символ уже в таблице
	for i := 0; i < m.size; i++ {
		if m.table[i].first == symbol {
			// Если символ есть в таблице, проверяем что отображение корректно
			expectedChar := m.table[i].second
			for j := 0; j < len(str1); j++ {
				if str1[j] == symbol && str2[j] != expectedChar {
					return false
				}
			}
			return true
		}
	}

	// Если символа нет в таблице, находим соответствующий символ из str2
	mappedChar := byte(0)
	for i := 0; i < len(str1); i++ {
		if str1[i] == symbol {
			if mappedChar == 0 {
				mappedChar = str2[i]
			} else if mappedChar != str2[i] {
				return false // Один символ отображается в разные - не изоморфно
			}
		}
	}

	// Проверяем, что mappedChar не используется для другого символа
	for i := 0; i < m.size; i++ {
		if m.table[i].second == mappedChar {
			return false // Два разных символа отображаются в один - не изоморфно
		}
	}

	// Добавляем новое отображение в таблицу
	if m.size == len(m.table) {
		// Увеличиваем емкость в 2 раза
		m.cap *= 2
		newTable := make([]MPair, m.cap)
		copy(newTable, m.table)
		m.table = newTable
	}
	m.table[m.size] = MPair{symbol, mappedChar}
	m.size++
	return true
}

// IsMorf - проверка изоморфизма строк
func IsMorf(str1 string, str2 string) {
	if len(str1) != len(str2) {
		fmt.Println("FALSE")
		return
	}
	table := NewMorfTable()
	for i := 0; i < len(str1); i++ {
		if !table.SearchSimbolMorf(str1, str2, str1[i]) {
			fmt.Println("FALSE")
			return
		}
	}
	fmt.Println("TRUE")
}

func morf() {
	var str1, str2 string
	fmt.Println("Введите строки")
	fmt.Scan(&str1)
	fmt.Scan(&str2)
	IsMorf(str1, str2)
}