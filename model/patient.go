package model

import (
	"github.com/fatih/structs"
)

// Patient details.
type Patient struct {
	ID         string `json:"id" structs:"id" db:"id" bson:"_id"`
	Name       string `json:"name" structs:"name" db:"name" bson:"name"`
	Age        int    `json:"age" structs:"age" db:"age" bson:"age"`
	Gender     string `json:"gender" structs:"gender" db:"gender" bson:"gender"`
	Phone      string `json:"phone" structs:"phone" db:"phone" bson:"phone"`
	Conditions string `json:"conditions" structs:"conditions" db:"conditions" bson:"conditions"`
}

// Map conversion.
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
