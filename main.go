package main

import (
	"VAtcoder/controllers"
	"VAtcoder/db"
	"VAtcoder/routers"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

func startRouter() {
	fmt.Println("Starting router...")
	err := routers.Router().Run(":80")
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
	log.Println("Starting reverse proxy server on :90")
	err = http.ListenAndServe(":90", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func yuanshen() {
	//并行启动startRouter startProxy
	go startRouter()
	go startProxy()
}

func main() {
	//insertProblems()
	yuanshen()
	log.Println("Error exit")
	select {}
}
