package main

import (
	_ "ci-api-go-4al/http/request"
	"ci-api-go-4al/http/route"

	"fmt"

	"github.com/System-Glitch/goyave/v2"
	_ "github.com/System-Glitch/goyave/v2/validation"

	// Import the approriate GORM dialect for the database you're using.
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
	// This is the entry point of your application.
	// Most applications don't need more than this, but
	// if you are running multiple services, such as a
	// websocket server, you'll need to run Goyave in a routine.
	// See: https://system-glitch.github.io/goyave/guide/advanced/multi-services.html
	goyave.RegisterStartupHook(func() {
		fmt.Println("Server started.")
	})
	goyave.Start(route.Register)
	fmt.Println("Server shutdown.")
}
