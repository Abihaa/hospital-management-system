package model

import (
	"reflect"
	"testing"
)

func TestPatient_Map(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		Age       int
		Gender    string
		Phone     string
		Condition string
	}
	tests := []struct {
		name     string
		fields   fields
		expected map[string]interface{}
	}{
		{
			name: " successful",
			fields: fields{
				ID:        1,
				Name:      "ironman",
				Age:       40,
				Gender:    "male",
				Phone:     "0312-4567890",
				Condition: "moderate",
			},
			expected: map[string]interface{}{

				"id":        1,
				"name":      "ironman",
				"age":       40,
				"gender":    "male",
				"phone":     "0312-4567890",
				"condition": "moderate",
			},
		},
		{
			name: "Success in case of fewer fields",
			fields: fields{
				ID:     1,
				Name:   "ironman",
				Age:    40,
				Gender: "male",
				Phone:  "0312-4567890",
			},
			expected: map[string]interface{}{
				"id":        1,
				"name":      "ironman",
				"age":       40,
				"gender":    "male",
				"phone":     "0312-4567890",
				"condition": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Patient{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Age:       tt.fields.Age,
				Gender:    tt.fields.Gender,
				Phone:     tt.fields.Phone,
				Condition: tt.fields.Condition,
			}
			if got := p.Map(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Map() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPatient_Names(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		Age       int
		Gender    string
		Phone     string
		Condition string
	}
	tests := []struct {
		name     string
		fields   fields
		expected []string
	}{
		{
			name: " successful",
			fields: fields{
				ID:        1,
				Name:      "ironman",
				Age:       40,
				Gender:    "male",
				Phone:     "0312-4567890",
				Condition: "moderate",
			},
			expected: []string{"id", "name", "age", "gender", "phone", "condition"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Patient{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Age:       tt.fields.Age,
				Gender:    tt.fields.Gender,
				Phone:     tt.fields.Phone,
				Condition: tt.fields.Condition,
			}
			if got := p.Names(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Names() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
