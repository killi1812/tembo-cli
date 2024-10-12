package helpers

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
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

func printLines(lines *[][]string, lengts *[]int) {
	printHeader(((*lines)[0]), lengts)
	for i := 1; i < len(*lines); i++ {
		printLine(&(*lines)[i], lengts)
	}
}

func printHeader(header []string, lenghs *[]int) {
	printLine(&header, lenghs)
	var out bytes.Buffer
	sum := 0
	for i := 0; i < len(*lenghs); i++ {
		sum += (*lenghs)[i]
	}

	for i := 0; i < sum+3*len(*lenghs)-1; i++ {
		out.WriteString("-")
	}
	fmt.Printf("|%s|\n", out.String())
}

func printLine(line *[]string, lenghs *[]int) {
	for i := 0; i < len(*line); i++ {
		minLen := len((*line)[i])
		maxlen := (*lenghs)[i]
		diff := maxlen - minLen

		out := bytes.NewBufferString((*line)[i])

		for i := 0; i < diff; i++ {
			out.WriteString(" ")
		}
		fmt.Printf("| %s ", out.String())
	}
	fmt.Println("|")
}

func prepLines(rows *pgx.Rows) ([][]string, []int) {
	var lines [][]string = make([][]string, 0)
	header := (*rows).FieldDescriptions()
	rowLen := len(header)
	var line []string = make([]string, rowLen)
	var colLens []int = make([]int, rowLen)

	for i := 0; i < len(header); i++ {
		line[i] = header[i].Name
	}
	lines = append(lines, line)

	for (*rows).Next() {
		line = make([]string, rowLen)
		values, err := (*rows).Values()
		for i := 0; i < rowLen; i++ {
			line[i] = fmt.Sprint(values[i])

			if l := len(line[i]); colLens[i] < l {
				colLens[i] = l
			}
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		lines = append(lines, line)
	}

	return lines, colLens
}

func PrintTable(rows *pgx.Rows) {
	lines, lens := prepLines(rows)
	printLines(&lines, &lens)
}
