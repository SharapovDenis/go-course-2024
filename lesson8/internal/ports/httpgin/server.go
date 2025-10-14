package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"homework8/internal/app"
)

type Server struct {
	port   string
	engine *gin.Engine
}

func NewHTTPServer(port string, app app.Service) Server {
	gin.SetMode(gin.ReleaseMode)
	server := Server{port: port, engine: gin.New()}

	server.engine.Use(gin.Logger(), gin.Recovery())

	adAPI := server.engine.Group("/api/v1")
	adAPI.GET("ads", listAd(app.Ad)) // queryParams:[text, title, author_id, published, created_date]
	adAPI.POST("ads", createAd(app.Ad, app.User))
	adAPI.PUT("/ads/:ad_id/status", changeAdStatus(app.Ad))
	adAPI.PUT("/ads/:ad_id", updateAd(app.Ad))

	userAPI := server.engine.Group("/api/v1")
	userAPI.GET("users", listUsers(app.User))
	userAPI.POST("users", createUser(app.User))
	userAPI.PUT("/users/:user_id", updateUser(app.User))

	return server
}

func (s *Server) Listen() error {
	return s.engine.Run(s.port)
}

func (s *Server) Handler() http.Handler {
	return s.engine
}
