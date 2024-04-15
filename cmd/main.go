package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	add "github.com/XRS0/gitHelper/internal/handlers/add"
	commit "github.com/XRS0/gitHelper/internal/handlers/commit"
)

func main() {
	Init()
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	fmt.Printf("Commands executed: \n")
	cmds := strings.Split(text, " ")
	printCommands(cmds)
	rangeCommands(cmds)
}

func Init() {
	flag := "ɢɪᴛHᴇʟᴘᴇʀ"
	fmt.Println(flag)
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

	// var arrayFunc []func()
	var wg sync.WaitGroup
	wg.Add(len(commands))

	for _, word := range commands {
		switch word {
		case "commit":
			go commit.GitCommitWithMessage(&wg)
			// arrayFunc = append(arrayFunc, func() {
			// 	commit.GitCommitWithMessage(&wg)
			// })
		case "add":
			go add.GitAdd(&wg)
			// arrayFunc = append(arrayFunc, func() {
			// 	add.GitAdd(&wg)
			// })
		}
	}

	// for _, fn := range arrayFunc {
	// 	go func(f func()) {
	// 		f()
	// 	}(fn)
	// }

	wg.Wait()
}

// file, err := os.Create("data.txt")
// reader := bufio.NewReader(os.Stdin)
// input, err := reader.ReadString('\n')
// _, err = file.WriteString(input)
// file, err := os.Create("data.txt")

// fmt.Printf("combined out:\n%s\n", string(out))
// out, err := cmd.CombinedOutput()
