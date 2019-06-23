package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/moviegeek/moviesub/config"
	"github.com/moviegeek/moviesub/database"
	"github.com/moviegeek/moviesub/files"
	"github.com/moviegeek/moviesub/model"
	"github.com/moviegeek/moviesub/router"
)

func main() {
	conf := config.Load()

	filescaner := files.New(conf.FileSystem.MediaRootDir)
	db, err := database.New(conf.Database.URI,
		conf.Database.Auth.Username,
		conf.Database.Auth.Password,
	)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	movieFiles := []*model.MovieFile{}

	filescaner.ScanMediaFiles(func(m *model.MovieFile) {
		movieFiles = append(movieFiles, m)
	})

	log.Printf("found %d movies", len(movieFiles))

	for _, m := range movieFiles {
		log.Printf("Add movie to db: %s", m.Title)
		err := db.AddMovieFile(m)
		if err != nil {
			fmt.Printf("add fail: %+v", err)
		}
	}

	engine := router.Create(db, conf)

	addr := fmt.Sprintf("%s:%d", conf.Server.ListenAddr, conf.Server.Port)
	fmt.Println("Started Listening for plain HTTP connection on " + addr)
	log.Fatal(http.ListenAndServe(addr, engine))

}
