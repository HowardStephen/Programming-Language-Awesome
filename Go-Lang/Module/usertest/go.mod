module usertest

go 1.23.2

require (
	"henrygomodule.com/test/peopletest" v0.0.0
)

replace (
	"henrygomodule.com/test/peopletest" => "../peopletest"
)