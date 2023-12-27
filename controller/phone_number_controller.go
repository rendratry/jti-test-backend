package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PhoneNumberController interface {
	Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Edit(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Auto(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
