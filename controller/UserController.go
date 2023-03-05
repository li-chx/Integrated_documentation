package controller

import (
	"fmt"
	"github.com/Haroxa/Integrated_documentation/common"
	"github.com/Haroxa/Integrated_documentation/helper"
	"github.com/Haroxa/Integrated_documentation/model"
	"github.com/Haroxa/Integrated_documentation/service"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"net/http"
)

func Mail(c *gin.Context) {
	err := service.SendMail("2379008409@qq.com", 1, "\ntest000\n", "\ntest111\n")
	if err != nil { //   读取  用户 并判断
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(-1, "发送失败", err))
		return
	}
	c.JSON(http.StatusBadRequest, helper.ApiReturn(1, "发送成功", nil))
}

func Register(c *gin.Context) {
	userLogin := &model.UserLogin{}                     //  初始化  用户  模型
	if err := c.ShouldBindJSON(userLogin); err != nil { //   读取  用户 并判断
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}
	if userLogin.Email == "" || userLogin.Password == "" {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "邮箱或密码不能为空", nil))
		return
	}
	User, _ := model.FindUser(userLogin.Email)
	if User.Email != "" {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "用户已存在", nil))
		return
	}
	//User.Id = int(time.Now().Unix() - 1677922335)
	User.Email = userLogin.Email
	User.Password = userLogin.Password
	if err := model.CreateUser(&User); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "创建用户失败", err))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "注册成功，请前往登录", User))
	c.Redirect(200, "/user/login")
}

func Login(c *gin.Context) {
	//fmt.Println("OK,start to login")
	userLogin := &model.UserLogin{}                     //  初始化  用户  模型
	if err := c.ShouldBindJSON(userLogin); err != nil { //   读取  用户 并判断
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}

	User, _ := model.FindUser(userLogin.Email)
	if User.Email == "" {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "用户不存在", nil))
		return
	}
	if User.Password != userLogin.Password {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "密码错误", nil))
		return
	}

	token, err := helper.CreatToken(User.Id)
	//fmt.Println("3\n%+v", token)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "创建token失败", err))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "登录成功", token))
}

func UpdateUser(c *gin.Context) {
	UserId := c.MustGet("user_id").(int)
	user, err := model.GetUserById(UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取用户失败", err))
		return
	}
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "用户不存在", nil))
		return
	}
	u := &model.User{}
	//fmt.Println("\n\n\n", user)
	if err = c.ShouldBindJSON(u); err != nil { //   读取  用户 并判断
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}

	//data, _ := json.Marshal(u)
	//mp := make(map[string]interface{})
	//json.Unmarshal(data, &mp)
	mp := structs.Map(u)
	for k, v := range mp {
		if v == "" || v == 0 {
			delete(mp, k)
		} else {
			//fmt.Printf("[ %v : %v ]\n", k, v)
		}
	}

	if err = model.UpdateUser(user, mp); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "信息更新失败", err))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "信息更新成功", nil))
}

func GetUserById(c *gin.Context) {
	//获取用户
	UserId := c.MustGet("user_id").(int)
	user, err := model.GetUserById(UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取用户失败", err))
		return
	}
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "用户不存在", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "获取成功", user))
}

func GetAllUser(c *gin.Context) {
	//获取所有用户切片和用户数
	users, count, err := model.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取用户失败", err))
		return
	}
	msg := fmt.Sprintf("获取成功，得到用户数为：%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, users))
}

func DeleteUser(c *gin.Context) {
	UserId := c.MustGet("user_id").(int)
	var user model.User
	var err error
	if user, err = model.GetUserById(UserId); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取用户失败", err))
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "用户不存在", nil))
		return
	}

	if err = model.DeleteUser(user); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "删除用户失败", err))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "删除用户成功", nil))
}
