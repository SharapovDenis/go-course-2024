package main

import (
	"homework8/internal/app"
	"homework8/internal/ports/httpfiber"
)

func main() {
	svc := app.New()
	r := httpfiber.NewHTTPServer(":8080", svc)
	r.Listen()
}
