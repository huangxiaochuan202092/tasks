package handlers

import (
	"log"
	"net/http"
	"proapp/models"
	"proapp/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 创建任务
func CreateTask(c *gin.Context) {
	var input struct {
		Title       string    `json:"title" binding:"required"`
		Description string    `json:"description"`
		Priority    string    `json:"priority" binding:"oneof=low medium high"`
		DueDate     time.Time `json:"due_date" binding:"required"`
	}

	// 绑定 JSON 输入
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("JSON 绑定失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数无效"})
		return
	}

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Priority:    input.Priority,
		Status:      "todo",
		DueDate:     input.DueDate,
	}

	// 调用服务层创建任务
	if err := services.CreateTask(&task); err != nil {
		log.Printf("创建任务失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误，请稍后重试"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "任务创建成功",
		"task": gin.H{
			"id":          task.ID,
			"title":       task.Title,
			"description": task.Description,
			"priority":    task.Priority,
			"status":      task.Status,
			"due_date":    task.DueDate.Format(time.RFC3339),
		},
	})
}

// 获取所有任务
func GetAllTasks(c *gin.Context) {
	tasks, err := services.GetTasks()
	if err != nil {
		log.Printf("获取任务列表失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误，请稍后重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// 获取单个任务
func GetTask(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的任务ID"})
		return
	}

	task, err := services.GetTask(uint(idUint))
	if err != nil {
		log.Printf("获取任务失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误，请稍后重试"})
		return
	}

	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "任务不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

// 更新任务
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的任务ID"})
		return
	}

	var input struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Priority    string    `json:"priority" binding:"omitempty,oneof=low medium high"`
		DueDate     time.Time `json:"due_date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("JSON 绑定失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数无效"})
		return
	}

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Priority:    input.Priority,
		DueDate:     input.DueDate,
	}

	if err := services.UpdateTask(uint(idUint), &task); err != nil {
		log.Printf("更新任务失败: %v", err)
		if err.Error() == "任务不存在" {
			c.JSON(http.StatusNotFound, gin.H{"error": "任务不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误，请稍后重试"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "任务更新成功"})
}

// 删除任务
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的任务ID"})
		return
	}

	if err := services.DeleteTask(uint(idUint)); err != nil {
		log.Printf("删除任务失败: %v", err)
		if err.Error() == "任务不存在" {
			c.JSON(http.StatusNotFound, gin.H{"error": "任务不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误，请稍后重试"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "任务删除成功"})
}
