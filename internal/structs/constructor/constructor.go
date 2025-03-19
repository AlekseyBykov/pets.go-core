package constructor

import (
	"github.com/AlekseyBykov/pets.go-core/internal/structs/custom_type"
)

type User struct {
	Name string
	Age  custom_type.Age
}

func NewUser(name string, age custom_type.Age) User {
	return User{
		Name: name,
		Age:  age,
	}
}

func NewUserPtr(name string, age custom_type.Age) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}
