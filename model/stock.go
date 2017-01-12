package model

import (
	"time"
)

type Stock struct {
	ID             int
	StockName      string
	StockId        string
	STOCKCHINANAME string
	CREATEDAT      time.Time
	UPDATEAT       time.Time
}