package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/jackc/pgx/v5"
)

func connInput() *pgx.Conn {
	var db *pgx.Conn
	var connString string
	for true {
		fmt.Print("Input a connection string: ")
		fmt.Scanln(&connString)
		connString = strings.Trim(connString, "' \"")
		conConf, err := pgx.ParseConfig(connString)
		if err != nil {
			fmt.Println("❌ Bad Connection string")
			continue
		}
		dbo, err := pgx.Connect(context.Background(), conConf.ConnString())

		if err == nil {
			fmt.Println("✅ Connected to remote database")
			db = dbo
			break
		}
		fmt.Println("❌ failed to connect")
		fmt.Println(err)
	}
	return db
}

func runCliInput(db *pgx.Conn) {

	//TODO write a while that takes lines of sql code
	//Command ends when ; is written
	for true {
		query := readComand()
		fmt.Println("Exacuting query")
		rez, err := db.Exec(context.Background(), query)
		fmt.Printf("%d rows affected\n", rez.RowsAffected())
		if err != nil {
			fmt.Print("❌")
		} else {
			fmt.Print("✅")
		}
	}

}

func scanln() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(3)
	}
	return line
}

func readComand() string {
	var command bytes.Buffer
	var line string

	//TODO add user and database names
	fmt.Printf("~\n")
	for true {
		line = scanln()
		line = strings.TrimSpace(line)
		//command.WriteString("\n")
		command.WriteString(line)

		if strings.Contains(line, ";") {
			break
		}

	}

	return command.String()
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			signal.Ignore(sig)
			fmt.Println("\nBye 👋")
			os.Exit(0)
		}
	}()
	db := connInput()
	defer db.Close(context.Background())
	runCliInput(db)
}
