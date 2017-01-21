package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)
	insertRow(db)
	read(db)
	transaction(db)
}

func createTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (`id` int(10) unsigned NOT NULL AUTO_INCREMENT, `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL, `age` int(3) unsigned, PRIMARY KEY (`id`));")
	if err != nil {
		log.Fatal(err)
	}
}

func insertRow(db *sql.DB) {
	result, err := db.Exec(
		"INSERT INTO users (name, age) VALUES (?, ?)",
		"John",
		20,
	)
	if err != nil {
		log.Fatal(err)
	}
	lastInsertId, _ := result.LastInsertId()
	affectedRows, _ := result.RowsAffected()
	fmt.Printf("Last id: %d, affected rows: %d\n", lastInsertId, affectedRows)
}

func read(db *sql.DB) {
	var (
		id, age int
		name    string
	)

	rows, err := db.Query("SELECT id, name, age FROM users WHERE id = ?", 1)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func transaction(db *sql.DB) {
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	result, err := tx.Exec(
		"INSERT INTO users (id, name, age) VALUES (?, ?, ?)",
		1,
		"John",
		20)

	if err != nil {
		tx.Rollback()
		fmt.Println("Transaction has been rolled back")
		log.Fatal(err)
	} else {
		err = tx.Commit()

		if err != nil {
			log.Fatal(err)
		}

		lastInsertId, _ := result.LastInsertId()
		fmt.Printf("Last id: %d\n", lastInsertId)
	}
}
