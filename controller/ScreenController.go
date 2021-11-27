package controller


import (
    "../models"
)

type ScreenController struct {
    Screens map[string]models.Screen
    showHistory map[string]models.Show
}

func NewScreenController() *ScreenController {
    return &ScreenController{}
}

func (c *ScreenController) CreateScreen(name string, screenType models.ScreenType, theaterID string, seats []models.Seat)  {
    Screen := models.NewScreen(name, screenType, theaterID, seats)
    c.Screens[Screen.Id] = *Screen
}

func (c *ScreenController) FindScreenByMovie(movieId string) (screens []models.Screen) {
    for _, s := range c.Screens {
        if s.GetCurrentMovie() == movieId {
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
