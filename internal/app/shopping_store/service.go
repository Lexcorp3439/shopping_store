package shopping_store

import (
	"podlodka/shopping_store/internal/pkg/repository"
	desc "podlodka/shopping_store/pkg/api/shopping_store"
)

type Implementation struct {
	desc.UnsafeShoppingStoreServer

	repo repository.ShoppingStore
}

// NewShoppingStoreService return new instance of Implementation.
func NewShoppingStoreService(
	repo repository.ShoppingStore,
) *Implementation {
	return &Implementation{
		repo: repo,
	}
}
