package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

type Book struct {
	isbn string
	title string
	author string
	price float32
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to the database")
}

func main() {
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.ListenAndServe(":8080", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		books = append(books, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range books {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

func booksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	bk := Book{}
	err := row.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
}
