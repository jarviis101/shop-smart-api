package otp

type Sender interface {
	SendOTP() error
}

type sender struct {
}

func CreateSender() Sender {
	return &sender{}
}

func (s *sender) SendOTP() error {
	return nil
}
