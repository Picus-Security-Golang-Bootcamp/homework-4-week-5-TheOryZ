package author

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/model"
	services "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}
func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&Author{})
}

//FindAll Get all records <SELECT * FROM Authors>
func (a *AuthorRepository) FindAll() []Author {
	var authors []Author
	a.db.Find(&authors)
	return authors
}

//FindById Get By Id <SELECT * FROM Authors WHERE ID = id>
func (a *AuthorRepository) FindById(id int) Author {
	var author Author
	a.db.Where("id = ?", id).Find(&author)
	return author
}

//FindByName Get by Name <SELECT * FROM Authors WHERE name = name>
func (a *AuthorRepository) FindByName(name string) []Author {
	var authors []Author
	a.db.Where("name LIKE ?", "%"+name+"%").Find(&authors)
	return authors
}

//GetNonDeleted Get non deleted authors
func (a *AuthorRepository) GetNonDeleted() []Author {
	var authors []Author
	a.db.Where("deleted_at = ?", nil).Find(&authors)
	return authors
}

//GetByIdWithBooks Get Author by id and with books
func (a *AuthorRepository) GetByIdWithBooks(id int) []model.BookWithAuthor {
	var model []model.BookWithAuthor
	a.db.Joins("left join books on authors.id = books.author_id").
		Where("authors.id = ?", id).
		Table("authors").
		Select("books.id ,books.title, authors.name").
		Scan(&model)
	return model
}

//Insert Create new Author
func (a *AuthorRepository) Insert(author Author) error {
	result := a.db.Create(&author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Update Update book
func (a *AuthorRepository) Update(author Author) error {
	result := a.db.Save(&author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete Delete author
func (a *AuthorRepository) Delete(author Author) error {
	result := a.db.Delete(&author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//DeleteById Delete by Id author
func (a *AuthorRepository) DeleteById(id int) error {
	result := a.db.Delete(&Author{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//InsertSeedData Insert seed data
func (a *AuthorRepository) InsertSeedData() error {
	authors, err := services.GetAllAuthors()
	if err != nil {
		return err
	}
	for _, author := range authors.Authors {
		a.db.Unscoped().Where(Author{Name: author.Name}).
			Attrs(Author{Name: author.Name}).
			FirstOrCreate(&author)
	}
	return nil
}

//SumOfAuthor Get sum of all authors
func (a *AuthorRepository) SumOfAuthor() int64 {
	var count int64
	a.db.Table("authors").Count(&count)
	return count
}

//HandleFindAll Find all authors
func (a *AuthorRepository) HandleFindAll(w http.ResponseWriter, r *http.Request) {
	authors := a.FindAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authors)
}

//HandleFindById Find by id
func (a *AuthorRepository) HandleFindById(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	idNumber, _ := strconv.Atoi(id)
	author := a.FindById(idNumber)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
}

//HandleFindByName Find by name
func (a *AuthorRepository) HandleFindByName(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	name := urlParams["name"]
	authors := a.FindByName(name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authors)
}

//HandleGetNonDeleted Get non deleted authors
func (a *AuthorRepository) HandleGetNonDeleted(w http.ResponseWriter, r *http.Request) {
	authors := a.GetNonDeleted()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authors)
}

//HandleGetByIdWithBooks Get author by id and with books
func (a *AuthorRepository) HandleGetByIdWithBooks(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	idNumber, _ := strconv.Atoi(id)
	books := a.GetByIdWithBooks(idNumber)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

//HandleSumOfAuthor Get sum of all authors
func (a *AuthorRepository) HandleSumOfAuthor(w http.ResponseWriter, r *http.Request) {
	sum := a.SumOfAuthor()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sum)
}

//HandleInsert Insert new author
func (a *AuthorRepository) HandleInsert(w http.ResponseWriter, r *http.Request) {
	var author Author
	_ = json.NewDecoder(r.Body).Decode(&author)
	err := a.Insert(author)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

//HandleUpdate Update author
func (a *AuthorRepository) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var author Author
	_ = json.NewDecoder(r.Body).Decode(&author)
	err := a.Update(author)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
}

//HandleDelete Delete author
func (a *AuthorRepository) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var author Author
	_ = json.NewDecoder(r.Body).Decode(&author)
	err := a.Delete(author)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
}

//HandleDeleteById Delete by id author
func (a *AuthorRepository) HandleDeleteById(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	idNumber, _ := strconv.Atoi(id)
	err := a.DeleteById(idNumber)
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
