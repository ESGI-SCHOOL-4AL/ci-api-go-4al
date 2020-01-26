package test

import (
	"ci-api-go-4al/database/seeder"
	"ci-api-go-4al/http/route"
	"testing"

	"github.com/System-Glitch/goyave/v2"
)

type TextTestSuite struct { // Create a test suite for the Hello controller
	goyave.TestSuite
}

func (suite *TextTestSuite) SetupTest() {
	suite.ClearDatabase()
	seeder.Text()
}

func (suite *TextTestSuite) TestHello() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/text", nil)
		suite.Nil(err)
		if err == nil {
			json := map[string]interface{}{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				// suite.Equal("value", json["title"])
				// suite.Equal("test", json["content"])
				suite.Equal(json["title"], "test")
			}
		}
	})

}

func TestTextSuite(t *testing.T) { // Run the test suite
	goyave.RunTest(t, new(TextTestSuite))
}
