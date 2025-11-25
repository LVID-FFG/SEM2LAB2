package main

import (
	"fmt"
)

type MChain struct {
	key  string
	data string
	next *MChain
}

func genHashMorf(size int, key string) int {
	result := 1245
	for i := 0; i < len(key); i++ {
		result += (i * int(key[i])) % size
	}
	return result % size
}

type HashTableMChain struct {
	table []*MChain
	size  int
}

func NewHashTableMChain(sz int) *HashTableMChain {
	table := make([]*MChain, sz)
	return &HashTableMChain{
		table: table,
		size:  sz,
	}
}

func (h *HashTableMChain) insert(key, value string) {
	Hash := genHashMorf(h.size, key)

	if h.table[Hash] == nil {
		h.table[Hash] = &MChain{key: key, data: value}
	} else {
		address := h.table[Hash]
		for address.next != nil && address.key != key {
			address = address.next
		}
		if address.key == key {
			address.data = value
		} else {
			address.next = &MChain{key: key, data: value}
		}
	}
}

func (h *HashTableMChain) find(key string) string {
	Hash := genHashMorf(h.size, key)
	address := h.table[Hash]
	if address == nil {
		return ""
	}
	for address != nil {
		if address.key == key {
			return address.data
		}
		address = address.next
	}
	return ""
}

type morfTable struct {
	tableIn  *HashTableMChain
	tableOut *HashTableMChain
}

func newMorfTable() *morfTable {
	return &morfTable{
		tableIn:  NewHashTableMChain(10),
		tableOut: NewHashTableMChain(10),
	}
}

func (m *morfTable) searchSimbolMorf(str1, str2 string, symbol string) bool {
	existingMapping := m.tableIn.find(symbol)

	if existingMapping != "" {
		for i := 0; i < len(str1); i++ {
			if str1[i] == symbol[0] && str2[i] != existingMapping[0] {
				return false
			}
		}
		return true
	}

	candidate := byte(0)
	candidateFound := false

	for i := 0; i < len(str1); i++ {
		if str1[i] == symbol[0] {
			if !candidateFound {
				candidate = str2[i]
				candidateFound = true
			} else if str2[i] != candidate {
				return false
			}
		}
	}

	candidateStr := string(candidate)
	reverseMapping := m.tableOut.find(candidateStr)
	if reverseMapping != "" && reverseMapping != symbol {
		return false
	}

	m.tableIn.insert(symbol, candidateStr)
	m.tableOut.insert(candidateStr, symbol)
	return true
}

func isMorf(str1, str2 string) {
	if len(str1) != len(str2) {
		fmt.Println("FALSE")
		return
	}

	table := newMorfTable()

	for i := 0; i < len(str1); i++ {
		symbol := string(str1[i])
		if !table.searchSimbolMorf(str1, str2, symbol) {
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
	isMorf(str1, str2)
}
