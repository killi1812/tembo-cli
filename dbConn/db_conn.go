package dbConn

import (
	"context"
	"fmt"
	"strings"
	"tembo-cli/helpers"

	"github.com/jackc/pgx/v5"
)

var conConf pgx.ConnConfig

func ConnInput() *pgx.Conn {
	var db *pgx.Conn
	var connString string
	for true {
		fmt.Print("Input a connection string: ")
		fmt.Scanln(&connString)
		connString = strings.Trim(connString, "' \"")
		conf, err := pgx.ParseConfig(connString)

		if err != nil {
			fmt.Println("❌ Bad Connection string")
			continue
		}

		conConf = *conf
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

func RunCliInput(db *pgx.Conn) {
	//TODO write a while that takes lines of sql code
	//Command ends when ; is written
	for true {

		fmt.Printf("User: %s on db: %s ~\n", conConf.User, conConf.Database)
		query := helpers.ReadComand()
		fmt.Println("Exacuting query")
		rez, err := db.Exec(context.Background(), query)
		fmt.Printf("%d rows affected\n", rez.RowsAffected())
		if err != nil {
			fmt.Println(query)
			fmt.Println(err)
			fmt.Print("❌")
		} else {
			fmt.Print("✅")
		}
	}

}
