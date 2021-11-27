package models

import (
    "github.com/google/uuid"
)

type BaseModel struct {
    Id string
}

func NewBaseModel() BaseModel {
    return BaseModel{
        Id: uuid.New().String(),
    }
}

