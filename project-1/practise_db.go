package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func practiseDB() {

	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres password=kaireke dbname=REST_API_GO sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("CONNECTED to DB")

	result, err := db.Exec(
		"INSERT INTO tasks (title, completed, priority) VALUES ($1, $2, $3)",
		"Learn SQL",
		false,
		3,
	)
	if err != nil {
		log.Fatal(err)
	}

	rows, _ := result.RowsAffected()
	fmt.Println("Rows inserted: ", rows)

	var id int
	err = db.QueryRow(
		"INSERT INTO tasks (title, completed, priority) VALUES ($1, $2, $3) RETURNING id",
		"Learn QueryRow",
		false,
		2,
	).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New task id: ", id)

	var title string
	var completed bool
	var priority int

	err = db.QueryRow(
		"SELECT title, completed, priority FROM tasks where id = $1",
		4,
	).Scan(&title, &completed, &priority)

	if err == sql.ErrNoRows {
		fmt.Println("Task not found")
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("task", title, completed, priority)

	rows2, err := db.Query(
		"SELECT id, title, completed, priority FROM tasks",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows2.Close()

	for rows2.Next() {
		var id int
		var title string
		var completed bool
		var priority int

		err := rows2.Scan(&id, &title, &completed, &priority)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, title, completed, priority)
	}

	res, err := db.Exec(
		"UPDATE tasks set completed = true where id = $1",
		3,
	)
	if err != nil {
		log.Fatal(err)
	}
	affected, _ := res.RowsAffected()
	fmt.Println(affected)

	res, err = db.Exec(
		"DELETE FROM tasks where id = $1",
		3,
	)
	if err != nil {
		log.Fatal(err)
	}

	affected, _ = res.RowsAffected()
	fmt.Println(affected)

}
