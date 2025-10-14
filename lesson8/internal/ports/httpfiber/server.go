package httpfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"homework8/internal/app"
)

type Server struct {
	port   string
	engine *fiber.App
}

func NewHTTPServer(port string, app app.Service) Server {
	server := Server{port: port, engine: fiber.New()}
	api := server.engine.Group("/api/v1")
	AppRouter(api, app)
	return server
}

func (s *Server) Listen() error {
	return s.engine.Listen(s.port)
}

func (s *Server) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return s.engine.Test(req, msTimeout...)
}
