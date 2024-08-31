package controllers

import (
	"VAtcoder/db"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type AdminController struct{}

func (ac AdminController) ReLoad(c *gin.Context) {
	if CheckAuthId(c) != 1 {
		//没有权限
		ReturnError(c, ERROR, "No Permission")
		return
	}
	data := map[string][]string{
		//"username": username,
		//"password": password,
		//"captcha":  "",
		"username": {"admin014"},
		"password": {"baogongtou"},
		"captcha":  {""},
	}
	bytes := vjudgeApi(apiURL+"/user/login", "POST", data)

	log.Println("JSESSIONID:", jSESSIONID)
	log.Println("JSESSlONID:", jSESSlONID)
	log.Println("Jax.Q:", jax_Q)

	ReturnSuccess(c, OK, "Success", string(bytes), 0)
}

func (ac AdminController) GetLevels(c *gin.Context) {
	if CheckAuthId(c) != 1 {
		//没有权限
		ReturnError(c, ERROR, "No Permission")
		return
	}
	var levels []db.Level
	db.DB.Find(&levels)
	ReturnSuccess(c, OK, "Success", levels, len(levels))
}

func (ac AdminController) GetAccount(c *gin.Context) {
	if CheckAuthId(c) != 1 {
		//没有权限
		ReturnError(c, ERROR, "No Permission")
		return
	}
	param := c.Param("level")
	level, err := strconv.Atoi(param)
	if err != nil {
		ReturnError(c, ERROR, "Error param")
		return
	}
	var data []db.ProblemCountView
	db.DB.Table("problem_count_view").Where("level = ?", level).Find(&data)
	ReturnSuccess(c, OK, "查询成功", data, len(data))
}

func (ac AdminController) AddProblem(c *gin.Context) {
	if CheckAuthId(c) != 1 {
		//没有权限
		ReturnError(c, ERROR, "No Permission")
		return
	}
	var problem db.Problem
	err := c.Bind(&problem)
	if err != nil || problem.Atcoderid == "" {
		ReturnError(c, ERROR, "Error param")
		return
	}
	if db.DB.Table("problem").Where("atcoderid = ? and level = ?", problem.Atcoderid, problem.Level).First(&db.Problem{}).RowsAffected != 0 {
		ReturnError(c, ERROR, "题目已存在")
		return
	}
	db.DB.Create(&problem)
	ReturnSuccess(c, OK, "Success", "Success", 0)
}

func (ac AdminController) AddUser(c *gin.Context) {
	if CheckAuthId(c) != 1 {
		//没有权限
		ReturnError(c, ERROR, "No Permission")
		return
	}
	var user db.User
	err := c.Bind(&user)
	if err != nil || user.Username == "" {
		ReturnError(c, ERROR, "Error param")
		return
	}
	if db.DB.Table("user").Where("username = ?", user.Username).First(&db.User{}).RowsAffected != 0 {
		ReturnError(c, ERROR, "用户名已存在")
		return
	}
	db.DB.Create(&user)
	ReturnSuccess(c, OK, "Success", "Success", 0)
}
