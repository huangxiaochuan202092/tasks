package handlers

import (
	"net/http"
	"proapp/models"
	"proapp/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 创建问卷
// 创建问卷
func CreateWenjuan(c *gin.Context) {
	var input struct {
		Title    string     `json:"title" binding:"required"`
		Content  string     `json:"content" binding:"required"`
		Status   string     `json:"status" binding:"required"`
		Deadline *time.Time `json:"deadline"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数无效", "data": err.Error()})
		return
	}

	wenjuan := models.Wenjuan{
		Title:    input.Title,
		Content:  input.Content,
		Status:   input.Status,
		Deadline: input.Deadline,
	}

	err := services.CreateWenjuan(&wenjuan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "创建问卷失败", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "问卷创建成功", "data": wenjuan})
}

// 获取所有问卷
func GetAllWenjuans(c *gin.Context) {
	wenjuans, err := services.GetAllWenjuans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wenjuans)
}

// 获取单个问卷
func GetWenjuanById(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	wenjuan, err := services.GetWenjuanById(int(idUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wenjuan)
}

// 更新问卷
func UpdateWenjuan(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	var input struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		Status   string `json:"status"`
		Deadline string `json:"deadline"` // 前端传递的是字符串
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if input.Title != "" {
		updates["title"] = input.Title
	}
	if input.Content != "" {
		updates["content"] = input.Content
	}
	if input.Status != "" {
		updates["status"] = input.Status
	}
	if input.Deadline != "" {
		deadline, err := time.Parse(time.RFC3339, input.Deadline)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的截止时间格式"})
			return
		}
		updates["deadline"] = deadline
	}

	if err := services.UpdateWenjuan(int(idUint), updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "问卷更新成功"})
}

// 删除问卷
func DeleteWenjuan(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := services.DeleteWenjuan(int(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "问卷删除成功"})
}

// 提交问卷答案
func SubmitWenjuanAnswer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的问卷ID",
			"error":   err.Error(),
		})
		return
	}

	var input struct {
		Answer string `json:"answer" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	if err := services.SubmitWenjuanAnswer(id, input.Answer); err != nil {
		if err.Error() == "问卷不存在" {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "问卷不存在",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "提交答案失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "答案提交成功",
	})
}
