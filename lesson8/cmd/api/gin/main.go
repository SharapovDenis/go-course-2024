package main

import (
	"homework8/internal/app"
	"homework8/internal/ports/httpgin"
)

func main() {
	svc := app.New()
	r := httpgin.NewHTTPServer(":8080", svc)
	r.Listen()
}
