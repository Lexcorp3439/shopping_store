package shopping_store_test

import (
	"os"
	"testing"

	"podlodka/shopping_store/internal/app/shopping_store"
	"podlodka/shopping_store/internal/pkg/repository"
	"podlodka/shopping_store/internal/pkg/store"
)

var (
	app  *shopping_store.Implementation
	repo repository.ShoppingStore
)

func TestMain(m *testing.M) {
	db, err := store.ConnectToPostgres()
	if err != nil {
		panic(err)
	}

	storage := store.NewStorage(db)
	repo = repository.NewShoppingStoreRepository(storage)

	app = shopping_store.NewShoppingStoreService(repo)

	code := m.Run()

	os.Exit(code)
}
