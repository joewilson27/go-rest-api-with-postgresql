package controller

import (
	"encoding/json" // package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	"fmt"
	"strconv" // package yang digunakan untuk mengubah string menjadi tipe int

	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api

	"go-rest-api-with-postgresql/models" //models package dimana Buku didefinisikan

	"github.com/gorilla/mux" // digunakan untuk mendapatkan parameter dari router
	_ "github.com/lib/pq"    // postgres golang driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message`
	Data    []models.Book `json:"data"`
}

// Add data
func AddBook(w http.ResponseWriter, r *http.Request) {

	// create an empty book of type models.Book
	// kita buat empty buku dengan tipe models.Book
	var book models.Book

	// decode data json request ke buku
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatalf("Cannot be decoded from the request body.  %v", err)
	}

	// call a method AddBook from models then get the return value ID
	insertID := models.AddBook(book)

	// format response objectnya
	res := response{
		ID:      insertID,
		Message: "Data buku telah ditambahkan",
	}

	// kirim response
	json.NewEncoder(w).Encode(res)
}

// Get a single data with parameter id
func GetBook(w http.ResponseWriter, r *http.Request) {
	// Set a header
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get the id book from the parameter requested, the key is "id"
	params := mux.Vars(r)

	// convert the id from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Cannot converted an id to type int.  %v", err)
	}

	// call a function GetBook from models.go with id parameter to get a single data
	buku, err := models.GetBook(int64(id))

	if err != nil {
		log.Fatalf("Can't get the data. %v", err)
	}

	// kirim response
	json.NewEncoder(w).Encode(buku)
}

// Get all data
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// call func GetAllBooks from models.go
	books, err := models.GetAllBooks()

	if err != nil {
		log.Fatalf("Cannot get the data. %v", err)
	}

	var response Response
	response.Status = 200
	response.Message = "Success"
	response.Data = books

	// sent response
	json.NewEncoder(w).Encode(response)
}

// Update data
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	// get request parameter id
	params := mux.Vars(r)

	// convert an id to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Cannot convert to int.  %v", err)
	}

	// create a book variable with type models.Book
	var book models.Book

	// decode json request to a book variable
	err = json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatalf("Cannot decode request body.  %v", err)
	}

	// call func UpdateBook
	updatedRows := models.UpdateBook(int64(id), book)

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Book has been successfully updated. Total updated data %v rows/record", updatedRows)

	// ini adalah format response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// sent response
	json.NewEncoder(w).Encode(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	// get request parameter id
	params := mux.Vars(r)

	// convert an id to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Cannot convert to int.  %v", err)
	}

	// call func DeleteBook, and convert int ke int64
	deletedRows := models.DeleteBook(int64(id))

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Book has been successfully deleted. Total data yang dihapus %v", deletedRows)

	// ini adalah format reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
