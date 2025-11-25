package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Chain struct {
	key  string
	data string
	next *Chain
}

func NewChain(key, data string) *Chain {
	return &Chain{
		key:  key,
		data: data,
		next: nil,
	}
}

func genHash(size int, key string) int {
	result := 1245
	for i := 0; i < len(key); i++ {
		result += i * int(key[i]) % size
	}
	return result % size
}

type HashTableChain struct {
	table []*Chain
	size  int
}

func NewHashTableChain(sz int) *HashTableChain {
	return &HashTableChain{
		table: make([]*Chain, sz),
		size:  sz,
	}
}

func (h *HashTableChain) Insert(key, value string) {
	hash := genHash(h.size, key)

	if h.table[hash] == nil {
		h.table[hash] = NewChain(key, value)
	} else {
		address := h.table[hash]
		for address.next != nil && address.key != key {
			address = address.next
		}
		if address.key == key {
			address.data = value
		} else {
			address.next = NewChain(key, value)
		}
	}
}

func (h *HashTableChain) Remove(key string) {
	hash := genHash(h.size, key)
	address := h.table[hash]

	if address == nil {
		fmt.Println("Элемент отсутствует в таблице")
		return
	}

	if address.key == key {
		h.table[hash] = address.next
		return
	}

	for address.next != nil && address.next.key != key {
		address = address.next
	}

	if address.next == nil {
		fmt.Println("Элемент отсутствует в таблице")
		return
	}

	address.next = address.next.next
}

func (h *HashTableChain) Find(key string) {
	hash := genHash(h.size, key)
	address := h.table[hash]
	if address == nil {
		fmt.Println("Элемент отсутствует в таблице")
		return
	}
	for address != nil {
		if address.key == key {
			fmt.Printf("Элемент найден, data = %s\n", address.data)
			return
		}
		address = address.next
	}
	fmt.Println("Элемент отсутствует в таблице")
}

func hashTableChain() {
	fmt.Println("Режим хеш-таблицы с цепочками (CHAINHASH)")
	fmt.Println("Введите размер хеш-таблицы:")

	var tableSize int
	fmt.Scan(&tableSize)

	if tableSize <= 0 {
		fmt.Println("Ошибка: размер таблицы должен быть положительным числом")
		return
	}

	hashTable := NewHashTableChain(tableSize)
	fmt.Printf("Хеш-таблица создана с размером %d\n", tableSize)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

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
				value := parts[2]
				hashTable.Insert(key, value)
				fmt.Printf("Элемент добавлен: ключ='%s', значение='%s'\n", key, value)
			} else {
				fmt.Println("Ошибка: неверный формат команды. Используйте: INSERT <ключ> <значение>")
			}
		} else if command == "REMOVE" {
			if len(parts) >= 2 {
				key := parts[1]
				fmt.Printf("Попытка удаления элемента с ключом='%s'\n", key)
				hashTable.Remove(key)
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
