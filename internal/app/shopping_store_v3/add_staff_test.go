package shopping_store_v3_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"math/rand"
	"podlodka/shopping_store/internal/app/shopping_store_v3"
	"podlodka/shopping_store/internal/pkg/repository"
	"podlodka/shopping_store/internal/pkg/testutils"
	desc "podlodka/shopping_store/pkg/api/shopping_store_v3"
)

func (s *testSuite) Test_ShoppingStoreV3_AddStaff() {
	ctx := context.Background()

	rand.Seed(231234)

	repo := repository.NewShoppingStoreRepository(s.storage)
	app := shopping_store_v3.NewShoppingStoreService(repo)

	userID := testutils.GetRandomInt(10000)
	staffID := testutils.GetRandomInt(10000)
	count := int32(3)

	req := &desc.AddStaffRequest{
		UserId: userID,
		Staff: &desc.Staff{
			StaffId: staffID,
			Count:   count,
		},
	}
	resp, err := app.AddStaff(ctx, req)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), resp)

	staffs, err := repo.GetByUserID(ctx, userID)
	require.Len(s.T(), staffs, 1)
	p := staffs[0]
	require.Equal(s.T(), userID, p.UserID)
	require.Equal(s.T(), staffID, p.StaffID)
	require.Equal(s.T(), count, p.Count)
}
