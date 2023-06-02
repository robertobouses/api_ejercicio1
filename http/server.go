package http

import "github.com/gin-gonic/gin"

type Server struct {
	engine  *gin.Engine
	handler Handler
}

type Handler struct {
}

func NewServer() Server {
	router := gin.Default()
	handler := Handler{}

	return Server{
		engine:  router,
		handler: handler,
	}
}
