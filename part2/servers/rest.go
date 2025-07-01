package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestEngine struct {
	engine *gin.Engine
}

func NewRestEngine() *RestEngine {
	r := gin.New()

	return &RestEngine{
		engine: r,
	}
}

func (re *RestEngine) Setup() {
	re.engine.Use(gin.Recovery())

	re.engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "App is healthy",
		})
	})

	api := re.engine.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello John Doe",
			})
		})
	}
}

func (re *RestEngine) Run(addr string) error {
	return re.engine.Run(addr)
}

func (re *RestEngine) RunTLS(addr, certPath, keyPath string) error {
	return re.engine.RunTLS(addr, certPath, keyPath)
}
