package model

import (
	"time"
)

const ShoppingStoreTable = "shopping_store"

// ShoppingStore - корзина
type ShoppingStore struct {
	ID        *int64    `db:"id"`
	UserID    int64     `db:"user_id"`
	StaffID   int64     `db:"staff_id"`
	Count     int32     `db:"count"`
	CreatedAt time.Time `db:"created_at"` // Дата создания
	UpdatedAt time.Time `db:"updated_at"` // Дата обновления
}
