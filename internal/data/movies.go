package data

import (
	"database/sql"
	"time"

	"greenlight.mounirbennacer.com/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genre,omitempty"`
	Rating    float64   `json:"rating"`
	Version   int32     `json:"version"`
	CreatedAt time.Time `json:"-"`
}

type MovieModel struct {
	DB *sql.DB
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(validator.NotBlank(movie.Title), "title", "must not be empty")
	v.Check(validator.MaxChars("title", 500), "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year > 1880, "year", "must be greater than 1880")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must not be empty")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive number")

	v.Check(movie.Genres != nil, "genres", "must not be empty")
	v.Check(len(movie.Genres) > 0, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genre")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}

func (m MovieModel) Insert(movie *Movie) error {
	return nil
}

func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m MovieModel) Delete(id int64) error {
	return nil
}
