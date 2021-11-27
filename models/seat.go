package models

type SeatType int
type SeatStatus int

const (
    Lounge SeatType = iota
    Economic
    Premium
)

const (
    BOOKED SeatStatus = iota
    Available
    Locked
)

type Seat struct {
    BaseModel
    seatType SeatType
    rowNo int
    colNo int
    Status SeatStatus
    Cost float64
}

func NewSeat(seatType SeatType, rowNo int, colNo int) *Seat {
    return &Seat{
        BaseModel: NewBaseModel(),
        seatType:  seatType,
        rowNo:     rowNo,
        colNo:     colNo,
    }
}

func (b *Seat) IsAvailable() bool {
    return b.Status == Available
}

func (b *Seat) BookSeat()  {
    b.Status = BOOKED
}

func (b *Seat) FreeSeat()  {
    b.Status = Available
}

func (b *Seat) LockSeat() {
    b.Status = Locked
}

