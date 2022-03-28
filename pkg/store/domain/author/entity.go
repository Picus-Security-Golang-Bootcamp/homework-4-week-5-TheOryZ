package author

import (
	"fmt"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)"`
}

func (Author) TableName() string {
	return "authors"
}
func (a *Author) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, CreatedAt : %s", a.ID, a.Name, a.CreatedAt)
}
func (a *Author) BeforeDelete(*gorm.DB) (err error) {
	fmt.Printf("Author (%s) is deleting.. Say Goodbye..", a.Name)
	return nil
}
