package seeder

import (
	"ci-api-go-4al/database/model"

	"github.com/System-Glitch/goyave/v2/database"
)

// Text seed 5 random Text records.
func Text() {
	database.NewFactory(model.TextGenerator).Save(5)
}
