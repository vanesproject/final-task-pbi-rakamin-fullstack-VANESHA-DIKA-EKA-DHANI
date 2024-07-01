package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `gorm:"primaryKey"`
    Username  string         `gorm:"unique;not null"`
    Email     string         `gorm:"unique;not null"`
    Password  string         `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Photos    []Photo
}
