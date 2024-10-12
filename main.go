package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"tembo-cli/dbConn"
)

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
	db := dbConn.ConnInput()
	defer db.Close(context.Background())

	dbConn.RunCliInput(db)
}
