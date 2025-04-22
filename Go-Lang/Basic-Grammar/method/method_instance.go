package main

import "fmt"

type User struct {
	name string
	uuid string
}

// u value method
func (u User) echoUserInfo() {
	fmt.Println(u.name, ": ", u.uuid)
}

// u pointer method
func (u *User) changeUserName(name string) {
	u.name = name
}

func main() {
	// value
	newUser := User{
		name: "henry",
		uuid: "2100770220",
	}
	newUser.echoUserInfo()
	newUser.changeUserName("green")
	newUser.echoUserInfo()

	// pointer
	newUserPtr := &newUser
	newUserPtr.echoUserInfo()
	newUserPtr.changeUserName("howard")
	newUserPtr.echoUserInfo()
}
