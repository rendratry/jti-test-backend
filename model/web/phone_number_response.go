package web

type PhoneNumberResponse struct {
	Id       int    `json:"id"`
	Number   string `json:"number"`
	Provider string `json:"provider"`
	Status   string `json:"status"`
}

type SavePhoneNumberResponse struct {
	Number   string `json:"number"`
	Provider string `json:"provider"`
	Status   string `json:"status"`
}
