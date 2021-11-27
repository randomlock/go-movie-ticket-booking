package models

type Coordinate struct {
    x, y float64
}

func NewCoordinate(x float64, y float64) *Coordinate {
    return &Coordinate{x: x, y: y}
}


type Location struct {
    coordinate Coordinate
    state string
    address string
}

func NewLocation(coordinate Coordinate, state string, address string) *Location {
    return &Location{ coordinate: coordinate, state: state, address: address}
}