package database

import (
	"context"
	"fmt"
	"log"

	"github.com/moviegeek/moviesub/model"
	"go.mongodb.org/mongo-driver/bson"
)

//AllMoviefiles get all movie files from the database
func (db *Database) AllMoviefiles() ([]*model.MovieFile, error) {
	cur, err := db.moviefileCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Printf("failed to get movies: %v", err)
		return nil, err
	}

	results := []*model.MovieFile{}
	for cur.Next(context.TODO()) {
		var m model.MovieFile
		err := cur.Decode(&m)
		if err != nil {
			log.Printf("failed to deocode item: %v", err)
			continue
		}
		results = append(results, &m)
	}
	return results, nil
}

//UpsertMovieFile update when exists, insert new moviefile when not exist
func (db *Database) UpsertMovieFile(m *model.MovieFile) error {
	//TOOD
	db.moviefileCollection.UpdateOne(context.TODO(), nil, m, nil)
	return nil
}

//AddMovieFile update when exists, insert new moviefile when not exist
func (db *Database) AddMovieFile(m *model.MovieFile) error {
	result, err := db.moviefileCollection.InsertOne(context.TODO(), *m)
	if err != nil {
		return err
	}

	fmt.Printf("insert movie result: %+v", result.InsertedID)
	return nil
}
