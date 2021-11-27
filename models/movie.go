package models

type Genre int
type AgeRating int

const (
    Romance Genre = iota
    Comedy
    Drama
    Thriller
)

const (
    kid AgeRating = iota
    adult
    all
)


type Movie struct {
    BaseModel
    name string
    description string
    genre []Genre
    rating int
    ageRating AgeRating
}

func (m Movie) Genre() []Genre {
    return m.genre
}

func NewMovie(name string, description string, genre []Genre, rating int, ageRating AgeRating) *Movie {
    return &Movie{ BaseModel: NewBaseModel(), name: name, description: description, genre: genre, rating: rating, ageRating: ageRating}
}
