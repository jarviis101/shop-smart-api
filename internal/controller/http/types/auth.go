package types

type (
	AuthUserRequest struct {
		Phone string `json:"phone" validate:"required"`
	}

	TokenResponse struct {
		Token string `json:"token"`
	}
)
