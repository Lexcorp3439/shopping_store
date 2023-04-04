package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"podlodka/shopping_store/internal/app/shopping_store"
	"podlodka/shopping_store/internal/app/shopping_store_v2"
	"podlodka/shopping_store/internal/app/shopping_store_v3"
	"podlodka/shopping_store/internal/pkg/repository"
	"podlodka/shopping_store/internal/pkg/store"

	desc "podlodka/shopping_store/pkg/api/shopping_store"
	desc2 "podlodka/shopping_store/pkg/api/shopping_store_v2"
	desc3 "podlodka/shopping_store/pkg/api/shopping_store_v3"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	db, err := store.ConnectToPostgres()
	if err != nil {
		panic(err)
	}

	storage := store.NewStorage(db)
	repo := repository.NewShoppingStoreRepository(storage)

	shoppingStoreV1 := shopping_store.NewShoppingStoreService(repo)
	shoppingStoreV2 := shopping_store_v2.NewShoppingStoreService(repo)
	shoppingStoreV3 := shopping_store_v3.NewShoppingStoreService(repo)

	s := grpc.NewServer()
	desc.RegisterShoppingStoreServer(s, shoppingStoreV1)
	desc2.RegisterShoppingStoreV2Server(s, shoppingStoreV2)
	desc3.RegisterShoppingStoreV3Server(s, shoppingStoreV3)

	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
