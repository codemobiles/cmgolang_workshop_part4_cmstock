package model

import "time"

type Product struct {
	ID        uint   `gorm:"primary_key"`
	Name      string
	Stock     int64 
	Price     float64
	Image     string 
	CreatedAt time.Time
}
