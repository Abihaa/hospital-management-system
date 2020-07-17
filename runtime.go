package runtime

import (
	"github.com/abihaa/hospital-management-system/db"
	"github.com/abihaa/hospital-management-system/db/mongo"
	"github.com/abihaa/hospital-management-system/service"
)

// Runtime initializes values for entry point to our application.
type Runtime struct {
	svc *service.Service
}

// NewRuntime creates a new database service.
func NewRuntime() (*Runtime, error) {
	store, err := mongo.NewClient(db.Option{})
	if err != nil {
		return nil, err
	}
	return &Runtime{svc: service.NewService(store)}, err
}

// Service returns pointer to our service variable.
func (r Runtime) Service() *service.Service {
	return r.svc
}
