package main

import (
	"VAtcoder/controllers"
	"VAtcoder/db"
	"VAtcoder/routers"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func insertProblems() {
	for i := 220; i <= 365; i++ {
		for c := 'c'; c <= 'g'; c++ {
			data := db.Problem{
				//abc365_d
				Atcoderid: "abc" + fmt.Sprintf("%03d", i) + "_" + string(c),
				Level:     1,
			}
			db.DB.Create(&data)
		}
	}

	for i := 220; i <= 365; i++ {
		for c := 'd'; c <= 'g'; c++ {
			data := db.Problem{
				//abc365_d
				Atcoderid: "abc" + fmt.Sprintf("%03d", i) + "_" + string(c),
				Level:     2,
			}
			db.DB.Create(&data)
		}
	}

	for i := 220; i <= 365; i++ {
		for c := 'c'; c <= 'e'; c++ {
			data := db.Problem{
				//abc365_d
				Atcoderid: "abc" + fmt.Sprintf("%03d", i) + "_" + string(c),
				Level:     3,
			}
			db.DB.Create(&data)
		}
	}

}

func insertUsers() {
	user := db.User{
		Level:    2,
		Truename: "付宇宸",
		Username: "fuyuchen",
		Password: "fuyuchen",
	}
	db.DB.Create(&user)

	user = db.User{
		Level:    2,
		Truename: "刘浩然",
		Username: "liuhaoran",
		Password: "liuhaoran",
	}
	db.DB.Create(&user)
	//赵彦哲、李秉泽、徐瑞泽、房泓宇、云晨轩、李宸朗、王奕山、颜庭泽、杜晨阳

	user = db.User{
		Level:    2,
		Truename: "赵彦哲",
		Username: "zhaoyanzhe",
		Password: "zhaoyanzhe",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    2,
		Truename: "李秉泽",
		Username: "libingze",
		Password: "libingze",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    2,
		Truename: "徐瑞泽",
		Username: "xuruize",
		Password: "xuruize",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    2,
		Truename: "房泓宇",
		Username: "fanghongyu",
		Password: "fanghongyu",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    2,
		Truename: "云晨轩",
		Username: "yunchenxuan",
		Password: "yunchenxuan",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    2,
		Truename: "李宸朗",
		Username: "lichenlang",
		Password: "lichenlang",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    2,
		Truename: "王奕山",
		Username: "wangyishan",
		Password: "wangyishan",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    2,
		Truename: "颜庭泽",
		Username: "yantingze",
		Password: "yantingze",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    2,
		Truename: "杜晨阳",
		Username: "duchenyang",
		Password: "duchenyang",
	}
	db.DB.Create(&user)
	//3:
	//胡赫轩、杨洋、穆鹏宇、石宇爀、于珈浩、程晟泰、梁钰涵、孙乐涵
	//刘佳赫、窦浩轩、刘子淇、左子毅、王景睿、周子航、刘浩、蔡云翔
	//冯嘉尧、催嘉睿、王龙晔、王彤睿、李荣桓、周亦帆
	user = db.User{
		Level:    3,
		Truename: "胡赫轩",
		Username: "huhexu",
		Password: "huhexu",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "杨洋",
		Username: "yangyang",
		Password: "yangyang",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "穆鹏宇",
		Username: "mupengyu",
		Password: "mupengyu",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "石宇爀",
		Username: "shiyiran",
		Password: "shiyiran",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "于珈浩",
		Username: "yujiahao",
		Password: "yujiahao",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "程晟泰",
		Username: "chenshengtai",
		Password: "chenshengtai",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "梁钰涵",
		Username: "liangyuhan",
		Password: "liangyuhan",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "孙乐涵",
		Username: "sunlehan",
		Password: "sunlehan",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "刘佳赫",
		Username: "liujiahe",
		Password: "liujiahe",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "窦浩轩",
		Username: "douhaoxuan",
		Password: "douhaoxuan",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "刘子淇",
		Username: "liuziqi",
		Password: "liuziqi",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "左子毅",
		Username: "zuoziyi",
		Password: "zuoziyi",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "王景睿",
		Username: "wangjingrui",
		Password: "wangjingrui",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "周子航",
		Username: "zhouzihang",
		Password: "zhouzihang",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "刘浩",
		Username: "liuhao",
		Password: "liuhao",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "蔡云翔",
		Username: "caiyunxiang",
		Password: "caiyunxiang",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "冯嘉尧",
		Username: "fengjiayao",
		Password: "fengjiayao",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "催嘉睿",
		Username: "cuijiarui",
		Password: "cuijiarui",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "王龙晔",
		Username: "wanglongye",
		Password: "wanglongye",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "王彤睿",
		Username: "wangtongrui",
		Password: "wangtongrui",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "李荣桓",
		Username: "lironghuan",
		Password: "lironghuan",
	}
	db.DB.Create(&user)
	user = db.User{
		Level:    3,
		Truename: "周亦帆",
		Username: "zhouyifan",
		Password: "zhouyifan",
	}
	db.DB.Create(&user)

}

func startRouter() {
	fmt.Println("Starting router...")
	err := routers.Router().Run(":3000")
	if err != nil {
		log.Fatalf("Failed to start router: %v", err)
	}
}
func startProxy() {
	targetURL, err := url.Parse("https://" + controllers.ApiHost + "/")
	if err != nil {
		log.Fatal(err)
	}
	proxyHandler := controllers.ReverseProxyHandler(targetURL)
	http.HandleFunc("/", proxyHandler)
	log.Println("Starting reverse proxy server")
	err = http.ListenAndServe(":3090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func yuanshen() {
	//并行启动startRouter startProxy
	go startRouter()
	go startProxy()
	//每隔一小时执行一次AutoLogin：
	go func() {
		f := true
		for {
			log.Println("Auto Login...........................")
			controllers.AutoLogin(f)
			f = !f
			time.Sleep(time.Hour)
		}
	}()
}

func main() {
	//insertProblems()
	//insertUsers()
	//yuanshen()
	startRouter()
	//log.Println("Error exit")
	//select {}

	//for i := 220; i <= 365; i++ {
	//	for c := 'g'; c <= 'g'; c++ {
	//		data := db.Problem{
	//			//abc365_d
	//			Atcoderid: "abc" + fmt.Sprintf("%03d", i) + "_" + string(c),
	//			Level:     1,
	//		}
	//		db.DB.Create(&data)
	//	}
	//}
	//
	//for i := 220; i <= 365; i++ {
	//	for c := 'g'; c <= 'g'; c++ {
	//		data := db.Problem{
	//			//abc365_d
	//			Atcoderid: "abc" + fmt.Sprintf("%03d", i) + "_" + string(c),
	//			Level:     2,
	//		}
	//		db.DB.Create(&data)
	//	}
	//}

}
