package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// genHashFree - генерация хеша (линейное пробирование)
func genHashFree(size int, key string) int {
	result := 1245
	for i := 0; i < len(key); i++ {
		result += i * int(key[i]) % size
	}
	return result % size
}

// HashTableFree - хеш-таблица с линейным пробированием
type HashTableFree struct {
	table []struct {
		first  string
		second string
	}
	size int
}

// NewHashTableFree - конструктор хеш-таблицы
func NewHashTableFree(sz int) *HashTableFree {
	table := make([]struct {
		first  string
		second string
	}, sz)

	// Инициализируем все ячейки как пустые
	for i := 0; i < sz; i++ {
		table[i].first = "_empty_"
	}

	return &HashTableFree{
		table: table,
		size:  sz,
	}
}

// Insert - вставка элемента
func (h *HashTableFree) Insert(key interface{}, value string) {
	strKey := keyToString(key)
	hash := genHashFree(h.size, strKey)

	if h.table[hash].first == "_empty_" {
		h.table[hash].first = strKey
		h.table[hash].second = value
	} else {
		i := 1
		for {
			index := (hash + i) % h.size
			if h.table[index].first == "_empty_" {
				h.table[index].first = strKey
				h.table[index].second = value
				return
			}
			i++
			if i > h.size*2 {
				fmt.Println("Свободное место отсутствует в таблице")
				return
			}
		}
	}
}

// Remove - удаление элемента
func (h *HashTableFree) Remove(key interface{}) {
	strKey := keyToString(key)
	hash := genHashFree(h.size, strKey)

	if h.table[hash].first == strKey {
		h.table[hash].first = "_empty_"
	} else {
		i := 1
		for {
			index := (hash + i) % h.size
			if h.table[index].first == strKey {
				h.table[index].first = "_empty_"
				return
			}
			i++
			if i > h.size*2 {
				fmt.Println("Элемент отсутствует в таблице")
				return
			}
		}
	}
}

// Find - поиск элемента
func (h *HashTableFree) Find(key interface{}) {
	strKey := keyToString(key)
	hash := genHashFree(h.size, strKey)

	if h.table[hash].first == strKey {
		fmt.Printf("Data = %s\n", h.table[hash].second)
	} else {
		i := 1
		for {
			index := (hash + i) % h.size
			if h.table[index].first == strKey {
				fmt.Printf("Data = %s\n", h.table[index].second)
				return
			}
			i++
			if i > h.size*2 {
				fmt.Println("Элемент отсутствует в таблице")
				return
			}
		}
	}
}

func hashTableFree() {
	fmt.Println("Введите размер хеш-таблицы:")

	var tableSize int
	_, err := fmt.Scan(&tableSize)
	if err != nil {
		fmt.Println("Ошибка: неверный формат размера таблицы")
		return
	}

	if tableSize <= 0 {
		fmt.Println("Ошибка: размер таблицы должен быть положительным числом")
		return
	}

	hashTable := NewHashTableFree(tableSize)
	fmt.Printf("Хеш-таблица создана с размером %d\n", tableSize)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Очищаем буфер после ввода числа

	for {
		fmt.Println("\nДоступные команды:")
		fmt.Println("INSERT <ключ> <значение> - добавить элемент")
		fmt.Println("REMOVE <ключ> - удалить элемент")
		fmt.Println("FIND <ключ> - найти элемент")
		fmt.Println("EXIT - выход из программы")
		fmt.Println("Введите команду:")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		if command == "EXIT" {
			fmt.Println("Выход из программы")
			break
		} else if command == "INSERT" {
			if len(parts) >= 3 {
				key := parts[1]
				value := strings.Join(parts[2:], " ")
				hashTable.Insert(key, value)
				fmt.Printf("Элемент добавлен: ключ='%s', значение='%s'\n", key, value)
			} else {
				fmt.Println("Ошибка: неверный формат команды. Используйте: INSERT <ключ> <значение>")
			}
		} else if command == "REMOVE" {
			if len(parts) >= 2 {
				key := parts[1]
				hashTable.Remove(key)
				fmt.Printf("Попытка удаления элемента с ключом='%s'\n", key)
			} else {
				fmt.Println("Ошибка: неверный формат команды. Используйте: REMOVE <ключ>")
			}
		} else if command == "FIND" {
			if len(parts) >= 2 {
				key := parts[1]
				hashTable.Find(key)
			} else {
				fmt.Println("Ошибка: неверный формат команды. Используйте: FIND <ключ>")
			}
		} else {
			fmt.Printf("Ошибка: неизвестная команда '%s'\n", command)
		}
	}
}
