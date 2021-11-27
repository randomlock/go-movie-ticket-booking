package controller

import (
    "../models"
)

type SeatController struct {
    Seats map[string]models.Seat
}

func NewSeatController() *SeatController {
    return &SeatController{}
}

func (c *SeatController) CreateSeat(seatType models.SeatType, rowNo int, colNo int)  {
    Seat := models.NewSeat(seatType, rowNo, colNo)
    c.Seats[Seat.Id] = *Seat
}

