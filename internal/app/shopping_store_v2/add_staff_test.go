package shopping_store_v2_test

import (
	"context"
	"math/rand"
	"podlodka/shopping_store/internal/app/shopping_store_v2"
	"podlodka/shopping_store/internal/pkg/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	shopping_repo_mock "podlodka/shopping_store/internal/pkg/mocks/repository"
	desc "podlodka/shopping_store/pkg/api/shopping_store_v2"
)

func Test_SHoppingStoreV2_AddStaff(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	repo := shopping_repo_mock.NewMockShoppingStore(ctrl)

	app := shopping_store_v2.NewShoppingStoreService(repo)

	modelID := rand.Int63n(10000)
	userID := rand.Int63n(10000)
	staffID := rand.Int63n(10000)
	count := int32(3)

	dbResponse := &model.ShoppingStore{
		ID:      &modelID,
		UserID:  userID,
		StaffID: staffID,
		Count:   count,
	}

	repo.EXPECT().Insert(ctx, gomock.Any()).Do(func(ctx context.Context, staff *model.ShoppingStore) {
		require.NotNil(t, staff)
		require.Equal(t, userID, staff.UserID)
		require.Equal(t, staffID, staff.StaffID)
		require.Equal(t, count, staff.Count)
	}).Return(dbResponse, nil).Times(1)

	req := &desc.AddStaffRequest{
		UserId: userID,
		Staff: &desc.Staff{
			StaffId: staffID,
			Count:   count,
		},
	}
	resp, err := app.AddStaff(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
