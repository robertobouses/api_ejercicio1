package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APP interface {
}

type Server struct {
	engine  *gin.Engine
	handler Handler
}

type Handler struct {
}

type RegRectangle struct {
	Long  int `json: "Long"`
	Short int `json: "Short"`
}

func NewServer() Server {
	router := gin.Default()
	handler := Handler{}

	return Server{
		engine:  router,
		handler: handler,
	}
}

func (H Handler) InsertRectangle(c *gin.Context) {
	var rectangle RegRectangle
	if err := c.ShouldBindJSON(&rectangle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	area := rectangle.Area()

	// Aquí debes guardar el área en la base de datos

	c.JSON(http.StatusOK, gin.H{"area": area})

}

func (s *Server) Run(port string) error {
	router := gin.Default()
	router.POST("/insert", s.handler.InsertRectangle)
	return s.engine.Run(fmt.Sprintf(":%s", port))

}
