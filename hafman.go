package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// HNode - узел дерева Хаффмана
type HNode struct {
	symbol     byte
	probability float64
	code       string
	parent     *HNode
	left       *HNode
	right      *HNode
}

// NewHNodeFromPair - конструктор из пары (вероятность, символ)
func NewHNodeFromPair(sym struct {
	first  float64
	second byte
}) *HNode {
	return &HNode{
		symbol:      sym.second,
		probability: sym.first,
		code:        "",
		parent:      nil,
		left:        nil,
		right:       nil,
	}
}

// NewHNodeFromChildren - конструктор из двух дочерних узлов
func NewHNodeFromChildren(l, r *HNode) *HNode {
	return &HNode{
		symbol:      '\x00',
		probability: l.probability + r.probability,
		code:        "",
		parent:      nil,
		left:        l,
		right:       r,
	}
}

// CharStringPair - аналог pair<char, string>
type CharStringPair struct {
	first  byte
	second string
}

// DoubleCharPair - аналог pair<double, char>
type DoubleCharPair struct {
	first  float64
	second byte
}

// String - реализация вывода для CharStringPair
func (p CharStringPair) String() string {
	return fmt.Sprintf("%c %s", p.first, p.second)
}

// String - реализация вывода для DoubleCharPair
func (p DoubleCharPair) String() string {
	return fmt.Sprintf("%f %c", p.first, p.second)
}

// createCode - функция для создания кодов
func createCode(node *HNode, currentCode string) {
	if node == nil {
		return
	}

	node.code = currentCode

	if node.left != nil {
		createCode(node.left, currentCode+"0")
	}
	if node.right != nil {
		createCode(node.right, currentCode+"1")
	}
}

// buildCodeTable - построение таблицы кодов
func buildCodeTable(codeTable []CharStringPair, currentNode *HNode, index *int) {
	if currentNode == nil {
		return
	}
	if currentNode.symbol == '\x00' {
		buildCodeTable(codeTable, currentNode.left, index)
		buildCodeTable(codeTable, currentNode.right, index)
		return
	}
	codeTable[*index] = CharStringPair{currentNode.symbol, currentNode.code}
	*index++
}

// HafmanCode - структура кода Хаффмана
type HafmanCode struct {
	codeTable []CharStringPair
	size      int
}

// NewHafmanCode - конструктор кода Хаффмана
func NewHafmanCode(str string) *HafmanCode {
	// Получаем уникальные символы
	unique := make([]byte, 0)
	for _, s := range str {
		found := false
		for _, u := range unique {
			if u == byte(s) {
				found = true
				break
			}
		}
		if !found {
			unique = append(unique, byte(s))
		}
	}

	// Создаем таблицу вероятностей
	table := make([]DoubleCharPair, len(unique))
	for i, char := range unique {
		count := 0
		for _, s := range str {
			if byte(s) == char {
				count++
			}
		}
		table[i] = DoubleCharPair{
			first:  float64(count) / float64(len(str)),
			second: char,
		}
	}

	// Сортируем по вероятности
	sort.Slice(table, func(i, j int) bool {
		return table[i].first < table[j].first
	})

	// Создаем рабочую таблицу узлов
	workTable := make([]*HNode, len(unique))
	for i := 0; i < len(unique); i++ {
		workTable[i] = NewHNodeFromPair(table[i])
	}

	sizeWorkTable := len(unique)
	for sizeWorkTable != 1 {
		// Создаем новый узел из двух наименьших
		node := NewHNodeFromChildren(workTable[0], workTable[1])
		workTable[0].parent = node
		workTable[1].parent = node
		sizeWorkTable--

		// Сдвигаем элементы
		for i := 0; i < sizeWorkTable; i++ {
			workTable[i] = workTable[i+1]
		}
		workTable[0] = node

		// Сортируем
		sort.Slice(workTable[:sizeWorkTable], func(i, j int) bool {
			return workTable[i].probability < workTable[j].probability
		})
	}

	head := workTable[0]
	createCode(head, "")

	// Создаем таблицу кодов
	codeTable := make([]CharStringPair, len(unique))
	index := 0
	buildCodeTable(codeTable, head, &index)

	result := ""
	for _, char := range str {
		for j := 0; j < len(unique); j++ {
			if codeTable[j].first == byte(char) {
				result += codeTable[j].second
			}
		}
	}
	fmt.Printf("Код: %s\n", result)
	fmt.Println("Таблица кодирования:")
	for i := 0; i < len(unique); i++ {
		fmt.Println(codeTable[i])
	}

	return &HafmanCode{
		codeTable: codeTable,
		size:      len(unique),
	}
}

// decode - декодирование строки
func (h *HafmanCode) decode(code string) {
	result := ""
	decodeSymbol := ""
	for i := 0; i < len(code); i++ {
		decodeSymbol += string(code[i])
		for j := 0; j < h.size; j++ {
			if h.codeTable[j].second == decodeSymbol {
				result += string(h.codeTable[j].first)
				decodeSymbol = ""
			}
		}
	}
	fmt.Printf("Декодированная строка: %s\n", result)
}

func hafman() {
	fmt.Println("Введите строку для кодирования:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	huffman := NewHafmanCode(input)

	fmt.Println("\nВведите код для декодирования (или 'stop' для выхода):")
	scanner.Scan()
	codeToDecode := scanner.Text()

	for codeToDecode != "stop" {
		huffman.decode(codeToDecode)

		fmt.Println("\nВведите следующий код для декодирования (или 'stop' для выхода):")
		scanner.Scan()
		codeToDecode = scanner.Text()
	}
}