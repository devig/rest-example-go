package entity

// User type is a struct for users.
type User struct {
	ID        int64  `json:"id,omitempty"` //если указатель на тип, то поле может быть null
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Sex       int    `json:"sex,omitempty"`
}
