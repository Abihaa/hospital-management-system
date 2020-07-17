package db

import (
	"log"

	"github.com/abihaa/hospital-management-system/models"
)

// DataStore is an interface for query ops.
type DataStore interface {
	SavePatient(patient *models.Patient) (string, error)
	GetPatientByID(id string) (*models.Patient, error)
	RemovePatient(id string) error
	UpdatePatient(id string, patient *models.Patient) error
	ListPatient(filter map[string]interface{}, limit int64, offset int64) ([]*models.Patient, error)
}

// Option holds configuration for data store clients.
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store.
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory.
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
	datastoreFactories[name] = factory
}
