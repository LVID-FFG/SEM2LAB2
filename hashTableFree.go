package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func genHashFree(size int, key string) int {
	result := 1245
	for i := 0; i < len(key); i++ {
		result += i * int(key[i]) % size
	}
	return result % size
}

type HashTableFree struct {
	table []pairF
	size  int
}

type pairF struct {
	first  string
	second string
}

func NewHashTableFree(sz int) *HashTableFree {
	table := make([]pairF, sz)
	for i := 0; i < sz; i++ {
		table[i].first = "_empty_"
	}
	return &HashTableFree{
		table: table,
		size:  sz,
	}
}

func (h *HashTableFree) insert(key, value string) {
	Hash := genHashFree(h.size, key)
	if h.table[Hash].first == key {
		h.table[Hash] = pairF{key, value}
		fmt.Printf("Элемент добавлен: ключ='%s', значение='%s'\n", key, value)
		return
	}
	if h.table[Hash].first == "_empty_" {
		h.table[Hash] = pairF{key, value}
		fmt.Printf("Элемент добавлен: ключ='%s', значение='%s'\n", key, value)
	} else {
		i := 1
		for {
			index := (Hash + i) % h.size
			if h.table[index].first == key {
				h.table[Hash] = pairF{key, value}
				fmt.Printf("Элемент добавлен: ключ='%s', значение='%s'\n", key, value)
				return
			}
			if h.table[index].first == "_empty_" {
				h.table[index] = pairF{key, value}
				fmt.Printf("Элемент добавлен: ключ='%s', значение='%s'\n", key, value)
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

func (h *HashTableFree) remove(key string) {
	Hash := genHashFree(h.size, key)
	if h.table[Hash].first == key {
		h.table[Hash].first = "_empty_"
		fmt.Println("Элемент успешно удалён")
	} else {
		i := 1
		for {
			index := (Hash + i) % h.size
			if h.table[index].first == key {
				h.table[index].first = "_empty_"
				fmt.Println("Элемент успешно удалён")
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

func (h *HashTableFree) find(key string) {
	Hash := genHashFree(h.size, key)
	if h.table[Hash].first == key {
		fmt.Printf("Data = %s\n", h.table[Hash].second)
	} else {
		i := 1
		for {
			index := (Hash + i) % h.size
			if h.table[index].first == key {
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
	fmt.Scan(&tableSize)

	if tableSize <= 0 {
		fmt.Println("Ошибка: размер таблицы должен быть положительным числом")
		return
	}

	hashTable := NewHashTableFree(tableSize)
	fmt.Printf("Хеш-таблица создана с размером %d\n", tableSize)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nДоступные команды:")
		fmt.Println("INSERT <ключ> <значение> - добавить элемент")
		fmt.Println("REMOVE <ключ> - удалить элемент")
		fmt.Println("FIND <ключ> - найти элемент")
		fmt.Println("EXIT - выход из программы")
		fmt.Println("Введите команду:")

		scanner.Scan()
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
				hashTable.insert(key, value)
			} else {
				fmt.Println("Ошибка: неверный формат команды. Используйте: INSERT <ключ> <значение>")
			}
		} else if command == "REMOVE" {
			if len(parts) >= 2 {
				key := parts[1]
				fmt.Printf("Попытка удаления элемента с ключом='%s'...\n", key)
				hashTable.remove(key)
			} else {
				fmt.Println("Ошибка: неверный формат команды. Используйте: REMOVE <ключ>")
			}
		} else if command == "FIND" {
			if len(parts) >= 2 {
				key := parts[1]
				hashTable.find(key)
			} else {
				fmt.Println("Ошибка: неверный формат команды. Используйте: FIND <ключ>")
			}
		} else {
			fmt.Printf("Ошибка: неизвестная команда '%s'\n", command)
		}
	}
}
