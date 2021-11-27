package models

type Status int 

const (
    Initiated Status = iota
    Booked
    Cancelled
)


type Booking struct {
    BaseModel
    screenId string
    movieId string
    cost   float64
    Seats  []Seat
    status Status
}

func NewBooking(screenId string, movieId string, cost float64, seats []Seat) *Booking {
    return &Booking{BaseModel: NewBaseModel(), screenId: screenId, movieId: movieId, cost: cost, Seats: seats, status: Initiated}
}

func (b *Booking) ConfirmPayment()  {
    b.status = Booked
}

func (b *Booking) CancelBooking()  {
    b.status = Cancelled
}

func (b Booking) Initiated() bool  {
    return b.status == Initiated
}

func (b Booking) CanCancel() bool  {
    return b.status == Booked || b.status == Initiated
}