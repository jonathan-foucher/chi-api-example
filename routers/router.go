package routers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "chi-api-example/models"
)
func InitRouter() *chi.Mux {
	router := chi.NewRouter()
    router.Use(middleware.Logger)

    router.Get("/api/movies", func(res http.ResponseWriter, req *http.Request) {
    	fmt.Println("Get all movies")
        res.Write([]byte("Get all movies"))
    })

    router.Post("/api/movies", func(res http.ResponseWriter, req *http.Request) {
    	movie := models.Movie{}
		json.NewDecoder(req.Body).Decode(&movie)

		fmt.Println("Post movie with id", movie.Id)
		fmt.Println("Post movie with title", movie.Title)
		fmt.Println("Post movie with release date", movie.ReleaseDate)
    })

    router.Delete("/api/movies/{movieId}", func(res http.ResponseWriter, req *http.Request) {
   		movieId, _ :=  strconv.Atoi(chi.URLParam(req, "movieId"))
     	fmt.Println("Delete movie with id", movieId)
    })

    return router
}