package middleware

import (
	"github.com/joho/godotenv"
	"jti-test-backend/helper"
	"jti-test-backend/model/web"
	"log"
	"net/http"
	"os"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file apakah yang ini")
	}

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	writer.Header().Set("Access-Control-Expose-Headers", "Authorization")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-api-key")
	if request.Method == http.MethodOptions {
		writer.WriteHeader(http.StatusOK)
		return
	}
	xApiKey := os.Getenv("X_API_KEY")
	if xApiKey == request.Header.Get("X-API-Key") {
		// ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
