package shopping_store_v2_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	shopping_repo_mock "podlodka/shopping_store/internal/pkg/mocks/repository"
	"podlodka/shopping_store/internal/pkg/model"
	"podlodka/shopping_store/internal/pkg/testutils"
)

func Test_ShoppingStore_Mock(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	repo := shopping_repo_mock.NewMockShoppingStore(ctrl)

	modelID := testutils.GetRandomInt(10000)
	userID := testutils.GetRandomInt(10000)
	staffID := testutils.GetRandomInt(10000)
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
	}).Return(dbResponse, nil).
		Times(0)
}
