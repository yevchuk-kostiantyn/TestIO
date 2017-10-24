package TestIO

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Student struct {
	Name string
	Surname string
	Group string
	Username string
	Password string
}

type Instructor struct {
	Name string
	Surname string
	Subject string
	Username string
	Password string
}

type Response struct {
	Position string
}

type NewUser struct {
	FirstName string 	`json:"first_name"`
	LastName string 	`json:"last_name"`
	Email string 		`json:"email"`
	Password string 	`json:"password"`
	Position string 	`json:"position"`
}