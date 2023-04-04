package repository

import (
	"context"
	"fmt"
	"podlodka/shopping_store/internal/pkg/model"
	"podlodka/shopping_store/internal/pkg/store"
)

// ShoppingStore repo interface
type ShoppingStore interface {
	Insert(ctx context.Context, staff *model.ShoppingStore) (*model.ShoppingStore, error)
	Update(ctx context.Context, staff *model.ShoppingStore) (*model.ShoppingStore, error)
	Delete(ctx context.Context, staff *model.ShoppingStore) error
	GetByUserID(ctx context.Context, userID int64) ([]model.ShoppingStore, error)
}

type storeRepo struct {
	store *store.Storage
}

func NewShoppingStoreRepository(db *store.Storage) ShoppingStore {
	return &storeRepo{store: db}
}

func (r *storeRepo) Insert(ctx context.Context, staff *model.ShoppingStore) (*model.ShoppingStore, error) {
	stmt := store.PSQL().
		Insert(model.ShoppingStoreTable).
		Columns("user_id", "staff_id", "count").
		Values(staff.UserID, staff.StaffID, staff.Count).
		Suffix("RETURNING *")

	err := r.store.Selectx(ctx, stmt,
		&staff.ID,
		&staff.StaffID,
		&staff.UserID,
		&staff.Count,
		&staff.CreatedAt,
		&staff.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return staff, nil
}

func (r *storeRepo) Update(ctx context.Context, staff *model.ShoppingStore) (*model.ShoppingStore, error) {
	stmt := store.PSQL().
		Update(model.ShoppingStoreTable).
		Set("count", staff.Count).
		Where("user_id = ?", staff.UserID).
		Where("staff_id = ?", staff.StaffID).
		Suffix("RETURNING *")

	err := r.store.Selectx(ctx, stmt,
		&staff.ID,
		&staff.StaffID,
		&staff.UserID,
		&staff.Count,
		&staff.CreatedAt,
		&staff.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return staff, nil
}

func (r *storeRepo) Delete(ctx context.Context, staff *model.ShoppingStore) error {
	stmt := store.PSQL().
		Delete(model.ShoppingStoreTable).
		Where("user_id = ?", staff.UserID).
		Where("staff_id = ?", staff.StaffID)

	err := r.store.Execx(ctx, stmt)
	if err != nil {
		return err
	}
	return nil
}

func (r *storeRepo) GetByUserID(ctx context.Context, userID int64) ([]model.ShoppingStore, error) {
	stmt := store.PSQL().
		Select("*").
		From(model.ShoppingStoreTable).
		Where("user_id = ?", userID)

	fmt.Println(stmt.ToSql())

	return store.Select[model.ShoppingStore](ctx, r.store, stmt)
	//var result []model.ShoppingStore
	//rows, err := r.store.Select(ctx, stmt, &result)
	//if err != nil {
	//	return nil, err
	//}
	//return pgx.CollectRows(rows, pgx.RowToStructByName[model.ShoppingStore])
}
