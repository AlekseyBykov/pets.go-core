package receivers

import (
	"fmt"
	"github.com/AlekseyBykov/pets.go-core/internal/structs/custom_type"
)

type User struct {
	Name string
	Age  custom_type.Age
}

func (u *User) PrintUserInfo() {
	u.Name = "Renamed"
	fmt.Println(u.Name, u.Age)
}
