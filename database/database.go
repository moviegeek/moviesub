package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoDatabaseName        = "moviesub"
	mongoMovieFileCollection = "moviefiles"
)

//Database is the structure holds db connection and query interfaces
//used by application and api
type Database struct {
	mongoClient         *mongo.Client
	moviefileCollection *mongo.Collection
}

//New creates a new database instance, connecting to mongodb
func New(URI string, user, password string) (*Database, error) {
	clientOptions := options.Client().ApplyURI(URI).SetAuth(options.Credential{
		Username: user,
		Password: password,
	})

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	moviefileCollection := client.Database(mongoDatabaseName).Collection(mongoMovieFileCollection)

	db := &Database{client, moviefileCollection}

	return db, nil
}

//Close close the database, with the underlying db connection
func (db *Database) Close() error {
	err := db.mongoClient.Disconnect(context.TODO())

	if err != nil {
		return err
	}

	log.Println("Connection to MongoDB closed.")

	return nil
}
