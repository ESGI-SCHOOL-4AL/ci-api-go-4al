package text

import (
	"net/http"

	"github.com/System-Glitch/goyave/v2/database"

	"github.com/System-Glitch/goyave/v2"
)

func GetText(response *goyave.Response, request *goyave.Request) {
	db := database.GetConnection()
	response.String(http.StatusOK, "Get text")
}

func GetTextById(response *goyave.Response, request *goyave.Request) {
	db := database.GetConnection()
	// text :=

	// db.First(&, request.Params["id"])
	response.String(http.StatusOK, "Get text by id --> "+request.Params["id"])
}

func ModifyTextById(response *goyave.Response, request *goyave.Request) {
	db := database.GetConnection()

	title := request.String("title")
	content := request.String("content")

	response.String(http.StatusOK, "Modify Text")
}

func AddText(response *goyave.Response, request *goyave.Request) {
	db := database.GetConnection()

	response.String(http.StatusOK, "Add text")
}
