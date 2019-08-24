package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moviegeek/moviesub/api"
	"github.com/moviegeek/moviesub/config"
	"github.com/moviegeek/moviesub/database"
	"github.com/moviegeek/moviesub/model"
)

//Create create gin server
func Create(db *database.Database, conf *config.Config) *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger(), gin.Recovery())
	g.LoadHTMLGlob("templates/*")

	moviesHandler := api.Movies{
		DB: db,
	}

	apiRoot := g.Group("/api")
	{
		movies := apiRoot.Group("/movies")
		movies.GET("", moviesHandler.GetAllMovies)
	}

	homeRoot := g.Group("")
	{
		homeRoot.GET("", func(ctx *gin.Context) {
			movies, err := db.AllMovies()
			if err != nil {
				movies = []*model.Movie{}
			}

			ctx.HTML(http.StatusOK, "movies.tmpl", gin.H{
				"Movies": movies,
			})
		})
	}

	return g
}
