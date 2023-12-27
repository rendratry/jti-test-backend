package controller

import (
	"github.com/julienschmidt/httprouter"
	"jti-test-backend/helper"
	"jti-test-backend/model/web"
	"jti-test-backend/service"
	"net/http"
	"strconv"
)

type PhoneNumberControllerImpl struct {
	PhoneNumberService service.PhoneNumberService
}

func NewPhoneNumberControllerImpl(phoneNumberService service.PhoneNumberService) *PhoneNumberControllerImpl {
	return &PhoneNumberControllerImpl{PhoneNumberService: phoneNumberService}
}

func (controller *PhoneNumberControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	numberRequest := web.SavePhoneNumberRequest{}
	helper.ReadFromRequestBody(request, &numberRequest)

	numberResponse := controller.PhoneNumberService.Save(request.Context(), numberRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   numberResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *PhoneNumberControllerImpl) Edit(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ID := params.ByName("Id")
	intId, err := strconv.Atoi(ID)
	helper.PanicIfError(err)
	editNumberRequest := web.EditPhoneNumberRequest{}
	helper.ReadFromRequestBody(request, &editNumberRequest)

	editNumberResponse := controller.PhoneNumberService.Edit(request.Context(), intId, editNumberRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   editNumberResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PhoneNumberControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ID := params.ByName("Id")
	intId, err := strconv.Atoi(ID)
	helper.PanicIfError(err)

	deleteResponse := controller.PhoneNumberService.Delete(request.Context(), intId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   deleteResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PhoneNumberControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	getAllNumbersResponse := controller.PhoneNumberService.GetAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getAllNumbersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PhoneNumberControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ID := params.ByName("Id")
	intId, err := strconv.Atoi(ID)
	helper.PanicIfError(err)

	getNumberResponse := controller.PhoneNumberService.GetById(request.Context(), intId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getNumberResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PhoneNumberControllerImpl) Auto(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	autoResponse := controller.PhoneNumberService.Auto(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   autoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
