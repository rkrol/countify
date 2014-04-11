package models

import (
	"time"
)

type Expense struct {
	Id        int
	Name      string
	Amount    float64
	Date      time.Time
	AccountId int
	PayerId   int
}
