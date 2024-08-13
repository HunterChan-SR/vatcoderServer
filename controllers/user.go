package controllers

import (
	"VAtcoder/db"
	"VAtcoder/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type UserController struct{}

func CheckToken(token string) int {
	dtoken := utils.Decrypt(token)
	info := strings.Split(dtoken, ":")
	if len(info) != 2 {
		return 0
	}
	username := info[0]
	password := info[1]

	if len(username) == 0 || len(password) == 0 {
		return 0
	}
	var user db.User
	if db.DB.Where("username = ?", username).Where("password = ?", password).First(&user).RowsAffected == 0 {
		return 0
	}
	return user.Id
}

func CheckAuthId(c *gin.Context) int {
	//获取token
	token := c.GetHeader("Authorization")
	token1 := ""
	//http://localhost/******?token=9mrtzn;ojneb
	qToken := c.Request.URL.RawQuery
	if len(qToken) > 0 && strings.Contains(qToken, "token=") {
		token1 = strings.Split(qToken, "token=")[1]
	}

	if len(token) == 0 && len(token1) == 0 {
		return 0
	}
	if len(token) == 0 {
		token = token1
	}
	return CheckToken(token)
}

func (u UserController) PostLogin(c *gin.Context) {
	//获取表达 Username Password
	var user db.User
	_ = c.Bind(&user)
	log.Println(user)
	if len(user.Username) == 0 || len(user.Password) == 0 {
		ReturnError(c, ERROR, "用户名或密码不能为空")
		return
	}
	if db.DB.Where("username = ?", user.Username).Where("password = ?", user.Password).First(&user).RowsAffected == 0 {
		ReturnError(c, ERROR, "用户名或密码错误")
		return
	}
	token := strings.Join([]string{user.Username, user.Password}, ":")
	token = utils.Encrypt(token)
	c.SetCookie("token", token, 365*24*3600, "/", "", false, true)
	ReturnSuccess(c, OK, "success", token, 1)
}

func (u UserController) GetOnline(c *gin.Context) {
	id := CheckAuthId(c)
	if id == 0 {
		ReturnError(c, ERROR, "用户未登录")
		return
	} else {
		user := db.UserView{
			Id: id,
		}
		db.DB.Table("user_view").Where("id = ?", id).Find(&user)

		ReturnSuccess(c, OK, "success", user, 1)
	}
}

//type User struct {
//	Id       int    `json:"id"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//	Truename string `json:"truename"`
//	Level    int    `json:"level"`
//}

type NewPwd struct {
	NewPwd string `json:"newpwd"`
	OldPwd string `json:"oldpwd"`
}

func (u UserController) PostPwd(c *gin.Context) {
	id := CheckAuthId(c)
	if id == 0 {
		ReturnError(c, ERROR, "用户未登录")
		return
	}
	var newpwd NewPwd
	_ = c.Bind(&newpwd)
	if len(newpwd.NewPwd) == 0 || len(newpwd.OldPwd) == 0 || newpwd.NewPwd == newpwd.OldPwd {
		ReturnError(c, ERROR, "密码错误")
		return
	}
	var user db.User
	if db.DB.Where("id = ?", id).Where("password = ?", newpwd.OldPwd).First(&user).RowsAffected == 0 {
		ReturnError(c, ERROR, "旧密码错误")
		return
	}
	user.Password = newpwd.NewPwd
	db.DB.Model(&user).Update("password", newpwd.NewPwd)
	ReturnSuccess(c, OK, "success", "", 1)
}
