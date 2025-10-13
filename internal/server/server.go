package server

import (
	"fmt"

	"github.com/ItsMeSouptik/insightly/internal/config"
	"github.com/ItsMeSouptik/insightly/internal/handlers"
	"github.com/gin-gonic/gin"
)

func New(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		notes := v1.Group("/notes")
		{
			notes.GET("", handlers.ListNotes)
			notes.POST("", handlers.CreateNote)
			notes.GET("/:id", handlers.GetNote)
			notes.PUT("/:id", handlers.UpdateNote)
			notes.DELETE("/:id", handlers.DeleteNote)
		}
	}
	return r
}

func Run(cfg *config.Config) {
	r := New(cfg)
	addr := fmt.Sprintf(":%s", cfg.Port)
	r.Run(addr)
}
