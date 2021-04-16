package sqlRequest

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	//"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB
var server = "localhost"
var port = 1434
var user = `adm_test` // SORESHNIKOVPC\soreshnikov
var password = "Qq1234567890"
var database = "test_user"

func Test() {
	fmt.Println("Done")
}

func SqlCon() error {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("DB connected!\n")
	return err
}

func CreateUser(name string, lastname string, email string, chatid int64) (int64, error) {
	ctx := context.Background()
	var err error

	if db == nil {
		err = errors.New("CreateEmployee: db is null")
		return -1, err
	}

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := `
      INSERT INTO [dbo].[profile1] (name, lastname, email, chatid) VALUES (@name, @lastname,@email,@chatid);
      select isNull(SCOPE_IDENTITY(), -1);
    `

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("name", name),
		sql.Named("lastname", lastname),
		sql.Named("email", email),
		sql.Named("chatid", chatid),
	)
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

func ReadUser() (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT id, name, lastname,chatid FROM [test_user].[dbo].[profile1];")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int64

	// Iterate through the result set.
	for rows.Next() {
		var name, lastname string
		var id, chatid int

		// Get values from row.
		err := rows.Scan(&id, &name, &lastname, &chatid)
		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %d, Name: %s, lastname: %s, ChatID: %d\n", id, name, lastname, chatid)
		count++
	}

	return count, nil
}

func CheckUser(ChatID int64) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT count(*) FROM [test_user].[dbo].[profile1] where chatid='%d'", ChatID)

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	// Iterate through the result set.

	/*
		for rows.Next() {
			var name, location string
			var id int

			// Get values from row.
			err := rows.Scan(&id, &name, &location)
			if err != nil {
				return -1, err
			}

			fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
			count++
		}
	*/
	var count int
	var profileExists int
	for rows.Next() {
		// Get values from row.
		err = rows.Scan(&profileExists)
		if err != nil {
			return -1, err
		}
		if profileExists != 0 {
			return 1, err // 1 = exists
		} else {
			return 0, err // 0 = not exists
		}
		count++
	}
	return -1, nil
}
