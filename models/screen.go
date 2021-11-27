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
    currentShow Show
    Seats   []Seat
}

func NewScreen(name string, screenType ScreenType, theaterID string, seat []Seat) *Screen {
    return &Screen{
        BaseModel:  NewBaseModel(),
        name:       name,
        screenType: screenType,
        theaterID:  theaterID,
        Seats:      seat,
    }
}

func (s *Screen) SetCurrentShow(show Show)  {
    s.currentShow = show
}

func (s Screen) CurrentShow() Show {
    return s.currentShow
}

func (s Screen) GetCurrentMovie() string  {
    return s.currentShow.GetMovieId()
}




