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

type Text struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100)"`
	Content string `gorm:"type:text"`
}

func TextGenerator() interface{} {
	text := &Text{}

	text.Title = faker.Name()

	text.Content = faker.Paragraph()

	return text
}
