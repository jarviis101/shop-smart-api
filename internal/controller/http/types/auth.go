package types

type (
	AuthUserRequest struct {
		Channel  string `json:"channel" validate:"required"`
		Resource string `json:"resource" validate:"required"`
	}

	TokenResponse struct {
		Token string `json:"token"`
	}
)
