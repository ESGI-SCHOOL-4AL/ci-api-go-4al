package test

import (
	"bytes"
	"ci-api-go-4al/database/model"
	"ci-api-go-4al/database/seeder"
	"ci-api-go-4al/http/route"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/System-Glitch/goyave/v2"
	"github.com/System-Glitch/goyave/v2/database"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TextTestSuite struct { // Create a test suite for the Hello controller
	goyave.TestSuite
}

func (suite *TextTestSuite) SetupTest() {
	if database.GetConnection().HasTable(&model.Text{}) {
		seeder.Text()
	}
}

func (suite *TextTestSuite) TearDownTest() {
	suite.ClearDatabase()
}

// Count number of result and check if not nil
func (suite *TextTestSuite) TestGetAll() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/text", nil)
		suite.Nil(err)
		if err == nil {
			json := []*model.Text{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				suite.Len(json, 5)
			}
		}
	})
}

// Count and check if content not nil
func (suite *TextTestSuite) TestGetOne() {
	suite.RunServer(route.Register, func() {
		text := model.Text{Title: "Test", Content: "Test content"}
		database.GetConnection().Create(&text)
		resp, err := suite.Get("/text/"+strconv.FormatUint(uint64(text.ID), 10), nil)
		suite.Nil(err)

		suite.Equal(http.StatusOK, resp.StatusCode)
		if err == nil {
			json := model.Text{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				suite.Equal("Test", json.Title)
				suite.Equal("Test content", json.Content)
			}
		}
	})
}

// Check if the target is deleted
func (suite *TextTestSuite) TestDeleteOne() {
	suite.RunServer(route.Register, func() {
		text := model.Text{Title: "Test", Content: "Test content"}
		database.GetConnection().Create(&text)
		url := "/text/" + strconv.FormatUint(uint64(text.ID), 10)
		resp, err := suite.Delete(url, nil, nil)
		suite.Equal(http.StatusNoContent, resp.StatusCode)

		suite.Nil(err)
		resp, err = suite.Get(url, nil)
		suite.Equal(http.StatusNotFound, resp.StatusCode)
	})
}

// Check if data evolve
func (suite *TextTestSuite) TestUpdateOne() {
	suite.RunServer(route.Register, func() {
		text := model.Text{Title: "Test", Content: "Test content"}
		database.GetConnection().Create(&text)
		url := "/text/" + strconv.FormatUint(uint64(text.ID), 10)
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{"title": "New title", "content": "New content"})
		resp, err := suite.Put(url, headers, bytes.NewReader(body))

		suite.Nil(err)
		suite.Equal(http.StatusNoContent, resp.StatusCode)

		resp, err = suite.Get(url, nil)
		suite.Nil(err)

		suite.Equal(http.StatusOK, resp.StatusCode)
		if err == nil {
			json := model.Text{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				suite.Equal("New title", json.Title)
				suite.Equal("New content", json.Content)
			}
		}

	})
}

// Check if data evolve
func (suite *TextTestSuite) TestStoreOne() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{"title": "New title", "content": "New content"})
		resp, err := suite.Post("/text", headers, bytes.NewReader(body))

		suite.Nil(err)
		suite.Equal(http.StatusCreated, resp.StatusCode)

		if err == nil {
			json := map[string]uint64{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				resp, err = suite.Get("/text/"+strconv.FormatUint(json["id"], 10), nil)
				suite.Nil(err)

				suite.Equal(http.StatusOK, resp.StatusCode)
				if err == nil {
					insertData := model.Text{}
					err := suite.GetJSONBody(resp, &insertData)
					suite.Nil(err)
					if err == nil {
						suite.Equal("New title", insertData.Title)
						suite.Equal("New content", insertData.Content)
					}
				}
			}
		}

	})
}
func TestTextSuite(t *testing.T) { // Run the test suite
	goyave.RunTest(t, new(TextTestSuite))
}
