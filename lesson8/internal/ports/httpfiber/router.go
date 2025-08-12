package httpfiber

import (
	"github.com/gofiber/fiber/v2"

	"homework8/internal/service"
)

func AppRouter(r fiber.Router, s service.Service) {
	r.Get("/ads", listAd(s))
	r.Post("/ads", createAd(s))
	r.Put("/ads/:ad_id/status", changeAdStatus(s))
	r.Put("/ads/:ad_id", updateAd(s))

	r.Get("users", listUsers(s))
	r.Post("users", createUser(s))
	r.Put("/users/:user_id", updateUser(s))
}
