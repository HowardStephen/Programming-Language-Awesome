package language

type goLang struct {
	Lang
	Template string
}

func (l *goLang) DefineTemplate() {
	l.Template = `
package main
		
import "fmt"
		
func main() {
	fmt.Println("Hello World")
}`
}

func (l *goLang) EchoLangInfo() {
	println("lang:\t", l.LangType)
	println("\ntemplate:\n", l.Template)
}

var MyGoLang Langer = &goLang{
	Lang: Lang{
		LangType: GoLang,
	},
	Template: "",
}
