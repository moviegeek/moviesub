package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moviegeek/moviesub/database"
	"github.com/moviegeek/moviesub/model"
)

//Movies api handler for movies
type Movies struct {
	DB *database.Database
}

//GetAllMovies return all movie files record in db
func (m *Movies) GetAllMovies(ctx *gin.Context) {
	movies, err := m.DB.AllMoviefiles()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	results := []model.MovieFileExternal{}
	for _, m := range movies {
		me := &model.MovieFileExternal{}
		me.FromMovieFile(m)
		results = append(results, *me)
	}

	ctx.JSON(http.StatusOK, results)
}
