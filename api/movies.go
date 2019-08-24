package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moviegeek/moviesub/database"
)

//Movies api handler for movies
type Movies struct {
	DB *database.Database
}

//GetAllMovies return all movie files record in db
func (m *Movies) GetAllMovies(ctx *gin.Context) {
	movies, err := m.DB.AllMovies()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, movies)
}
