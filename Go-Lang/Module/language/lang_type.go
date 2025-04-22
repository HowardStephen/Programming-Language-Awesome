package language

type LangType string

const (
	GoLang     LangType = "Go"
	Java       LangType = "Java"
	Python     LangType = "Python"
	Swift      LangType = "Swift"
	CPlusPlus  LangType = "C++"
	Kotlin     LangType = "Kotlin"
	JavaScript LangType = "JavaScript"
	TypeScript LangType = "TypeScript"
	Lua        LangType = "Lua"
	Rust       LangType = "Rust"
)

type Lang struct {
	LangType LangType
}
