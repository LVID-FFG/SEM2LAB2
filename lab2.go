package main

import (
	"fmt"
	"os"
)

// Command - перечисление команд
type Command int

const (
	EXIT Command = iota
	ASTEROID
	GENOME
	DICTIONARY
	HAFMAN
	OPENHASH
	CHAINHASH
	MORF
	LRUCASH
	ERROR
)

func main() {
	for {
		fmt.Println("Доступные программы: ")
		fmt.Println("Выход - EXIT")
		fmt.Println("ASTEROID")
		fmt.Println("GENOME")
		fmt.Println("DICTIONARY")
		fmt.Println("HAFMAN")
		fmt.Println("OPENHASH")
		fmt.Println("CHAINHASH")
		fmt.Println("MORF")
		fmt.Println("LRUCASH")

		var usCin string
		fmt.Scan(&usCin)

		var command Command

		switch usCin {
		case "EXIT":
			command = EXIT
		case "ASTEROID":
			command = ASTEROID
		case "GENOME":
			command = GENOME
		case "DICTIONARY":
			command = DICTIONARY
		case "HAFMAN":
			command = HAFMAN
		case "OPENHASH":
			command = OPENHASH
		case "CHAINHASH":
			command = CHAINHASH
		case "MORF":
			command = MORF
		case "LRUCASH":
			command = LRUCASH
		default:
			command = ERROR
		}

		switch command {
		case EXIT:
			fmt.Println("Выход из программы")
			os.Exit(0)
		case ASTEROID:
			fmt.Println("\nЗапуск задания №1")
			asteroid()
		case GENOME:
			fmt.Println("\nЗапуск задания №3")
			genome()
		case DICTIONARY:
			fmt.Println("\nЗапуск задания №4")
			dictionary()
		case HAFMAN:
			fmt.Println("\nЗапуск задания №5")
			hafman()
		case OPENHASH:
			fmt.Println("\nЗапуск задания №6")
			hashTableFree()
		case CHAINHASH:
			fmt.Println("\nЗапуск задания №7")
			hashTableChain()
		case MORF:
			fmt.Println("\nЗапуск задания №8")
			morf()
		case LRUCASH:
			fmt.Println("\nЗапуск задания №9")
			LRU()
		default:
			fmt.Println("Ошибка ввода!")
			os.Exit(1)
		}

		command = ERROR
		fmt.Println()
	}
}
