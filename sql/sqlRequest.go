package sqlRequest

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func CreateUser(name string, location string, email string) (int64, error) {
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
      INSERT INTO [dbo].[profile] (Name, Location, email) VALUES (@Name, @Location,@email);
      select isNull(SCOPE_IDENTITY(), -1);
    `

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", name),
		sql.Named("Location", location),
		sql.Named("email", email),
	)
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

func ReadUser() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM [test_user].[dbo].[profile];")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
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

	return count, nil
}

func CheckUser(ChatID string) (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT count(*) FROM [test_user].[dbo].[profile] where ChatID='%q'", ChatID)

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
	var profileExists int
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
	//return -1, nil
}
