package controller

import (
    "fmt"

    "../models"
)

type MovieController struct {
    movies map[string]models.Movie
    moviesByGenre map[models.Genre][]models.Movie
}

func NewMovieController() *MovieController {
    return &MovieController{}
}

func (c *MovieController) CreateMovies(name string, description string, genre []models.Genre, rating int, ageRating models.AgeRating)  {
    movie := models.NewMovie(name, description, genre, rating, ageRating)
    c.movies[movie.Id] = *movie

    for _, genre := range movie.Genre() {
        c.moviesByGenre[genre] = append(c.moviesByGenre[genre], *movie)
    }
}

func (c *MovieController) GetMovies(movieId string) (movie models.Movie, err error) {
    if _ , exists := c.movies[movieId]; !exists {
        return movie, fmt.Errorf("invalid movie")
    }
    return c.movies[movieId], nil
}

func (c *MovieController) GetMovieByGenre(genre models.Genre) (movie []models.Movie) {
    return c.moviesByGenre[genre]
}