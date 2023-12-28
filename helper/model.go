package helper

import (
	"jti-test-backend/model/domain"
	"jti-test-backend/model/web"
)

func ToSavePhoneNumberResponse(number domain.PhoneNumber) web.SavePhoneNumberResponse {
	return web.SavePhoneNumberResponse{
		Number:   number.Number,
		Provider: number.Provider,
		Status:   "save success",
	}
}

func ToEditPhoneNumberResponse(number domain.PhoneNumber) web.PhoneNumberResponse {
	return web.PhoneNumberResponse{
		Id:       number.ID,
		Number:   number.Number,
		Provider: number.Provider,
		Status:   "edit success",
	}
}

func ToDeletePhoneNumberResponse(Id string) string {
	return "id " + Id + " success deleted"
}

func ToGetPhoneNumberResponse(number domain.PhoneNumber) web.PhoneNumberResponse {
	return web.PhoneNumberResponse{
		Id:       number.ID,
		Number:   number.Number,
		Provider: number.Provider,
	}
}
