package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"log"
	"net/http"
)

func main() {
	app.Route("/", func() app.Composer { return &hello{Name: "Henry"} })
	app.Route("/hello", func() app.Composer { return &hello{} })
	
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "hello",
		Description: "Hello World",
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
