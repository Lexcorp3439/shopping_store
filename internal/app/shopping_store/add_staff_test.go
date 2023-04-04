package shopping_store_test

import (
	"context"
	"podlodka/shopping_store/internal/pkg/testutils"
	"testing"

	"github.com/stretchr/testify/require"
	desc "podlodka/shopping_store/pkg/api/shopping_store"
)

func Test_ShoppingStoreV1_AddStaff(t *testing.T) {
	ctx := context.Background()

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
	t.Log(req)
	resp, err := app.AddStaff(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	staffs, err := repo.GetByUserID(ctx, userID)
	t.Log(staffs)
	require.Len(t, staffs, 1)
	s := staffs[0]
	require.Equal(t, userID, s.UserID)
	require.Equal(t, staffID, s.StaffID)
	require.Equal(t, count, s.Count)
}
