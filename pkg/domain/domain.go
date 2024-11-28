package domain

import "time"

type BookingSlot struct {
	Id        int64
	TimeFrom  time.Time
	TimeTo    time.Time
	Price     int
	CreatedAt time.Time
}
