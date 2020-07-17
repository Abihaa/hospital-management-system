package handlers

import (
	genModels "github.com/abihaa/hospital-management-system/gen/models"
	domain "github.com/abihaa/hospital-management-system/models"
)

// ToGenPatient for conversion from domainModels to genModels
func ToGenPatient(patient *domain.Patient) *genModels.Patient {
	return &genModels.Patient{
		ID:         patient.ID,
		Name:       patient.Name,
		Age:        int64(patient.Age),
		Gender:     patient.Gender,
		Phone:      patient.Phone,
		Conditions: patient.Conditions}
}

// ToDomainPatient for conversion from genModels to domainModels
func ToDomainPatient(patient *genModels.Patient) *domain.Patient {
	return &domain.Patient{
		ID:         patient.ID,
		Name:       patient.Name,
		Age:        int(patient.Age),
		Gender:     patient.Gender,
		Phone:      patient.Phone,
		Conditions: patient.Conditions,
	}
}
