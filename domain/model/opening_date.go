package model

import "time"

type OpeningDate struct {
	Day  uint
	From time.Time
	To   time.Time
}
