package server

import (
	"fmt"
	"log"

	entities "github.com/DanielQuerolBeltran/Climbing-notebook-api/internal/platform"
	"github.com/DanielQuerolBeltran/Climbing-notebook-api/internal/platform/server/handler/climb"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr        string
	engine          *gin.Engine
	climbRepository entities.ClimbRepository
}

func New(host string, port uint, climbRepository entities.ClimbRepository) Server {
	srv := Server{
		httpAddr:        fmt.Sprintf(":%d", port),
		engine:          gin.New(),
		climbRepository: climbRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on port", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/check", climb.CheckHandler())
	s.engine.POST("/climb", climb.CreateHandler(s.climbRepository))
	s.engine.GET("/climb", climb.FetchHandler(s.climbRepository))
}
