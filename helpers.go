package beacon

import (
	"bufio"
	"fmt"
	"os"
)

// Prompt gets text input from the user and returns it as a string.
func Prompt(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	message, _ := reader.ReadString('\n')

	return message
}
