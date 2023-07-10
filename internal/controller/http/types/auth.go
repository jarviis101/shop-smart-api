package types

type (
	LoginUserRequest struct {
		Phone string `json:"phone" validate:"required"`
	}

	TokenResponse struct {
		Token string `json:"token"`
	}
)
