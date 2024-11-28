package parser

import "time"

type BookingSlotDTO struct {
	TimeFrom time.Time
	TimeTo   time.Time
	Price    int
}
