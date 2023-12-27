package repository

import (
	"context"
	"gorm.io/gorm"
	"jti-test-backend/model/domain"
)

type PhoneNumberRepository interface {
	Save(ctx context.Context, db *gorm.DB, number domain.PhoneNumber)
	Edit(ctx context.Context, db *gorm.DB, number domain.PhoneNumber) domain.PhoneNumber
	Delete(ctx context.Context, db *gorm.DB, ID int)
	GetAll(ctx context.Context, db *gorm.DB) []domain.PhoneNumber
	GetById(ctx context.Context, db *gorm.DB, Id int) domain.PhoneNumber
}
