#go get -u github.com/go-sql-driver/mysql

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Conectarse a la base de datos
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// CREATE operation - creating a new table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(100) NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// INSERT operation - adding data to the table
	res, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", "john_doe", "john.doe@example.com")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted ID:", lastId)

	// READ operation - querying data from the table
	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var username string
		var email string
		if err := rows.Scan(&id, &username, &email); err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID:", id, "Username:", username, "Email:", email)
	}

	// UPDATE operation - updating data in the table
	_, err = db.Exec("UPDATE users SET email = ? WHERE id = ?", "john.newemail@example.com", 1)
	if err != nil {
		log.Fatal(err)
	}

	// DELETE operation - deleting data from the table
	_, err = db.Exec("DELETE FROM users WHERE id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
}
