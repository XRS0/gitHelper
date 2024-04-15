package commit

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

func GitCommitWithMessage(wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Printf("Enter a commit name:\n")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Ошибка при чтении данных:", err)
	}
	fullCmd := fmt.Sprintf(`git commit -m "%s"`, text)
	cmd := exec.Command(fullCmd)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
