package sqlRequest

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

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

func main() {

	fmt.Println("Server: ", server)
	fmt.Println("port: ", port)
	fmt.Println("user: ", user)
	// fmt.Println("password: ", password)
	fmt.Println("database: ", database)

	// Build connection string
	//connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",

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
	fmt.Printf("Connected!\n")

	// Create employee
	createID, err := CreateEmployee("Jake", "United States")
	if err != nil {
		log.Fatal("Error creating Employee: ", err.Error())
	}
	fmt.Printf("Inserted ID: %d successfully.\n", createID)

	// Read employees
	count, err := ReadEmployees()
	if err != nil {
		log.Fatal("Error reading Employees: ", err.Error())
	}
	fmt.Printf("Read %d row(s) successfully.\n", count)

	// Update from database
	updatedRows, err := UpdateEmployee("Jake", "Poland")
	if err != nil {
		log.Fatal("Error updating Employee: ", err.Error())
	}
	fmt.Printf("Updated %d row(s) successfully.\n", updatedRows)

	// Delete from database
	deletedRows, err := DeleteEmployee("Jake")
	if err != nil {
		log.Fatal("Error deleting Employee: ", err.Error())
	}
	fmt.Printf("Deleted %d row(s) successfully.\n", deletedRows)

}

// CreateEmployee inserts an employee record
func CreateEmployee(name string, location string) (int64, error) {
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
      INSERT INTO [test_user].[dbo].[profile] (Name, Location) VALUES (@Name, @Location);
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
		sql.Named("Location", location))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

// ReadEmployees reads all employee records
func ReadEmployees() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees;")

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

// UpdateEmployee updates an employee's information
func UpdateEmployee(name string, location string) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("UPDATE TestSchema.Employees SET Location = @Location WHERE Name = @Name")

	// Execute non-query with named parameters
	result, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("Location", location),
		sql.Named("Name", name))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

// DeleteEmployee deletes an employee from the database
func DeleteEmployee(name string) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("DELETE FROM TestSchema.Employees WHERE Name = @Name;")

	// Execute non-query with named parameters
	result, err := db.ExecContext(ctx, tsql, sql.Named("Name", name))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func FirstLogin(clientId string) {

}
