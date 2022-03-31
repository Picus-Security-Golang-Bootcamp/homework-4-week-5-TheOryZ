package main

import (
	"log"
	"net/http"

	postgres "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/store/common/db"
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/store/domain/author"
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/store/domain/book"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

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
	router := mux.NewRouter()
	//Router validations
	handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	router.Use(loggingMiddleware)

	router.HandleFunc("/authors", authorRepo.HandleFindAll).Methods("GET")
	router.HandleFunc("/authors/{id}", authorRepo.HandleFindById).Methods("GET")
	router.HandleFunc("/authors/name/{name}", authorRepo.HandleFindByName).Methods("GET")
	router.HandleFunc("/authors/{id}/books", authorRepo.HandleGetByIdWithBooks).Methods("GET")
	router.HandleFunc("/authors/non-deleted", authorRepo.HandleGetNonDeleted).Methods("GET")
	router.HandleFunc("/authors/sum", authorRepo.HandleSumOfAuthor).Methods("GET")
	router.HandleFunc("/authors", authorRepo.HandleInsert).Methods("POST")
	router.HandleFunc("/authors", authorRepo.HandleUpdate).Methods("PUT")
	router.HandleFunc("/authors", authorRepo.HandleDelete).Methods("DELETE")
	router.HandleFunc("/authors/{id}", authorRepo.HandleDeleteById).Methods("DELETE")

	router.HandleFunc("/books", bookRepo.HandleFindAll).Methods("GET")
	router.HandleFunc("/books/{id}", bookRepo.HandleFindById).Methods("GET")
	router.HandleFunc("/books/title/{title}", bookRepo.HandleFindByTitle).Methods("GET")
	router.HandleFunc("/books/author/{id}", bookRepo.HandleFindByAuthorID).Methods("GET")
	router.HandleFunc("/books/withauthorname/{id}", bookRepo.HandleGetByIdWithAuthorName).Methods("GET")
	router.HandleFunc("/books/non-deleted", bookRepo.HandleGetNonDeleted).Methods("GET")
	router.HandleFunc("/books/sum", bookRepo.HandleSumOfBooks).Methods("GET")
	router.HandleFunc("/books", bookRepo.HandleInsert).Methods("POST")
	router.HandleFunc("/books", bookRepo.HandleUpdate).Methods("PUT")
	router.HandleFunc("/books", bookRepo.HandleDelete).Methods("DELETE")
	router.HandleFunc("/books/{id}", bookRepo.HandleDeleteById).Methods("DELETE")
	http.Handle("/", router)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

//loggingMiddleware is a middleware that logs the request as it goes in and the response as it goes out.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
