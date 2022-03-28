package author

import (
	model "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/model"
	services "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-TheOryZ/pkg/service"

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

//InsertSeedData
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
