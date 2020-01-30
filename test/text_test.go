package test

import (
	"ci-api-go-4al/database/seeder"
	"ci-api-go-4al/http/route"
	"testing"

	"github.com/System-Glitch/goyave/v2"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TextTestSuite struct { // Create a test suite for the Hello controller
	goyave.TestSuite
}

func (suite *TextTestSuite) SetupTest() {
	suite.ClearDatabase()
	seeder.Text()
}

// Count number of result and check if not nil
func (suite *TextTestSuite) TestGetAll() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/text", nil)
		suite.Nil(err)
		if err == nil {
			json := map[string]interface{}{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				suite.Equal(json["title"], "test")
			}
		}
	})
}

// Count and check if content not nil
func (suite *TextTestSuite) TestGetOne() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/text/1", nil)
		suite.Nil(err)
		if err == nil {
			json := map[string]interface{}{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				suite.Equal("test", json["title"])
			}
		}
	})
}

// Check error code
func (suite *TextTestSuite) TestGetUnexistedOne() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/text/666", nil)
		suite.Nil(err)
		if err == nil {
			json := map[string]interface{}{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				suite.Equal("test", json["title"])
			}
		}
	})
}

// Check if the target is deleted
func (suite *TextTestSuite) TestDeleteOne() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/text/1", nil)
		suite.Nil(err)
		if err == nil {
			json := map[string]interface{}{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				suite.Equal("test", json["title"])
			}
		}
	})
}

// Check if data evolve
func (suite *TextTestSuite) TestUpdateOne() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/text/1", nil)
		suite.Nil(err)
		if err == nil {
			json := map[string]interface{}{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				suite.Equal("test", json["title"])
			}
		}
	})
}
func TestTextSuite(t *testing.T) { // Run the test suite
	goyave.RunTest(t, new(TextTestSuite))
}
