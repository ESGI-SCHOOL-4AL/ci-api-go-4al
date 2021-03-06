package route

import (
	"ci-api-go-4al/http/controller/text"
	"ci-api-go-4al/http/request/textrequest"

	"github.com/System-Glitch/goyave/v2"
	"github.com/System-Glitch/goyave/v2/cors"
)

// Routing is an essential part of any Goyave application.
// Routes definition is the action of associating a URI, sometimes having
// parameters, with a handler which will process the request and respond to it.

// Routes are defined in routes registrer functions.
// The main route registrer is passed to "goyave.Start()" and is executed
// automatically with a newly created root-level router.

// Register all the application routes. This is the main route registrer.
func Register(router *goyave.Router) {

	// Applying default CORS settings (allow all methods and all origins)
	// Learn more about CORS options here: https://system-glitch.github.io/goyave/guide/advanced/cors.html
	router.CORS(cors.Default())

	router.Route("GET", "/robots.txt", func(resp *goyave.Response, req *goyave.Request) {
		resp.File("resources/robots.txt")
	}, nil)

	router.Route("GET", "/text", text.Index, nil)
	router.Route("POST", "/text", text.Store, textrequest.Store)

	router.Route("GET", "/text/{id:[0-9]+}", text.Show, nil)
	router.Route("PUT", "/text/{id:[0-9]+}", text.Update, textrequest.Store)
	router.Route("DELETE", "/text/{id:[0-9]+}", text.Destroy, nil)
}
