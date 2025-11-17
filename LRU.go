package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Node - узел двусвязного списка
type Node struct {
	key  string
	prev *Node
	next *Node
}

// NewNode - конструктор Node
func NewNode(k string) *Node {
	return &Node{
		key:  k,
		prev: nil,
		next: nil,
	}
}

// LRUCache - структура LRU кэша
type LRUCache struct {
	capacity int
	head     *Node // MRU (Most Recently Used)
	tail     *Node // LRU (Least Recently Used)
}

// NewLRUCache - конструктор LRUCache
func NewLRUCache(cap int) *LRUCache {
	// Создаем фиктивные голову и хвост для упрощения операций
	head := NewNode("")
	tail := NewNode("")
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: cap,
		head:     head,
		tail:     tail,
	}
}

// findNode - поиск узла по ключу
func (lru *LRUCache) findNode(key string) *Node {
	current := lru.head.next
	for current != lru.tail {
		if current.key == key {
			return current
		}
		current = current.next
	}
	return nil
}

// removeNode - удаление узла из списка
func (lru *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// addToFront - добавление узла в начало списка (после head)
func (lru *LRUCache) addToFront(node *Node) {
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next.prev = node
	lru.head.next = node
}

// moveToFront - перемещение узла в начало (обновление до MRU)
func (lru *LRUCache) moveToFront(node *Node) {
	lru.removeNode(node)
	lru.addToFront(node)
}

// removeLRU - удаление LRU элемента
func (lru *LRUCache) removeLRU() {
	lruNode := lru.tail.prev
	lru.removeNode(lruNode)
	// В Go сборщик мусора автоматически удалит объект
}

// Get - GET операция
func (lru *LRUCache) Get(key string) bool {
	node := lru.findNode(key)
	if node != nil {
		lru.moveToFront(node) // Обновляем как MRU
		return true
	}
	return false
}

// Set - SET операция
func (lru *LRUCache) Set(key string) {
	node := lru.findNode(key)
	if node != nil {
		// Ключ уже существует - просто перемещаем в начало
		lru.moveToFront(node)
	} else {
		// Новый ключ
		if lru.GetSize() >= lru.capacity {
			lru.removeLRU() // Удаляем LRU если достигнут лимит
		}

		// Создаем новый узел и добавляем в начало
		newNode := NewNode(key)
		lru.addToFront(newNode)
	}
}

// GetSize - вспомогательная функция для получения размера
func (lru *LRUCache) GetSize() int {
	size := 0
	current := lru.head.next
	for current != lru.tail {
		size++
		current = current.next
	}
	return size
}

// PrintCache - вспомогательная функция для отладки
func (lru *LRUCache) PrintCache() {
	fmt.Print("Кэш: ")
	current := lru.head.next
	for current != lru.tail {
		fmt.Printf("[\"%s\"] ", current.key)
		current = current.next
	}
	fmt.Println()
}

func LRU() {
	fmt.Println("Режим LRU-кеша (LRUCASH)")
	fmt.Println("Введите размер кеша:")

	var cacheSize int
	_, err := fmt.Scan(&cacheSize)
	if err != nil {
		fmt.Println("Ошибка: неверный формат размера кеша")
		return
	}

	if cacheSize <= 0 {
		fmt.Println("Ошибка: размер кеша должен быть положительным числом")
		return
	}

	cache := NewLRUCache(cacheSize)
	fmt.Printf("LRU-кеш создан с размером %d\n", cacheSize)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Очищаем буфер после ввода числа

	for {
		fmt.Println("\nДоступные команды:")
		fmt.Println("SET <ключ> - добавить элемент в кеш")
		fmt.Println("GET <ключ> - проверить наличие элемента в кеше")
		fmt.Println("PRINT - вывести текущее состояние кеша")
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
		} else if command == "SET" {
			if len(parts) >= 2 {
				key := parts[1]
				cache.Set(key)
				fmt.Printf("Элемент добавлен в кеш: ключ='%s'\n", key)
				cache.PrintCache()
			} else {
				fmt.Println("Ошибка: неверный формат команды. Используйте: SET <ключ>")
			}
		} else if command == "GET" {
			if len(parts) >= 2 {
				key := parts[1]
				found := cache.Get(key)
				if found {
					fmt.Printf("Элемент найден в кеше: ключ='%s'\n", key)
				} else {
					fmt.Printf("Элемент отсутствует в кеше: ключ='%s'\n", key)
				}
				cache.PrintCache()
			} else {
				fmt.Println("Ошибка: неверный формат команды. Используйте: GET <ключ>")
			}
		} else if command == "PRINT" {
			cache.PrintCache()
		} else {
			fmt.Printf("Ошибка: неизвестная команда '%s'\n", command)
		}
	}
}
