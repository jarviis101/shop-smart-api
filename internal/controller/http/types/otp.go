package types

type (
	VerifyOTPRequest struct {
		Code string `json:"code" validate:"required"`
	}
)
