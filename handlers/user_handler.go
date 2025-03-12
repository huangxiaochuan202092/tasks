package handlers

import (
	"net/http"
	"test/services"
	"test/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SendVerificationCode 发送验证码
func SendCodeHandler(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}

	// 绑定 JSON 输入
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 生成验证码
	code := utils.GenerateVerificationCode()
	// 发送验证码邮件
	if err := utils.SendVerificationEmail(input.Email, code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 存储验证码
	if err := utils.SetVerificationCode(input.Email, code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "验证码发送成功"})
}

// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	// 绑定 JSON 输入
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取存储的验证码
	storedCode, err := utils.GetVerificationCode(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 验证码匹配
	if storedCode != input.Code {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	// 注册用户
	if err := services.RegisterUser(input.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 删除验证码
	if err := utils.DelVerificationCode(input.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// LoginUser 登录用户
func LoginUser(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	// 绑定 JSON 输入
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取存储的验证码
	storedCode, err := utils.GetVerificationCode(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Verification code expired or not found"})
		return
	}

	// 验证码匹配
	if storedCode != input.Code {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	// 登录用户

	user, err := services.LoginUser(input.Email) // 接收正确返回值
	if err != nil {
		// 区分"用户不存在"和"数据库错误"
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "用户不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库错误"})
		}
		return
	}
	if err := utils.DelVerificationCode(input.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "user": user})
}
