package handler

import (
	"TodoPS/models"
	"TodoPS/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	repo *repository.Repository
}

func NewTodoHandler(repo *repository.Repository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.repo.TodoRepository.GetAll()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetByID(c *gin.Context) {
	Pid := c.Params.ByName("id")
	id, err := strconv.Atoi(Pid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	todo, err := h.repo.TodoRepository.GetOne(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, todo)
}

func (h *TodoHandler) Create(c *gin.Context) {
	var request models.Todo
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if request.Description == "" {
		c.JSON(400, gin.H{"Message": "Description is empty"})
		return
	}

	nowTime, err := GetTime()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	request.Created_at = nowTime
	request.Updated_at = nowTime

	if err := h.repo.TodoRepository.Create(request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Message": "Success create todo"})
}

func (h *TodoHandler) Update(c *gin.Context) {
	Pid := c.Params.ByName("id")
	id, err := strconv.Atoi(Pid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	todo, err := h.repo.TodoRepository.GetOne(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var json models.Todo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if json.Description == "" {
		c.JSON(400, gin.H{"error": "Description empty"})
		return
	}
	todo.Description = json.Description

	nowTime, err := GetTime()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	todo.Updated_at = nowTime

	if err := h.repo.TodoRepository.Update(*todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, todo)
}

func (h *TodoHandler) Delete(c *gin.Context) {
	Pid := c.Params.ByName("id")
	id, err := strconv.Atoi(Pid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	todo, err := h.repo.TodoRepository.GetOne(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.TodoRepository.Delete(*todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"Message": "Successful deletion of the todo",
	})
}
