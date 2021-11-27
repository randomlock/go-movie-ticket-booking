package models

type Theater struct {
    BaseModel
    location Location
    name string
}

func NewTheater(location Location, name string) *Theater {
    return &Theater{ BaseModel: NewBaseModel(), location: location, name: name}
}