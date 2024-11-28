package storage

import (
	"context"
	"database/sql"
	"fmt"
	"tennis-parser/pkg/domain"
	"time"
)

const slotsTable = "slots"

type Storage struct {
	conn *sql.DB
}

func NewStorage(conn *sql.DB) *Storage {
	return &Storage{
		conn: conn,
	}
}

func (s *Storage) Add(ctx context.Context, slot domain.BookingSlot) (domain.BookingSlot, error) {
	var (
		id        int64
		createdAt = time.Now().UTC()
	)

	query := fmt.Sprintf("insert into %s(time_from, time_to, price, created_at) values(?, ?, ?, ?) returning id", slotsTable)
	row := s.conn.QueryRowContext(ctx, query, slot.TimeFrom, slot.TimeTo, slot.Price, createdAt)

	if err := row.Scan(id); err != nil {
		return domain.BookingSlot{}, fmt.Errorf("failed to add slot to database row.Scan: %w", err)
	}

	slot.Id = id
	slot.CreatedAt = createdAt

	return slot, nil
}
