package models

import (
    "time"
    "gorm.io/gorm"
)

type Photo struct {
    ID        uint           `gorm:"primaryKey"`
    Title     string         `gorm:"not null"`
    Caption   string
    PhotoUrl  string         `gorm:"not null"`
    UserID    uint
    CreatedAt time.Time
    UpdatedAt time.Time
}