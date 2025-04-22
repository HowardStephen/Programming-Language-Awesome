package main

import (
	"henrygomodule.com/test/peopletest"
	"usertest/user"
)

func main() {
	newUser := user.User{}
	newUser.DefineUser("henry", "2100770220")
	newUser.ShowUserInfo()

	newPeople := people.People{}
	newPeople.DefinePeople("howard", 22)
	newPeople.ShowPeopleInfo()
}
