package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// SNode - узел стека
type SNode struct {
	data pair
	next *SNode
}

// Stack - структура стека
type Stack struct {
	head *SNode
}

// pair - замена std::pair<int, string>
type pair struct {
	first  int
	second string
}

// SPUSH - добавление элемента в стек
func SPUSH(list *Stack, data pair) {
	newNode := &SNode{data: data, next: list.head}
	list.head = newNode
}

// SPOP - удаление верхнего элемента стека
func SPOP(list *Stack) {
	if list.head == nil {
		return
	}
	nextNode := list.head.next
	list.head = nextNode
	// В Go сборщик мусора автоматически удалит память
}

// String - реализация вывода для pair (аналог operator<<)
func (p pair) String() string {
	return fmt.Sprintf("%d %s", p.first, p.second)
}

// SPRINT - вывод содержимого стека
func SPRINT(list *Stack) {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// ReverseStack - переворачивание стека
func ReverseStack(list *Stack) {
	reversed := &Stack{}
	current := list.head
	for current != nil {
		SPUSH(reversed, current.data)
		current = current.next
	}
	list.head = reversed.head
}

// AsteroidCycle - обработка столкновений астероидов
func AsteroidCycle(aster *Stack) {
	changed := false
	for {
		changed = false
		temp := &Stack{}

		current := aster.head
		for current != nil {
			if temp.head != nil &&
				temp.head.data.second == "right" &&
				current.data.second == "left" {

				// Обработка столкновения
				if temp.head.data.first == current.data.first {
					SPOP(temp)
				} else if temp.head.data.first > current.data.first {
					temp.head.data.first -= current.data.first
				} else {
					current.data.first -= temp.head.data.first
					SPOP(temp)
					SPUSH(temp, current.data)
				}
				changed = true
			} else {
				SPUSH(temp, current.data)
			}
			current = current.next
		}

		// Восстанавливаем исходный порядок
		aster.head = nil
		current = temp.head
		for current != nil {
			SPUSH(aster, current.data)
			current = current.next
		}

		if !changed {
			break
		}
	}
}

// asteroid - основная функция для работы с астероидами
func asteroid() {
	fmt.Println("Введите начальные данные в формате:")
	fmt.Println("<масса> <направление (l/r)> stop")

	stack := &Stack{}
	scanner := bufio.NewScanner(os.Stdin)
	
	if scanner.Scan() {
		input := scanner.Text()
		tokens := strings.Fields(input)
		
		mass := 0
		expectMass := true
		
		for _, token := range tokens {
			if token == "stop" {
				break
			}
			
			if expectMass {
				var err error
				mass, err = strconv.Atoi(token)
				if err != nil {
					fmt.Printf("Ошибка: масса должна быть числом, получено: %s\n", token)
					return
				}
				expectMass = false
			} else {
				var direction string
				if token == "l" || token == "r" {
					if token == "l" {
						direction = "left"
					} else {
						direction = "right"
					}
					ast := pair{mass, direction}
					SPUSH(stack, ast)
					expectMass = true
				} else {
					fmt.Printf("Ошибка: направление должно быть 'l' или 'r', получено: %s\n", token)
					return
				}
			}
		}
	}

	// Обрабатываем астероиды
	ReverseStack(stack)
	AsteroidCycle(stack)
	fmt.Println("Результат:")
	SPRINT(stack)
}
