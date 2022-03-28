package service

import (
	"encoding/json"
	"io/ioutil"
	"os"

	model "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/model"
)

// GetAllBooks Function that returns all books
func GetAllBooks() (*model.Books, error) {
	var books model.Books
	jsonFile, err := os.Open("pkg/docs/books.json")
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &books)
	defer jsonFile.Close()
	return &books, nil
}

//GetAllAuthors Function that returns all authors
func GetAllAuthors() (*model.Authors, error) {
	var authors model.Authors
	jsonFile, err := os.Open("pkg/docs/authors.json")
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &authors)
	defer jsonFile.Close()
	return &authors, nil
}
