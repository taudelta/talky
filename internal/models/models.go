package models

import "time"

type Phrase struct {
	ID         int64
	Text       string
	CategoryID int64
	CreatedAt  time.Time
}

type Rating struct {
	ID     int64
	Phrase int64
	Value  float64
}

type View struct {
	ID       int64
	Phrase   int64
	Datetime time.Time
}
