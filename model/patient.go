package model

import (
	"github.com/fatih/structs"
)

// Patient details.
type Patient struct {
	ID        int    `json:"id" structs:"id"`
	Name      string `json:"name" structs:"name"`
	Age       int    `json:"age" structs:"age"`
	Gender    string `json:"gender" structs:"gender"`
	Phone     string `json:"phone" structs:"phone"`
	Condition string `json:"condition" structs:"condition"`
}

// Map converts struct to Map.
func (p *Patient) Map() map[string]interface{} {
	return structs.Map(p)
}

// Names returns a slice of field names.
func (p *Patient) Names() []string {
	fields := structs.Fields(p)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
