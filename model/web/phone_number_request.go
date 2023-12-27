package web

type SavePhoneNumberRequest struct {
	Number   string `json:"number" validate:"required"`
	Provider string `json:"provider" validate:"required"`
}

type EditPhoneNumberRequest struct {
	Number   string `json:"number"`
	Provider string `json:"provider"`
}
