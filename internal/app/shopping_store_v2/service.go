package shopping_store_v2

import (
	"podlodka/shopping_store/internal/pkg/repository"
	desc "podlodka/shopping_store/pkg/api/shopping_store_v2"
)

type Implementation struct {
	desc.UnsafeShoppingStoreV2Server

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
