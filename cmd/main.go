package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"tennis-parser/pkg/clients"
	"time"
)

func main() {
	httpClient := http.Client{}
	luzhnikiClient := clients.NewLuznikiAPIClient(httpClient)

	ctx := context.Background()
	date := time.Now().Add(24 * time.Hour)
	slots, err := luzhnikiClient.GetSlotsForDate(ctx, date)
	if err != nil {
		log.Fatal(err)
	}

	availableDates := make(map[time.Time]float64)

	for _, slot := range slots.ByTrainer.NoTrainer.Slots {
		firstSlot := slot[0]
		availableDates[firstSlot.TimeFrom] = firstSlot.Price.From
	}
	fmt.Println(availableDates)
}
