package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"try/models"
	"try/repository"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var newTask models.CreateTask

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	task, err := repository.CreateTask(newTask.Title, newTask.Description, newTask.Priority)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(http.StatusCreated, task)

}

func GetTasks(c *gin.Context) {

	tasks, err := repository.GetTasks()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}
	c.JSON(200, tasks)

}

func GetTaskById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid json",
		})
		return
	}

	task, err := repository.GetTaskByID(id)
	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{
			"error": "not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(200, task)

}

func UpdateTask(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id is invalid",
		})
		return
	}
	var newTask models.UpdateTask
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid json",
		})
		return
	}

	affected, err := repository.UpdateTask(id, newTask)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	if affected == 0 {
		c.JSON(404, gin.H{
			"error": "NOT FOUND",
		})
		return
	}

	c.Status(200)
}

func PatchTask(c *gin.Context) {

	var patT models.PatchTask

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	if err := c.ShouldBindJSON(&patT); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid json",
		})
		return
	}

	err = repository.PatchTask(id, patT)
	if err == repository.ErrNothingToUpdate {
		c.JSON(400, gin.H{
			"error": "nothing to update",
		})
		return
	}

	if err == repository.Notfound {
		c.JSON(404, gin.H{
			"error": "not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	c.Status(200)

}

func DeleteTask(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	affected, err := repository.DeleteTask(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	if affected == 0 {
		c.JSON(404, gin.H{
			"error": "not found",
		})
		return
	}

	c.Status(204)

}
