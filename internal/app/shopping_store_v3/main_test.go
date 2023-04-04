package shopping_store_v3_test

import (
	"github.com/stretchr/testify/suite"
	"os"
	"podlodka/shopping_store/internal/pkg/store"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type testSuite struct {
	suite.Suite
	storage *store.Storage
}

func TestShoppingClient(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (s *testSuite) SetupSuite() {
	db, err := store.ConnectToPostgres()
	if err != nil {
		panic(err)
	}
	s.storage = store.NewStorage(db)
}

func (s *testSuite) TearDownSuite() {
	s.storage.Close()
}

func (s *testSuite) AfterTest(suiteName, testName string) {
}
