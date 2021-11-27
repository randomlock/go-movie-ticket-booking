package controller

import (
    "../models"
)

type TheaterController struct {
    Theaters map[string]models.Theater
}

func NewTheaterController() *TheaterController {
    return &TheaterController{}
}

func (c *TheaterController) CreateTheaters(location models.Location, name string)  {
    Theater := models.NewTheater(location, name)
    c.Theaters[Theater.Id] = *Theater
}

