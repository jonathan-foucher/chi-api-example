package routers

import (
	"context"
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	db "chi-api-example/database"
	"chi-api-example/models"
)

func InitRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/api/movies", getAllMovies)
	router.Post("/api/movies", saveMovie)
	router.Delete("/api/movies/{movieId}", deleteMovie)

	return router
}

func getAllMovies(res http.ResponseWriter, req *http.Request) {
	queries := db.New(db.GetDbConnection())
	
	fmt.Println("Get all movies")
	results, err := queries.GetMovies(context.Background())
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	}

	movies := make([]models.Movie, len(results))
	for i, result := range results {
	    movies[i] = convertDbMovieToModelsMovie(result)
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func saveMovie(res http.ResponseWriter, req *http.Request) {
	queries := db.New(db.GetDbConnection())

	movie := models.Movie{}
	json.NewDecoder(req.Body).Decode(&movie)

	fmt.Println("Post movie with id", movie.Id)
	fmt.Println("Post movie with title", movie.Title)
	fmt.Println("Post movie with release date", movie.ReleaseDate)
	queries.SaveMovie(context.Background(), db.SaveMovieParams{
		ID: movie.Id,
		Title: movie.Title,
		ReleaseDate: movie.ReleaseDate,
	})
}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	queries := db.New(db.GetDbConnection())

	movieId, err :=  strconv.Atoi(chi.URLParam(req, "movieId"))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	}

	fmt.Println("Delete movie with id", movieId)
	queries.DeleteMovie(context.Background(), int32(movieId))
}

func convertDbMovieToModelsMovie(dbMovie db.Movie) (models.Movie) {
	movie := models.Movie{
		Id: dbMovie.ID,
		Title: dbMovie.Title,
		ReleaseDate: dbMovie.ReleaseDate,
	}
	return movie
}
