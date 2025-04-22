package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type hello struct {
	app.Compo
	Name string
}

func (h *hello) Render() app.UI {
	return app.H1().Text("Hello, " + h.Name)
}
