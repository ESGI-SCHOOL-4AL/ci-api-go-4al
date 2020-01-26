package textrequest

import "github.com/System-Glitch/goyave/v2/validation"

var (
	Store validation.RuleSet = validation.RuleSet{
		"title":   {"required", "string", "between:1,255"},
		"content": {"required", "string", "min:1"},
	}
)
