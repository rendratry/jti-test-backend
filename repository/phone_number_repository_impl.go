package repository

import (
	"context"
	"gorm.io/gorm"
	"jti-test-backend/exception"
	"jti-test-backend/helper"
	"jti-test-backend/model/domain"
)

type PhoneNumberRepositoryImpl struct {
}

func NewPhoneNumberRepositoryImpl() *PhoneNumberRepositoryImpl {
	return &PhoneNumberRepositoryImpl{}
}

func (repository *PhoneNumberRepositoryImpl) Save(ctx context.Context, db *gorm.DB, number domain.PhoneNumber) {
	newPhoneNumber := domain.PhoneNumber{
		Number:   number.Number,
		Provider: number.Provider,
	}
	result := db.Create(&newPhoneNumber)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}
}

func (repository *PhoneNumberRepositoryImpl) Edit(ctx context.Context, db *gorm.DB, number domain.PhoneNumber) domain.PhoneNumber {
	newPhoneNumber := domain.PhoneNumber{
		Number:   number.Number,
		Provider: number.Provider,
	}
	result := db.Model(&domain.PhoneNumber{}).Where("id= ?", number.ID).Updates(newPhoneNumber)
	if result.RowsAffected == 0 {
		panic(exception.NewNotFoundError("data not found"))
	}
	return newPhoneNumber
}

func (repository *PhoneNumberRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, ID int) {
	db.Where("id = ?", ID).Delete(&domain.PhoneNumber{})
}

func (repository *PhoneNumberRepositoryImpl) GetAll(ctx context.Context, db *gorm.DB) []domain.PhoneNumber {
	var findNumbers []domain.PhoneNumber
	db.Find(&findNumbers)
	return findNumbers
}

func (repository *PhoneNumberRepositoryImpl) GetById(ctx context.Context, db *gorm.DB, Id int) domain.PhoneNumber {
	findNumber := domain.PhoneNumber{}
	result := db.Where("id = ?", Id).Find(&findNumber)
	if result.RowsAffected == 0 {
		panic(exception.NewNotFoundError("number not found"))
	}
	return findNumber
}
