package book

import (
	"fmt"

	authorRepo "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/store/domain/author"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title          string
	NumberOfPages  int
	NumberOfStocks int
	Price          float64
	ISBN           string
	ReleaseDate    string
	AuthorID       int
	Author         authorRepo.Author `gorm:"foreignKey:AuthorID;"`
}

func (Book) TableName() string {
	return "books"
}
func (b *Book) ToString() string {
	return fmt.Sprintf("ID : %d, Title : %s, Number Of Stocks : %d, Price : %f, ISBN : %s, Release Date : %s ", b.ID, b.Title, b.NumberOfStocks, b.Price, b.ISBN, b.ReleaseDate)
}
func (b *Book) BeforeDelete(*gorm.DB) (err error) {
	fmt.Printf("Book (%s) is deleting.. Say Goodbye..", b.Title)
	return nil
}
