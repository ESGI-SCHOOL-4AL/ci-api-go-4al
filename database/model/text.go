package model

import (
	"github.com/System-Glitch/goyave/v2/database"
	"github.com/bxcodec/faker/v3"
	"github.com/jinzhu/gorm"
)

func init() {
	// All models should be registered in an "init()" function inside their model file.
	database.RegisterModel(&Text{})
}

// Text a notepad record containing a title and a text content.
type Text struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100)"`
	Content string `gorm:"type:text"`
}

// TextGenerator generate a single record of the Text model.
func TextGenerator() interface{} {
	return &Text{
		Title:   faker.Name(),
		Content: faker.Paragraph(),
	}
}
