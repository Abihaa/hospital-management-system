package mongo

import (
	"os"
	"reflect"
	"testing"

	"github.com/abihaa/hospital-management-system/db"
	"github.com/abihaa/hospital-management-system/models"
	"github.com/google/go-cmp/cmp"
)

func Test_client_GetPatientByID(t *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "HMS-mongo-db")

	c, _ := NewClient(db.Option{})
	patient := &models.Patient{Name: "Ertugrl", Age: 50, Gender: "Male", Phone: "0321:2233445", Conditions: "Moderate"}
	_, _ = c.SavePatient(patient)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Patient
		wantErr bool
	}{
		{
			name:    "successful in getting patient details",
			args:    args{id: patient.ID},
			want:    patient,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetPatientByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPatientByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPatientByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_ListPatient(t *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "HMS-mongo-db")

	c, _ := NewClient(db.Option{})
	patient1 := &models.Patient{Name: "Turali", Age: 7, Gender: "Male", Phone: "021-1123456", Conditions: "Moderate"}
	_, _ = c.SavePatient(patient1)
	patient2 := &models.Patient{Name: "Dundar", Age: 15, Gender: "Male", Phone: "0900-78601", Conditions: "Moderate"}
	_, _ = c.SavePatient(patient2)

	type args struct {
		filter map[string]interface{}
		limit  int64
		offset int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.Patient
		wantErr bool
	}{
		{
			name:    "successful: listing with name filter.",
			args:    args{filter: map[string]interface{}{"name": "Turali"}, limit: 1, offset: 0},
			want:    []*models.Patient{patient1},
			wantErr: false,
		},
		{
			name:    "successful: listing with age filter.",
			args:    args{filter: map[string]interface{}{"age": 15}, limit: 1, offset: 0},
			want:    []*models.Patient{patient2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ListPatient(tt.args.filter, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListPatient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ListPatient() got = %v, want = %v, diff = %s", got, tt.want, diff)
			}
		})
	}
}

func Test_client_RemovePatient(t *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "HMS-mongo-db")

	c, _ := NewClient(db.Option{})
	patient := &models.Patient{Name: "Titus", Age: 45, Gender: "Male", Phone: "0321:5544332", Conditions: "Critical"}
	_, _ = c.SavePatient(patient)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "successful in removing Patient.",
			args:    args{id: patient.ID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.RemovePatient(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RemovePatient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_SavePatient(t *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "HMS-mongo-db")

	type args struct {
		patient *models.Patient
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    " successful in Adding new Patient.",
			args:    args{patient: &models.Patient{Name: "Turgut", Age: 45, Gender: "Male", Phone: "0312-3245532", Conditions: "Critical"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			_, err := c.SavePatient(tt.args.patient)
			if (err != nil) != tt.wantErr {
				t.Errorf("SavePatient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_UpdatePatient(t *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "HMS-mongo-db")

	c, _ := NewClient(db.Option{})
	patient := &models.Patient{Name: "Halime", Age: 35, Gender: "Female", Phone: "0321:5532643", Conditions: "moderate"}
	_, _ = c.SavePatient(patient)

	type args struct {
		id      string
		patient *models.Patient
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "successful in Updating Patient.",
			args:    args{id: patient.ID, patient: &models.Patient{ID: patient.ID, Name: "Halime", Age: 36, Gender: "Female", Phone: "0321:5532643", Conditions: "moderate"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.UpdatePatient(tt.args.id, tt.args.patient); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePatient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
