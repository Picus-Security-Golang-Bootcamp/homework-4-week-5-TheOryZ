package book

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	apiModel "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/api/models/api"
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

//HandleFindAll Get all books
func (b *BookRepository) HandleFindAll(w http.ResponseWriter, r *http.Request) {
	books := b.FindAll()
	var model []apiModel.Book
	for _, book := range books {
		model = append(model, apiModel.Book{
			ID:             int64(book.ID),
			Title:          book.Title,
			NumberOfStocks: int64(book.NumberOfStocks),
			NumberOfPages:  int64(book.NumberOfPages),
			Price:          book.Price,
			Isbn:           book.ISBN,
			ReleaseDate:    book.ReleaseDate,
			AuthorID:       int64(book.AuthorID),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

//HandleFindById Get book by id
func (b *BookRepository) HandleFindById(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	idNumber, _ := strconv.Atoi(id)
	book := b.FindById(idNumber)
	var model apiModel.Book
	model = apiModel.Book{
		ID:             int64(book.ID),
		Title:          book.Title,
		NumberOfStocks: int64(book.NumberOfStocks),
		NumberOfPages:  int64(book.NumberOfPages),
		Price:          book.Price,
		Isbn:           book.ISBN,
		ReleaseDate:    book.ReleaseDate,
		AuthorID:       int64(book.AuthorID),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

//HandleFindByAuthorID Get book by author id
func (b *BookRepository) HandleFindByAuthorID(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	authorID := urlParams["id"]
	authorIDNumber, _ := strconv.Atoi(authorID)
	books := b.FindByAuthorID(authorIDNumber)
	var model []apiModel.Book
	for _, book := range books {
		model = append(model, apiModel.Book{
			ID:             int64(book.ID),
			Title:          book.Title,
			NumberOfStocks: int64(book.NumberOfStocks),
			NumberOfPages:  int64(book.NumberOfPages),
			Price:          book.Price,
			Isbn:           book.ISBN,
			ReleaseDate:    book.ReleaseDate,
			AuthorID:       int64(book.AuthorID),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

//HandleFindByTitle Get book by title
func (b *BookRepository) HandleFindByTitle(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	title := urlParams["title"]
	books := b.FindByTitle(title)
	var model []apiModel.Book
	for _, book := range books {
		model = append(model, apiModel.Book{
			ID:             int64(book.ID),
			Title:          book.Title,
			NumberOfStocks: int64(book.NumberOfStocks),
			NumberOfPages:  int64(book.NumberOfPages),
			Price:          book.Price,
			Isbn:           book.ISBN,
			ReleaseDate:    book.ReleaseDate,
			AuthorID:       int64(book.AuthorID),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

//HandleGetNonDeleted Get non deleted books
func (b *BookRepository) HandleGetNonDeleted(w http.ResponseWriter, r *http.Request) {
	books := b.GetNonDeleted()
	var model []apiModel.Book
	for _, book := range books {
		model = append(model, apiModel.Book{
			ID:             int64(book.ID),
			Title:          book.Title,
			NumberOfStocks: int64(book.NumberOfStocks),
			NumberOfPages:  int64(book.NumberOfPages),
			Price:          book.Price,
			Isbn:           book.ISBN,
			ReleaseDate:    book.ReleaseDate,
			AuthorID:       int64(book.AuthorID),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

//HandleGetByIdWithAuthorName Get book by id and with authors names
func (b *BookRepository) HandleGetByIdWithAuthorName(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	idNumber, _ := strconv.Atoi(id)
	books := b.GetByIdWithAuthorName(idNumber)
	var model []apiModel.BookWithAuthorName
	for _, book := range books {
		model = append(model, apiModel.BookWithAuthorName{
			ID:    int64(book.ID),
			Title: book.Title,
			Name:  book.Name,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

//HandleSumOfBooks Get sum of books
func (b *BookRepository) HandleSumOfBooks(w http.ResponseWriter, r *http.Request) {
	count := b.SumOfBooks()
	var model apiModel.CountModel
	model.Sum = count
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

//HandleInsert Insert new Book
func (b *BookRepository) HandleInsert(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	err := b.Insert(book)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

//HandleUpdate Update book
func (b *BookRepository) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	err := b.Update(book)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

//HandleDelete Delete book
func (b *BookRepository) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	err := b.Delete(book)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

//HandleDeleteById Delete book by id
func (b *BookRepository) HandleDeleteById(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	idNumber, _ := strconv.Atoi(id)
	err := b.DeleteById(idNumber)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
