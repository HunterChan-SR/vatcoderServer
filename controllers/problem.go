package controllers

import (
	"VAtcoder/db"
	"github.com/gin-gonic/gin"
)

type ProblemController struct{}

func (pc ProblemController) Get(c *gin.Context) {
	userid := CheckAuthId(c)
	if userid == 0 {
		ReturnError(c, ERROR, "请先登录")
		return
	} else {
		//查询problem_view表中的满足要求的atcoder列
		var data []db.ProblemCountView
		db.DB.Table("problem_count_view").Where("userid = ?", userid).Find(&data)
		ReturnSuccess(c, OK, "查询成功", data, len(data))
	}
}
