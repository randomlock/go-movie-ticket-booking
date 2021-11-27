package strategy

import (
    "time"

    "../models"
)

type SeatLockStrategy interface {
    LockSeat(seat *models.Seat)
    UnLockSeat(seat *models.Seat)
    IsSeatLocked(seat models.Seat) bool
}

const MAX_LOCK_TIME_DURATION = time.Second * 120

type InMemoryLocking struct {
    lockedSeats map[string]time.Time
}

func (i *InMemoryLocking) LockSeat(seat *models.Seat) {
    i.lockedSeats[seat.Id] = time.Now()
    seat.LockSeat()
}

func (i *InMemoryLocking) UnLockSeat(seat *models.Seat) {
    delete(i.lockedSeats, seat.Id)
    seat.FreeSeat()
}

func (i InMemoryLocking) IsSeatLocked(seat models.Seat) bool {
    return int(time.Now().Sub(i.lockedSeats[seat.Id]).Seconds()) < int(MAX_LOCK_TIME_DURATION)
}

