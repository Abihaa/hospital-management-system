package service

import (
	"github.com/abihaa/hospital-management-system/models"
)

// SavePatient adds patient into database.
func (s *Service) SavePatient(patient *models.Patient) (string, error) {
	return s.db.SavePatient(patient)
}

// GetPatientByID retrieve patient from database.
func (s *Service) GetPatientByID(id string) (*models.Patient, error) {
	patient, err := s.db.GetPatientByID(id)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

// RemovePatient removes patient from database.
func (s *Service) RemovePatient(id string) error {
	_, err := s.db.GetPatientByID(id)
	if err != nil {
		return err
	}

	return s.db.RemovePatient(id)
}

// UpdatePatient updates patient record in database.
func (s *Service) UpdatePatient(patient *models.Patient) error {
	_, err := s.db.GetPatientByID(patient.ID)
	if err != nil {
		return err
	}

	return s.db.UpdatePatient(patient.ID, patient)
}

// ListPatient lists patient details.
func (s *Service) ListPatient(filter map[string]interface{}, limit int64, offset int64) ([]*models.Patient, error) {
	return s.db.ListPatient(filter, offset, limit)
}
