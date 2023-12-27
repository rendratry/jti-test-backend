//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"jti-test-backend/app"
	"jti-test-backend/controller"
	"jti-test-backend/middleware"
	"jti-test-backend/repository"
	"jti-test-backend/service"
	"net/http"
)

var phoneNumberSet = wire.NewSet(
	repository.NewPhoneNumberRepositoryImpl,
	wire.Bind(new(repository.PhoneNumberRepository), new(*repository.PhoneNumberRepositoryImpl)),
	service.NewPhoneNumberServiceImpl,
	wire.Bind(new(service.PhoneNumberService), new(*service.PhoneNumberServiceImpl)),
	controller.NewPhoneNumberControllerImpl,
	wire.Bind(new(controller.PhoneNumberController), new(*controller.PhoneNumberControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.GetConnection,
		validator.New,
		phoneNumberSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
