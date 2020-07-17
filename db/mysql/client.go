package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/abihaa/hospital-management-system/config"
	"github.com/abihaa/hospital-management-system/db"
	domainErr "github.com/abihaa/hospital-management-system/errors"
	"github.com/abihaa/hospital-management-system/models"
)

const (
	patientTableName = "patient"
)

func init() {
	db.Register("mysql", NewClient)
}

//The first implementation.
type client struct {
	db *sqlx.DB
}

func formatDSN() string {
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cfg.DBName = viper.GetString(config.DbName)
	cfg.ParseTime = true
	cfg.User = viper.GetString(config.DbUser)
	cfg.Passwd = viper.GetString(config.DbPass)
	return cfg.FormatDSN()
}

// NewClient initializes a mysql database connection.
func NewClient(conf db.Option) (db.DataStore, error) {
	log().Info("initializing mysql connection: " + formatDSN())
	cli, err := sqlx.Connect("mysql", formatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}
	return &client{db: cli}, nil
}

// To add  patient details.
func (c *client) SavePatient(patient *models.Patient) (string, error) {
	if patient.ID != "" {
		err := c.UpdatePatient(patient.ID, patient)
		return patient.ID, err
	}

	patient.ID = uuid.NewV4().String()

	if _, err := c.db.NamedExec(fmt.Sprintf(`INSERT INTO %s (%s) VALUES(%s)`, patientTableName, strings.Join(patient.Names(), ","), strings.Join(mkPlaceHolder(patient.Names(), ":", func(name, prefix string) string {
		return prefix + name
	}), ",")), patient); err != nil {
		return "", errors.Wrap(err, "failed to add patient details.")
	}

	return patient.ID, nil
}

// update patient details.
func (c *client) UpdatePatient(id string, patient *models.Patient) error {
	if _, err := c.db.NamedExec(fmt.Sprintf(`UPDATE %s SET %s WHERE id = '%s'`, patientTableName, strings.Join(mkPlaceHolder(patient.Names(), "=:", func(name, prefix string) string {
		return name + prefix + name
	}), ","), id), patient); err != nil {
		return errors.Wrap(err, "failed to update patient")
	}
	return nil
}

// To retrieve patient data by ID.
func (c *client) GetPatientByID(id string) (*models.Patient, error) {
	var pat models.Patient
	if err := c.db.Get(&pat, fmt.Sprintf(`SELECT * FROM %s WHERE id = '%s'`, patientTableName, id)); err != nil {
		if err == sql.ErrNoRows {
			return nil, domainErr.NewAPIError(domainErr.NotFound, err)
		}
		return nil, err
	}
	return &pat, nil
}

// To remove patient.
func (c *client) RemovePatient(id string) error {
	if _, err := c.db.Exec(fmt.Sprintf(`DELETE FROM %s WHERE id= '%s'`, patientTableName, id)); err != nil {
		return errors.Wrap(err, "failed to delete patient.")
	}

	return nil
}

// To list all available patient details.
func (c *client) ListPatient(filter map[string]interface{}, limit int64, offset int64) ([]*models.Patient, error) {
	var pat []*models.Patient

	if err := c.db.Select(&pat, fmt.Sprintf("SELECT * FROM %s %s LIMIT %b, %b", patientTableName, mkFilter(filter), offset, limit)); err != nil {
		if err == sql.ErrNoRows {
			return pat, domainErr.NewAPIError(domainErr.NotFound, err)
		}
		return pat, err
	}

	return pat, nil
}

// Placeholder function.
func mkPlaceHolder(names []string, prefix string, formatName func(name, prefix string) string) []string {
	ph := make([]string, len(names))
	for i, name := range names {
		ph[i] = formatName(name, prefix)
	}

	return ph
}

// Filter for where clause option in ListPatient.
func mkFilter(filter map[string]interface{}) string {
	if len(filter) < 1 {
		return ""
	}

	whereFilter := make([]string, 0)
	for k, v := range filter {
		whereFilter = append(whereFilter, fmt.Sprintf(" %s = '%s' ", k, v))
	}

	return fmt.Sprintf("WHERE %s", strings.Join(whereFilter, " AND "))
}
