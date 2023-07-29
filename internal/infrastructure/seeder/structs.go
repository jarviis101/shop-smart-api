package seeder

type User struct {
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
	Phone     string `faker:"phone_number"`
}

type Organization struct {
	Name string `faker:"name"`
}

type Transaction struct {
	TrxNumber string `faker:"uuid_digit"`
}
