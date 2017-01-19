package model

import (
	"time"
)

type StockList struct {
	ID             int
	StockName      string
	StockId        string
	STOCKCHINANAME string
	CREATEDAT      time.Time
	UPDATEAT       time.Time
}