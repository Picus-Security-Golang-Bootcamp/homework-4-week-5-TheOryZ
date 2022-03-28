package main

import (
	"log"

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
}
