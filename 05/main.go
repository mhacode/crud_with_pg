package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:0@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.HandleFunc("/books/create", booksCreate)
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

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

func booksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	bk := Book{}
	err := row.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
}

func booksCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	isbn := r.FormValue("isbn")
	title := r.FormValue("title")
	author := r.FormValue("author")
	if isbn == "" || title == "" || author == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	price, err := strconv.ParseFloat(r.FormValue("price"), 32)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	result, err := db.Exec("INSERT INTO books VALUES($1, $2, $3, $4)", isbn, title, author, price)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", isbn, rowsAffected)
}

// curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5.90" localhost:3000/books/create
