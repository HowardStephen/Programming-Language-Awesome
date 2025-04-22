package user

import "fmt"

type User struct {
	name string
	uuid string
}

func (u *User) DefineUser(name string, uuid string) {
	u.name = name
	u.uuid = uuid
}

func (u *User) ShowUserInfo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if u.name == "" || u.uuid == "" {
		panic("User does not define!")
	}

	fmt.Println(u.name, u.uuid)
}
