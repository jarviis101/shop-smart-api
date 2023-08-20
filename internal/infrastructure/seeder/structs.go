package seeder

type User struct {
	Email string `faker:"email"`
	Phone string `faker:"phone_number"`
}

type Organization struct {
	Name string `faker:"name"`
}

type Transaction struct {
	TrxNumber string `faker:"uuid_digit"`
}
