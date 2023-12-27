package service

import (
	"context"
	"jti-test-backend/model/domain"
	"jti-test-backend/model/web"
)

type PhoneNumberService interface {
	Save(ctx context.Context, request web.SavePhoneNumberRequest) web.PhoneNumberResponse
	Edit(ctx context.Context, ID int, request web.EditPhoneNumberRequest) web.PhoneNumberResponse
	Delete(ctx context.Context, ID int) string
	GetAll(ctx context.Context) []domain.PhoneNumber
	GetById(ctx context.Context, ID int) web.PhoneNumberResponse
	Auto(ctx context.Context) string
}
