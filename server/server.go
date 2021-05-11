package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New() (server *Server, err error) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	sever = Server{
		router: r,
	}
	return
}

func (s *Server) Start() {
	s.router.Run()
}

func (s *Server) Stop() {

}
