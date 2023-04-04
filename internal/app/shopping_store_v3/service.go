package shopping_store_v3

import (
	"podlodka/shopping_store/internal/pkg/repository"
	desc "podlodka/shopping_store/pkg/api/shopping_store_v3"
)

type Implementation struct {
	desc.UnsafeShoppingStoreV3Server

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
