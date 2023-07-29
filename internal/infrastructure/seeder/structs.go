package seeder

type User struct {
	FirstName  string   `faker:"first_name"`
	LastName   string   `faker:"last_name"`
	MiddleName string   `faker:"-"`
	Phone      string   `faker:"phone_number"`
	Roles      []string `faker:"-"`
}

type Organization struct {
	OwnerID int64  `faker:"-"`
	Name    string `faker:"name"`
}
