package main

import (
	"log"
	"net/http"

	postgres "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/store/common/db"
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/store/domain/author"
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/store/domain/book"

	"github.com/joho/godotenv"
)

func main() {
	//Set enviroment variables
	err := godotenv.Load("../../env/connection.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init", err)
	}
	log.Println("Postgres connected")

	//Connection DB and migrations
	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migrations()

	bookRepo := book.NewBookRepository(db)
	bookRepo.Migrations()

	//Insert Seed Datas
	authorRepo.InsertSeedData()
	bookRepo.InsertSeedData()
	log.Println("Seed Datas inserted")

	//Handle requests
	http.HandleFunc("/authors", authorRepo.HandleFindAll)
	http.HandleFunc("/authors/{id}", authorRepo.HandleFindById)
	http.HandleFunc("/authors/name/{name}", authorRepo.HandleFindByName)
	http.HandleFunc("/authors/non-deleted", authorRepo.HandleGetNonDeleted)
	http.HandleFunc("/authors/{id}/books", authorRepo.HandleGetByIdWithBooks)
	http.HandleFunc("/authors/sum", authorRepo.HandleSumOfAuthor)
	http.HandleFunc("/authors/insert", authorRepo.HandleInsert)
	http.HandleFunc("/authors/update", authorRepo.HandleUpdate)
	http.HandleFunc("/authors/delete", authorRepo.HandleDelete)
	http.HandleFunc("/authors/delete/{id}", authorRepo.HandleDeleteById)

	http.HandleFunc("/books", bookRepo.HandleFindAll)
	http.HandleFunc("/books/{id}", bookRepo.HandleFindById)
	http.HandleFunc("/books/title/{title}", bookRepo.HandleFindByTitle)
	http.HandleFunc("/books/non-deleted", bookRepo.HandleGetNonDeleted)
	http.HandleFunc("/books/authors/{id}", bookRepo.HandleFindByAuthorID)
	http.HandleFunc("/books/withauthorname/{id}", bookRepo.HandleGetByIdWithAuthorName)
	http.HandleFunc("/books/sum", bookRepo.HandleSumOfBooks)
	http.HandleFunc("/books/insert", bookRepo.HandleInsert)
	http.HandleFunc("/books/update", bookRepo.HandleUpdate)
	http.HandleFunc("/books/delete", bookRepo.HandleDelete)
	http.HandleFunc("/books/delete/{id}", bookRepo.HandleDeleteById)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
