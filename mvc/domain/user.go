package domain

type User struct {
	Id        uint64 `json:"id"` //handling json
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
