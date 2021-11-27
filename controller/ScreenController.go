package controller


import (
    "../models"
)

type ScreenController struct {
    Screens map[string]models.Screen
}

func NewScreenController() *ScreenController {
    return &ScreenController{}
}

func (c *ScreenController) CreateScreen(name string, screenType models.ScreenType, theaterID string, movieId string, seats []models.Seat)  {
    Screen := models.NewScreen(name, screenType, theaterID, movieId, seats)
    c.Screens[Screen.Id] = *Screen
}

func (c *ScreenController) FindScreenByMovie(movieId string) (screens []models.Screen) {
    for _, s := range c.Screens {
        if s.MovieID() == movieId {
            screens = append(screens, s)
        }
    }
    return
}

func (c *ScreenController) FindSeat(screenID string) (seats []models.Seat)  {

    screen := c.Screens[screenID]
    for _, s := range screen.Seats {
        if s.IsAvailable() {
            seats = append(seats, s)
        }
    }
    return
}
