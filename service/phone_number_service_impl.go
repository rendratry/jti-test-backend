package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"jti-test-backend/helper"
	"jti-test-backend/model/domain"
	"jti-test-backend/model/web"
	"jti-test-backend/repository"
	"strconv"
)

type PhoneNumberServiceImpl struct {
	PhoneNumberRepository repository.PhoneNumberRepository
	DB                    *gorm.DB
	Validate              *validator.Validate
}

func NewPhoneNumberServiceImpl(phoneNumberRepository repository.PhoneNumberRepository, db *gorm.DB, validate *validator.Validate) *PhoneNumberServiceImpl {
	return &PhoneNumberServiceImpl{
		PhoneNumberRepository: phoneNumberRepository,
		DB:                    db,
		Validate:              validate,
	}
}

func (service *PhoneNumberServiceImpl) Save(ctx context.Context, request web.SavePhoneNumberRequest) web.SavePhoneNumberResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	phoneNumber := domain.PhoneNumber{
		Number:   request.Number,
		Provider: request.Provider,
	}

	service.PhoneNumberRepository.Save(ctx, tx, phoneNumber)

	return helper.ToSavePhoneNumberResponse(phoneNumber)
}

func (service *PhoneNumberServiceImpl) Edit(ctx context.Context, ID int, request web.EditPhoneNumberRequest) web.PhoneNumberResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	phoneNumber := domain.PhoneNumber{
		ID:       ID,
		Number:   request.Number,
		Provider: request.Provider,
	}

	service.PhoneNumberRepository.Edit(ctx, tx, phoneNumber)

	return helper.ToEditPhoneNumberResponse(phoneNumber)
}

func (service *PhoneNumberServiceImpl) Delete(ctx context.Context, ID int) string {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	service.PhoneNumberRepository.Delete(ctx, tx, ID)
	strId := strconv.Itoa(ID)
	return helper.ToDeletePhoneNumberResponse(strId)
}

func (service *PhoneNumberServiceImpl) GetAll(ctx context.Context) []domain.PhoneNumber {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	numbers := service.PhoneNumberRepository.GetAll(ctx, tx)

	return numbers
}

func (service *PhoneNumberServiceImpl) GetById(ctx context.Context, ID int) domain.PhoneNumber {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	number := service.PhoneNumberRepository.GetById(ctx, tx, ID)

	return number
}

func (service *PhoneNumberServiceImpl) Auto(ctx context.Context) string {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	for i := 1; i <= 25; i++ {
		number := helper.RandomPhoneNumber()
		provider := helper.RandomProvider()

		dataToSave := domain.PhoneNumber{
			Number:   number,
			Provider: provider,
		}

		service.PhoneNumberRepository.Save(ctx, tx, dataToSave)
	}

	return "auto generate success"
}
