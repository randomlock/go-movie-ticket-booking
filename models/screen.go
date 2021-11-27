package models

type ScreenType int

const (
    ScreenType2D ScreenType = iota
    ScreenType3D
    ScreenTypeImax
) 


type Screen struct {
    BaseModel
    name string
    screenType ScreenType
    theaterID string
    movieID string
    Seats   []Seat
}

func NewScreen(name string, screenType ScreenType, theaterID string, movieId string, seat []Seat) *Screen {
    return &Screen{
        BaseModel:  NewBaseModel(),
        name:       name,
        screenType: screenType,
        theaterID:  theaterID,
        movieID:    movieId,
        Seats:      seat,
    }
}

func (s Screen) MovieID() string {
    return s.movieID
}




