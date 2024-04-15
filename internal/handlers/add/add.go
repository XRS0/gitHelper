package add

import (
	"fmt"
	"os/exec"
)

func GitAdd(file string) error {
	cmd := exec.Command("git", "add", file)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	return nil
}
