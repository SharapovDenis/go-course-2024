package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"homework8/internal/service"
)

type Server struct {
	port   string
	engine *gin.Engine
}

func NewHTTPServer(port string, svc service.Service) Server {
	gin.SetMode(gin.ReleaseMode)
	server := Server{port: port, engine: gin.New()}

	server.engine.Use(gin.Logger(), gin.Recovery())

	adAPI := server.engine.Group("/api/v1")
	adAPI.GET("ads", listAd(svc)) // queryParams:[text, title, author_id, published, created_date]
	adAPI.POST("ads", createAd(svc))
	adAPI.PUT("/ads/:ad_id/status", changeAdStatus(svc))
	adAPI.PUT("/ads/:ad_id", updateAd(svc))

	userAPI := server.engine.Group("/api/v1")
	userAPI.GET("users", listUsers(svc))
	userAPI.POST("users", createUser(svc))
	userAPI.PUT("/users/:user_id", updateUser(svc))

	return server
}

func (s *Server) Listen() error {
	return s.engine.Run(s.port)
}

func (s *Server) Handler() http.Handler {
	return s.engine
}
