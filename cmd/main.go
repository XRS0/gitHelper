package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	add "github.com/XRS0/gitHelper/internal/handlers/add"
	// commit "github.com/XRS0/gitHelper/internal/handlers/commit"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Ошибка при чтении данных:", err)
		return
	}

	fmt.Printf("Выполнены комманды: \n")
	printCommands(strings.Split(text, " "))
}

func printCommands(commands []string) {
	if len(commands) == 0 {
		return
	}
	for _, word := range commands {
		fmt.Printf("%s\n", word)
	}
}

func rangeCommands(commands []string) {
	if len(commands) == 0 {
		return
	}
	for _, word := range commands {
		switch word {
		case "commit":
			add.GitAdd(".")
		}
	}
}

// file, err := os.Create("data.txt")
// reader := bufio.NewReader(os.Stdin)
// input, err := reader.ReadString('\n')
// _, err = file.WriteString(input)
// file, err := os.Create("data.txt")
