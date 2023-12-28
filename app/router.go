package app

import (
	"github.com/julienschmidt/httprouter"
	"jti-test-backend/controller"
	"jti-test-backend/exception"
)

func NewRouter(
	PhoneNumberController controller.PhoneNumberController,
) *httprouter.Router {
	router := httprouter.New()

	router.POST("/jti/api/save", PhoneNumberController.Save)
	router.PUT("/jti/api/edit/:Id", PhoneNumberController.Edit)
	router.DELETE("/jti/api/delete/:Id", PhoneNumberController.Delete)
	router.GET("/jti/api/get-all", PhoneNumberController.GetAll)
	router.GET("/jti/api/get/:Id", PhoneNumberController.GetById)
	router.GET("/jti/api/auto", PhoneNumberController.Auto)

	router.PanicHandler = exception.ErrorHandler

	return router
}
