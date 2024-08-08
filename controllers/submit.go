package controllers

import (
	"VAtcoder/db"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type SubmitController struct{}

//
//type SubmitStruct struct {
//	Captcha  string `json:"captcha"`
//	Oj       string `json:"oj"`
//	ProbNum  string `json:"probNum"`
//	Method   string `json:"method"`
//	Language string `json:"language"`
//	Open     string `json:"open"`
//	Source   string `json:"source"`
//}

func (sc SubmitController) Submit(c *gin.Context) {
	//获取post发来的请求
	//解析x-www-form-urlencoded格式的请求

	userid := CheckAuthId(c)
	if userid == 0 {
		ReturnError(c, ERROR, "请先登录")
		return
	}

	type codeSubmit struct {
		Source    string `json:"source"`
		Atcoderid string `json:"atcoderid"`
	}
	var codesubmit codeSubmit
	err := c.Bind(&codesubmit)

	data := map[string][]string{
		"probNum":  []string{codesubmit.Atcoderid},
		"source":   []string{codesubmit.Source},
		"captcha":  []string{""},
		"oj":       []string{"AtCoder"},
		"method":   []string{"0"},
		"language": []string{"5001"},
		"open":     []string{"0"},
	}

	resBytes := vjudgeApi(apiURL+"/problem/submit", "POST", data)
	//resBytes = {"runId":53251234} //截取提交id
	resStr := string(resBytes)
	if !strings.Contains(resStr, "runId") {
		ReturnError(c, ERROR, "Failed to submit code")
		return
	}
	leftindex := strings.Index(resStr, ":")
	rightindex := strings.Index(resStr, "}")
	submitid, err := strconv.Atoi(resStr[leftindex+1 : rightindex])
	if err != nil || resBytes == nil {
		ReturnError(c, ERROR, "Failed to submit code")
		return
	}
	submit := db.Submit{
		Id:              submitid,
		Userid:          userid,
		Atcoderid:       data["probNum"][0],
		Statuscanonical: "PENDING",
	}

	db.DB.Create(submit)
	ReturnSuccess(c, OK, "提交成功", submitid, 1)
}

func (sc SubmitController) Data(c *gin.Context) {

	userid := CheckAuthId(c)
	if userid == 0 {
		ReturnError(c, ERROR, "请先登录")
		return
	}

	id := c.Param("id")

	for i := 1; i <= 60; i++ {
		resBytes := vjudgeApi(apiURL+"/solution/data/"+id, "POST", nil)
		//使用map解析resBytes，resBytes为json格式
		var res map[string]interface{}
		err := json.Unmarshal(resBytes, &res)
		if err != nil {
			ReturnError(c, ERROR, "Failed to get data")
			return
		}
		if res["statusCanonical"] != "PENDING" {
			db.DB.Model(db.Submit{}).Where("id = ?", id).Update("statuscanonical", res["statusCanonical"])
			ReturnSuccess(c, OK, "success", res, 1)
			return
		}
		time.Sleep(time.Second * 2)
	}
	ReturnError(c, ERROR, "Failed to get data")
}
