package models

import "time"

type Show struct {
    BaseModel
    startTime time.Time
    duration time.Duration
    movieID string
    screenId string
}

func NewShow(startTime time.Time, duration time.Duration, movieID string, screenId string) *Show {
    return &Show{BaseModel: NewBaseModel(), startTime: startTime, duration: duration, movieID: movieID, screenId: screenId}
}

func (s Show) IsAvailable() bool {
    return time.Now().Sub(s.startTime) > 0
}

func (s Show) GetMovieId() string {
    return s.movieID
}