package add

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func GitAdd(wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Printf("Enter a commit name:\n")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "add", text)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
