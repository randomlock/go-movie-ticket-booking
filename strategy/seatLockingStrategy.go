package strategy

import (
    "sync"
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
    *sync.Mutex
}

func (i *InMemoryLocking) LockSeat(seat *models.Seat) {
    i.Mutex.Lock()
    i.lockedSeats[seat.Id] = time.Now()
    i.Mutex.Unlock()
}

func (i *InMemoryLocking) UnLockSeat(seat *models.Seat) {
    i.Mutex.Lock()
    delete(i.lockedSeats, seat.Id)
    i.Mutex.Unlock()

}

func (i InMemoryLocking) IsSeatLocked(seat models.Seat) bool {
    if _, exists := i.lockedSeats[seat.Id]; !exists {
        return false
    }
    return int(time.Now().Sub(i.lockedSeats[seat.Id]).Seconds()) < int(MAX_LOCK_TIME_DURATION)
}

