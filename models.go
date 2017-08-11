package TestKnowledge

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