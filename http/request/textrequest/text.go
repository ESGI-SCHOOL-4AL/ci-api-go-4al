package textrequest

import "github.com/System-Glitch/goyave/v2/validation"

var (
	// Store validate a store or update request for the Text model.
	Store validation.RuleSet = validation.RuleSet{
		"title":   {"required", "string", "between:1,100"},
		"content": {"required", "string", "min:1"},
	}
)
