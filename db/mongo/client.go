package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/abihaa/hospital-management-system/config"
	"github.com/abihaa/hospital-management-system/db"
	domainErr "github.com/abihaa/hospital-management-system/errors"
	"github.com/abihaa/hospital-management-system/models"
)

const (
	patCollection = "Patient"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn *mongo.Client
}

// NewClient initializes a mysql database connection.
func NewClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &client{conn: cli}, nil
}

func (c *client) SavePatient(patient *models.Patient) (string, error) {
	// if patient exists, this updates its record.
	if patient.ID != "" {
		err := c.UpdatePatient(patient.ID, patient)
		return patient.ID, err
	}

	// Otherwise it will create a new patient
	patient.ID = uuid.NewV4().String()
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(patCollection)
	if _, err := collection.InsertOne(context.TODO(), patient); err != nil {
		return "", errors.Wrap(err, "Failed to add patient")
	}

	return patient.ID, nil
}

func (c *client) GetPatientByID(id string) (*models.Patient, error) {
	var pat *models.Patient
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(patCollection)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&pat); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domainErr.NewAPIError(domainErr.NotFound, err)
		}
		return nil, err
	}
	return pat, nil
}

func (c *client) RemovePatient(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(patCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to remove patient '%s' ", id))
	}

	return nil
}

func (c *client) UpdatePatient(id string, patient *models.Patient) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(patCollection)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": patient.ID}, bson.M{"$set": patient}); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to update patient '%s'", id))
	}

	return nil
}

func (c *client) ListPatient(filter map[string]interface{}, limit int64, offset int64) ([]*models.Patient, error) {
	var pat []*models.Patient

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(patCollection)
	cursor, err := collection.Find(context.TODO(), filter, &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domainErr.NewAPIError(domainErr.NotFound, err)
		}
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var p *models.Patient
		err = cursor.Decode(&p)
		if err != nil {
			return nil, errors.Wrap(err, "Error occurred while scanning rows")
		}
		pat = append(pat, p)
	}

	return pat, nil
}
