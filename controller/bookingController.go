package controller


import (
    "fmt"

    "../models"
    "../strategy"
)

type BookingController struct {
    bookings map[string]models.Booking
    seatBookingStrategy strategy.SeatLockStrategy
}

func NewBookingController() *BookingController {
    return &BookingController{}
}

func (c *BookingController) ReserveSeat(screenId string, movieId string, seats []models.Seat) error {

    var totalCosts float64

    for _, seat := range seats {
        if c.seatBookingStrategy.IsSeatLocked(seat) {
            return fmt.Errorf("the seat is locked")
        }
        if !seat.IsAvailable() {
            return fmt.Errorf("seat is already booked")
        }
        totalCosts += seat.Cost
    }

    for _, seat := range  seats {
        c.seatBookingStrategy.LockSeat(&seat)
    }

    booking := models.NewBooking(screenId, movieId, totalCosts, seats)
    c.bookings[booking.Id] = *booking
    return nil
}

func (c BookingController) ConfirmPayment(bookingId string) error {
    booking, exists := c.bookings[bookingId]
    if !exists {
        return fmt.Errorf("booking is not created")
    }

    if !booking.Initiated() {
        return fmt.Errorf("booking is not in initiation state")
    }

    booking.ConfirmPayment()
    for _, seat := range booking.Seats {
        seat.BookSeat()
    }
    return nil
}

func (c BookingController) CancelBooking(bookingId string) error  {
    booking, exists := c.bookings[bookingId]
    if !exists {
        return fmt.Errorf("booking is not created")
    }
    if !booking.CanCancel() {
        return fmt.Errorf("booking cannot be cancelled")
    }
    booking.CancelBooking()
    for _, seat := range booking.Seats {
        c.seatBookingStrategy.UnLockSeat(&seat)
    }
    return nil
}

func (c BookingController) UpdateBooking()  {

}

func (c BookingController) GetBooking(bookingID string) (booking models.Booking, err error)  {
    booking, exists := c.bookings[bookingID]
    if !exists {
        return booking, fmt.Errorf("booking is not created")
    }
    return
}