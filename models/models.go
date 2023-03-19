package models

import (
	"database/sql"
	"fmt"
	"go-rest-api-with-postgresql/config"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// Book schema from table book
// jika return datanya ada yg null, silahkan pake NullString, contohnya dibawah
// Penulis       config.NullString `json:"penulis"`
type Book struct {
	ID             int64  `json:"id"`
	Book_name      string `json:"book_name"`
	Author         string `json:"author"`
	Date_published string `json:"date_published"`
}

// Add data
func AddBook(book Book) int64 {
	// connect to database
	db := config.CreateConnection()

	// close the connection in a rest processing
	defer db.Close()

	// create insert statement
	// return the id of inserted book
	sqlStatement := `INSERT INTO book (book_name, author, date_published) VALUES ($1, $2, $3) RETURNING id`

	// a variable that holds the ID
	var id int64

	// Scan function akan menyimpan insert id didalam id id
	err := db.QueryRow(sqlStatement, book.Book_name, book.Author, book.Date_published).Scan(&id)

	if err != nil {
		log.Fatalf("Cannot execute the query. %v", err)
	}

	fmt.Printf("Insert data single record %v", id)

	// return insert id
	return id
}

// Get a single data
func GetBook(id int64) (Book, error) {
	// Connect to database
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var book Book

	// buat sql query
	sqlStatement := `SELECT * FROM book WHERE id=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&book.ID, &book.Book_name, &book.Author, &book.Date_published)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Data not found!")
		return book, nil
	case nil:
		return book, nil
	default:
		log.Fatalf("Cannot get the data. %v", err)
	}

	return book, err
}

// Get all data
func GetAllBooks() ([]Book, error) {
	// Connect to database
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var books []Book

	// create a select query
	sqlStatement := `SELECT * FROM book`

	// execute the query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("cannot executed the query. %v", err)
	}

	// close the execution query sql
	defer rows.Close()

	// iterate data rows
	for rows.Next() {
		var book Book

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&book.ID, &book.Book_name, &book.Author, &book.Date_published)

		if err != nil {
			log.Fatalf("cannot get the data. %v", err)
		}

		// masukkan ke dalam slice books
		books = append(books, book)

	}

	// return empty buku atau jika error
	return books, err
}

// Update data
func UpdateBook(id int64, book Book) int64 {

	// Connect to database
	db := config.CreateConnection()

	// close connection in rest process
	defer db.Close()

	// create a single query
	sqlStatement := `UPDATE book SET book_name=$2, author=$3, date_published=$4 WHERE id=$1`

	// execute sql statement
	res, err := db.Exec(sqlStatement, id, book.Book_name, book.Author, book.Date_published)

	if err != nil {
		log.Fatalf("Cannot executed query. %v", err)
	}

	// check a total data affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error check rows/data while updating. %v", err)
	}

	fmt.Printf("Total rows/record being affected %v\n", rowsAffected)

	return rowsAffected
}

// Delete data
func DeleteBook(id int64) int64 {

	// Connect to database
	db := config.CreateConnection()

	// close connection in rest process
	defer db.Close()

	// create a sql query
	sqlStatement := `DELETE FROM book WHERE id=$1`

	// execute sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("cannot executed the query. %v", err)
	}

	// check total data deleted
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("cannot find the data. %v", err)
	}

	fmt.Printf("Total data deleted %v", rowsAffected)

	return rowsAffected
}
