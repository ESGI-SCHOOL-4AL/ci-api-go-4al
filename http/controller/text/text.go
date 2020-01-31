package text

import (
	"ci-api-go-4al/database/model"
	"net/http"
	"strconv"

	"github.com/System-Glitch/goyave/v2/database"

	"github.com/System-Glitch/goyave/v2"
)

// Store create a new text record.
func Store(response *goyave.Response, request *goyave.Request) {
	text := model.Text{
		Title:   request.String("title"),
		Content: request.String("content"),
	}

	if err := database.GetConnection().Create(&text).Error; err != nil {
		response.Error(err)
	} else {
		response.JSON(http.StatusCreated, map[string]interface{}{"id": text.ID})
	}
}

// Index list all records.
func Index(response *goyave.Response, request *goyave.Request) {
	texts := []model.Text{}

	if err := database.GetConnection().Find(&texts).Error; err != nil {
		response.Error(err)
	} else {
		response.JSON(http.StatusOK, texts)
	}
}

// Show a single record
func Show(response *goyave.Response, request *goyave.Request) {
	text := model.Text{}
	id, _ := strconv.ParseUint(request.Params["id"], 10, 64)
	result := database.GetConnection().First(&text, id)
	if result.RecordNotFound() {
		response.Status(http.StatusNotFound)
	} else if err := result.Error; err != nil {
		response.Error(err)
	} else {
		response.JSON(http.StatusOK, text)
	}
}

// Update a record.
func Update(response *goyave.Response, request *goyave.Request) {
	text := model.Text{}
	id, _ := strconv.ParseUint(request.Params["id"], 10, 64)

	db := database.GetConnection()
	if db.Select("id").First(&text, id).RecordNotFound() {
		response.Status(http.StatusNotFound)
	} else {
		err := db.Model(&text).Update(model.Text{
			Title:   request.String("title"),
			Content: request.String("content"),
		}).Error

		if err != nil {
			response.Error(err)
		}
	}
}

// Destroy a record.
func Destroy(response *goyave.Response, request *goyave.Request) {
	text := model.Text{}
	id, _ := strconv.ParseUint(request.Params["id"], 10, 64)

	db := database.GetConnection()
	if db.Select("id").First(&text, id).RecordNotFound() {
		response.Status(http.StatusNotFound)
	} else if err := db.Delete(&text).Error; err != nil {
		response.Error(err)
	}
}
