package controllers

import (
	"net/http"
	"strconv"
	"todo-gin/models"

	"github.com/gin-gonic/gin"
)

// GET /todos
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, models.Todos)
}

// POST /todos
func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo.ID = len(models.Todos) + 1
	models.Todos = append(models.Todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// PUT /todos/:id
func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, t := range models.Todos {
		if t.ID == id {
			updatedTodo.ID = id
			models.Todos[i] = updatedTodo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// DELETE /todos/:id
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, t := range models.Todos {
		if t.ID == id {
			models.Todos = append(models.Todos[:i], models.Todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
