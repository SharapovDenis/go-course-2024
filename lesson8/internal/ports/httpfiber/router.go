package httpfiber

import (
	"github.com/gofiber/fiber/v2"

	"homework8/internal/app"
)

func AppRouter(r fiber.Router, app app.Service) {
	r.Get("/ads", listAd(app.Ad))
	r.Post("/ads", createAd(app.Ad, app.User))
	r.Put("/ads/:ad_id/status", changeAdStatus(app.Ad))
	r.Put("/ads/:ad_id", updateAd(app.Ad))

	r.Get("users", listUsers(app.User))
	r.Post("users", createUser(app.User))
	r.Put("/users/:user_id", updateUser(app.User))
}
