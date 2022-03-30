package book

import (
	"gorm.io/gorm"

	model "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/model"
	services "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/service"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}
func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Book{})
}

//FindAll Get all records <SELECT * FROM Books>
func (b *BookRepository) FindAll() []Book {
	var books []Book
	b.db.Find(&books)
	return books
}

//FindById Get By Id <SELECT * FROM Books WHERE ID = id>
func (b *BookRepository) FindById(id int) Book {
	var book Book
	b.db.Where("id = ?", id).Find(&book)
	return book
}

//FindByAuthorID Get by Author ID <SELECT * FROM Books WHERE AuthorID = authorID ORDER BY ID DESC>
func (b *BookRepository) FindByAuthorID(authorID int) []Book {
	var books []Book
	b.db.Where("author_id = ?", authorID).Order("id desc").Find(&books)
	return books
}

//FindByTitle Get by Title <SELECT * FROM Books WHERE Title = title>
func (b *BookRepository) FindByTitle(title string) []Book {
	var books []Book
	b.db.Where("title LIKE ?", "%"+title+"%").Find(&books)
	return books
}

//GetNonDeleted Get non deleted books
func (b *BookRepository) GetNonDeleted() []Book {
	var books []Book
	b.db.Where("deleted_at = ?", nil).Find(&books)
	return books
}

//GetByIdWithAuthorName Get book by id and with authors names
func (b *BookRepository) GetByIdWithAuthorName(id int) []model.BookWithAuthor {
	var model []model.BookWithAuthor
	b.db.Joins("left join authors on authors.id = books.author_id").
		Where("books.id = ?", id).
		Table("books").
		Select("books.id ,books.title, authors.name").
		Scan(&model)
	return model
}

//Insert Create new Book
func (b *BookRepository) Insert(book Book) error {
	result := b.db.Create(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Update Update book
func (b *BookRepository) Update(book Book) error {
	result := b.db.Save(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete Delete book
func (b *BookRepository) Delete(book Book) error {
	result := b.db.Delete(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//DeleteById Delete by Id book
func (b *BookRepository) DeleteById(id int) error {
	result := b.db.Delete(&Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//InsertSeedData Insert seed data
func (b *BookRepository) InsertSeedData() error {
	books, err := services.GetAllBooks()
	if err != nil {
		return err
	}
	for _, book := range books.Books {
		b.db.Unscoped().Where(Book{Title: book.Title}).
			Attrs(Book{Title: book.Title, NumberOfPages: book.NumberOfPages, NumberOfStocks: book.NumberOfStocks, Price: book.Price, ISBN: book.ISBN, ReleaseDate: book.ReleaseDate, AuthorID: book.AuthorID}).
			FirstOrCreate(&book)
	}
	return nil
}

//SumOfBooks Get sum of books
func (b *BookRepository) SumOfBooks() int64 {
	var count int64
	b.db.Table("books").Count(&count)
	return count
}
