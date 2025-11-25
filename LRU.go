package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	key  string
	data string
	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
}

func NewLRUCache(cap int) *LRUCache {
	head := &Node{key: "", data: ""}
	tail := &Node{key: "", data: ""}
	head.next = tail
	tail.prev = head
	return &LRUCache{
		capacity: cap,
		head:     head,
		tail:     tail,
	}
}

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

func (lru *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) addToFront(node *Node) {
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) moveToFront(node *Node) {
	lru.removeNode(node)
	lru.addToFront(node)
}

func (lru *LRUCache) removeLRU() {
	lruNode := lru.tail.prev
	lru.removeNode(lruNode)
}

func (lru *LRUCache) Get(key string) *Node {
	node := lru.findNode(key)
	if node != nil {
		lru.moveToFront(node)
		return node
	}
	return nil
}

func (lru *LRUCache) Set(key, data string) {
	node := lru.findNode(key)
	if node != nil {
		lru.moveToFront(node)
		node.data = data
	} else {
		if lru.GetSize() >= lru.capacity {
			lru.removeLRU()
		}
		newNode := &Node{key: key, data: data}
		lru.addToFront(newNode)
	}
}

func (lru *LRUCache) GetSize() int {
	size := 0
	current := lru.head.next
	for current != lru.tail {
		size++
		current = current.next
	}
	return size
}

func (lru *LRUCache) PrintCache() {
	fmt.Print("Кэш: ")
	current := lru.head.next
	for current != lru.tail {
		fmt.Printf("[\"%s %s\"] ", current.key, current.data)
		current = current.next
	}
	fmt.Println()
}

func LRU() {
	fmt.Println("Режим LRU-кеша")
	fmt.Println("Введите размер кеша:")

	var cacheSize int
	fmt.Scan(&cacheSize)

	if cacheSize <= 0 {
		fmt.Println("Ошибка: размер кеша должен быть положительным числом")
		return
	}

	cache := NewLRUCache(cacheSize)
	fmt.Printf("LRU-кеш создан с размером %d\n", cacheSize)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nДоступные команды:")
		fmt.Println("SET <ключ> - добавить элемент в кеш")
		fmt.Println("GET <ключ> - проверить наличие элемента в кеше")
		fmt.Println("PRINT - вывести текущее состояние кеша")
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
		} else if command == "SET" {
			if len(parts) >= 3 {
				key := parts[1]
				data := strings.Join(parts[2:], " ")
				cache.Set(key, data)
				fmt.Printf("Элемент добавлен в кеш: ключ='%s'\n", key)
				cache.PrintCache()
			} else {
				fmt.Println("Ошибка: неверный формат команды. Используйте: SET <ключ>")
			}
		} else if command == "GET" {
			if len(parts) >= 2 {
				key := parts[1]
				found := cache.Get(key)
				if found != nil {
					fmt.Printf("Элемент найден в кеше: ='%s'\n", found.data)
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
