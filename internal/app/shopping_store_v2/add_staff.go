package shopping_store_v2

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"podlodka/shopping_store/internal/pkg/model"
	desc "podlodka/shopping_store/pkg/api/shopping_store_v2"
)

func (i *Implementation) AddStaff(ctx context.Context, req *desc.AddStaffRequest) (*desc.AddStaffResponse, error) {
	if err := validateCreateUser(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err := i.repo.Insert(ctx, &model.ShoppingStore{
		UserID:  req.UserId,
		StaffID: req.Staff.StaffId,
		Count:   req.Staff.Count,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.AddStaffResponse{}, nil
}

func validateCreateUser(req *desc.AddStaffRequest) error {
	err := validation.ValidateStruct(req,
		validation.Field(&req.UserId, validation.Required),
		validation.Field(&req.Staff, validation.Required),
	)
	if err != nil {
		return err
	}
	err = validation.ValidateStruct(req.Staff,
		validation.Field(&req.Staff.StaffId, validation.Required),
		validation.Field(&req.Staff.Count, validation.Required),
	)
	return err
}
