package main

import (
	"fmt"
	"jti-test-backend/app"
	"jti-test-backend/helper"
	"jti-test-backend/model/domain"
	"testing"
)

func TestNumberRandom(t *testing.T) {
	number := helper.RandomPhoneNumber()
	provider := helper.RandomProvider()
	fmt.Println(number, provider)
}

func TestMigrate(t *testing.T) {
	db := app.GetConnection()
	err := db.AutoMigrate(&domain.PhoneNumber{})
	if err != nil {
		return
	}
}
