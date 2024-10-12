package helpers

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func scanln() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(3)
	}
	return line
}

func ReadComand() string {
	var command bytes.Buffer
	var line string

	for true {
		line = scanln()
		line = strings.TrimSpace(line)
		command.WriteString("\n")
		command.WriteString(line)

		if strings.Contains(line, ";") {
			break
		}

	}

	return command.String()
}

func PrintTable(out string) {

}
