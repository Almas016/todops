package handler

import (
	"TodoPS/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	TodoHandler *TodoHandler
	repo        *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{
		TodoHandler: NewTodoHandler(repo),
		repo:        repo,
	}
}

func GetTime() (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Almaty")
	if err != nil {
		return time.Time{}, err
	}
	now := time.Now().In(loc)
	return now, nil
}

func (h *Handler) Run() error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/success", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
		})
	})

	r.GET("//todos", h.TodoHandler.GetAllTodos)
	r.GET("/todo/:id", h.TodoHandler.GetByID)
	r.POST("/todo", h.TodoHandler.Create)
	r.PUT("/todo/:id", h.TodoHandler.Update)
	r.DELETE("/todo/:id", h.TodoHandler.Delete)
	return r.Run(":8080")
}
