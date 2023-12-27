package main

import (
	"jti-test-backend/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: authMiddleware,
	}
}

func main() {
	//server := InitializedServer()
	//
	//err := server.ListenAndServe()
	//helper.PanicIfError(err)
}
