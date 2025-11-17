package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Chain - элемент цепочки хеш-таблицы
type Chain struct {
	key  string
	data string
	next *Chain
}

// NewChain - конструктор Chain
func NewChain(key string, data string) *Chain {
	return &Chain{
		key:  key,
		data: data,
		next: nil,
	}
}

// genHash - генерация хеша
func genHash(size int, key string) int {
	result := 1245
	for i := 0; i < len(key); i++ {
		result += i * int(key[i]) % size
	}
	return result % size
}

// HashTableChain - хеш-таблица с цепочками
type HashTableChain struct {
	table []*Chain
	size  int
}

// NewHashTableChain - конструктор хеш-таблицы
func NewHashTableChain(sz int) *HashTableChain {
	table := make([]*Chain, sz)
	for i := range table {
		table[i] = nil
	}
	return &HashTableChain{
		table: table,
		size:  sz,
	}
}

// keyToString - преобразование ключа в строку
func keyToString(key interface{}) string {
	switch k := key.(type) {
	case string:
		return k
	case int:
		return strconv.Itoa(k)
	default:
		return fmt.Sprintf("%v", k)
	}
}

// Insert - вставка элемента
func (h *HashTableChain) Insert(key interface{}, value string) {
	strKey := keyToString(key)
	hash := genHash(h.size, strKey)
	
	if h.table[hash] == nil {
		h.table[hash] = NewChain(strKey, value)
	} else {
		address := h.table[hash]
		for address.next != nil {
			address = address.next
		}
		address.next = NewChain(strKey, value)
	}
}

// Remove - удаление элемента
func (h *HashTableChain) Remove(key interface{}) {
	strKey := keyToString(key)
	hash := genHash(h.size, strKey)
	address := h.table[hash]
	
	if address == nil {
		fmt.Println("Элемент отсутствует в таблице")
		return
	}
	
	if address.key == strKey {
		h.table[hash] = address.next
		return
	}
	
	for address.next != nil && address.next.key != strKey {
		address = address.next
	}
	
	if address.next == nil {
		fmt.Println("Элемент отсутствует в таблице")
		return
	}
	
	deleteChain := address.next
	address.next = address.next.next
	_ = deleteChain // В Go сборщик мусора удалит объект автоматически
}

// Find - поиск элемента
func (h *HashTableChain) Find(key interface{}) {
	strKey := keyToString(key)
	hash := genHash(h.size, strKey)
	address := h.table[hash]
	
	if address == nil {
		fmt.Println("Элемент отсутствует в таблице")
		return
	}
	
	for address != nil {
		if address.key == strKey {
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
	_, err := fmt.Scan(&tableSize)
	if err != nil {
		fmt.Println("Ошибка: неверный формат размера таблицы")
		return
	}
	
	if tableSize <= 0 {
		fmt.Println("Ошибка: размер таблицы должен быть положительным числом")
		return
	}
	
	hashTable := NewHashTableChain(tableSize)
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