package models

import (
	"time"
)

type Order struct {
	OrderId      uint      `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `gorm:"not null; type:varchar(50)" json:"customerName"`
	OrderedAt    time.Time `gorm:"not null;" json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:OrderId" json:"items"`
}
