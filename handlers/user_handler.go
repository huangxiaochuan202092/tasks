package handlers

import (
	"log"
	"net/http"
	"proapp/services"
	"proapp/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 发送验证码
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

// 登录或注册
func LoginOrRegisterHandler(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	// 绑定 JSON 输入
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("绑定 JSON 失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取存储的验证码
	storedCode, err := utils.GetVerificationCode(input.Email)
	if err != nil {
		log.Printf("获取验证码失败: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "验证码获取失败"})
		return
	}

	// 验证码匹配
	if storedCode != input.Code {
		log.Printf("验证码不匹配: 输入验证码=%s, 存储验证码=%s", input.Code, storedCode)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "验证码不匹配"})
		return
	}

	// 检查用户是否存在
	user, err := services.GetUserByEmail(input.Email)
	if err != nil {
		log.Printf("获取用户失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库错误"})
		return
	}

	log.Printf("获取用户结果: %+v", user)

	// 如果用户不存在，则创建用户
	if user == nil {
		log.Printf("用户不存在，尝试创建新用户: email=%s", input.Email)
		user, err = services.CreateUser(input.Email)
		if err != nil {
			log.Printf("创建用户失败: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
			return
		}
	}

	// 删除验证码
	if err := utils.DelVerificationCode(input.Email); err != nil {
		log.Printf("删除验证码失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "验证码删除失败"})
		return
	}

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"user": gin.H{
			"email": user.Email,
			"id":    user.ID,
		},
	})
}

// 获取所有用户
func GetAllUsersHandler(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// 根据id获取用户
func GetUserByIdHandler(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	user, err := services.GetUserById(uint(idUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// 更新用户
func UpdateUserHandler(c *gin.Context) {
	// 获取路径参数 id
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	// 绑定 JSON 输入
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层更新用户
	if err := services.UpdateUser(uint(idUint), input.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新用户成功"})
}

// 删除用户
func DeleteUserHandler(c *gin.Context) {
	// 获取路径参数 id
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	// 调用服务层删除用户
	if err := services.DeleteUserByID(uint(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除用户成功"})
}
