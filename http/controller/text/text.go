package text

import (
	"ci-api-go-4al/database/model"
	"net/http"
	"strconv"

	"github.com/System-Glitch/goyave/v2/database"

	"github.com/System-Glitch/goyave/v2"
)

func Store(response *goyave.Response, request *goyave.Request) {
	text := model.Text{
		Title:   request.String("title"),
		Content: request.String("content"),
	}

	database.GetConnection().Create(&text)
	response.JSON(http.StatusCreated, map[string]interface{}{"id": text.ID})
}

func Index(response *goyave.Response, request *goyave.Request) {
	texts := []model.Text{}

	database.GetConnection().Find(&texts)
	response.JSON(http.StatusOK, texts)
}

func Show(response *goyave.Response, request *goyave.Request) {
	text := model.Text{}
	id, _ := strconv.ParseUint(request.Params["id"], 10, 64)

	if database.GetConnection().First(&text, id).RecordNotFound() {
		response.Status(http.StatusNotFound)
	} else {
		response.JSON(http.StatusOK, text)
	}
}

func Update(response *goyave.Response, request *goyave.Request) {
	text := model.Text{}
	id, _ := strconv.ParseUint(request.Params["id"], 10, 64)

	db := database.GetConnection()
	if db.Select("id").First(&text, id).RecordNotFound() {
		response.Status(http.StatusNotFound)
	} else {
		db.Model(&text).Update(model.Text{
			Title:   request.String("title"),
			Content: request.String("content"),
		})
	}
}

func Destroy(response *goyave.Response, request *goyave.Request) {
	text := model.Text{}
	id, _ := strconv.ParseUint(request.Params["id"], 10, 64)

	db := database.GetConnection()
	if db.Select("id").First(&text, id).RecordNotFound() {
		response.Status(http.StatusNotFound)
	} else {
		db.Delete(&text)
	}
}
