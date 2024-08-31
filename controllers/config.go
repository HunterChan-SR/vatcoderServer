package controllers

import (
	"VAtcoder/db"
	"VAtcoder/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const OK = 0
const ERROR = 1

const apiURL = "https://vjudge.net"

//const apiURL = "http://localhost:8080"

const username = "team_014"
const password = "myrbwgdwa"

var jSESSIONID = "DDD8D0D79F51EA9AEEAC00BAC3AE8ECA"
var jSESSlONID = "3OF0CRL65CCNUAGTFK2QE54MM0JJDD4T"
var jax_Q = ""

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"cnt"`
}

type JsonErrStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int) {
	json := &JsonStruct{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JsonErrStruct{
		Code: code,
		Msg:  msg,
	}
	c.JSON(200, json)
}

func vjudgeApi(targetURL string, method string, data map[string][]string) []byte {
	log.Println("======>vjudgeApi<======")
	log.Println("REQUEST")
	log.Println("targetURL: " + targetURL)
	log.Println("method: " + method)
	payloadData := url.Values{}
	for k, v := range data {
		payloadData[k] = v
	}
	log.Println("data:", payloadData)
	payload := strings.NewReader(payloadData.Encode())
	client := &http.Client{}
	req, err := http.NewRequest(method, targetURL, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if !strings.Contains(targetURL, "login") {
		reqCookie := ""
		if jSESSIONID != "" {
			reqCookie += "JSESSIONID=" + jSESSIONID + "; "
		}
		if jSESSlONID != "" {
			reqCookie += "JSESSlONID=" + jSESSlONID + "; "
		}
		if jax_Q != "" {
			reqCookie += "Jax.Q=" + jax_Q + "; "
		}
		req.Header.Add("Cookie", reqCookie)
		log.Println("Cookie: " + reqCookie)
	}
	resp, err := client.Do(req)
	log.Println("RESPONSE")
	log.Println("status: " + resp.Status)
	log.Println("cookies: ")
	for _, v := range resp.Cookies() {
		if v.Name == "JSESSIONID" {
			jSESSIONID = v.Value
			log.Println("JSESSIONID: " + jSESSIONID)
		} else if v.Name == "JSESSlONID" {
			jSESSlONID = v.Value
			log.Println("JSESSlONID: " + jSESSlONID)
		} else if v.Name == "Jax.Q" {
			jax_Q = v.Value
			log.Println("Jax.Q: " + jax_Q)
		}
	}

	all, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("vjudgeApi error: " + err.Error())
		}
	}(resp.Body)

	if err != nil {
		log.Println("vjudgeApi error: " + err.Error())
		return nil
	}
	log.Println("Body:", string(all))
	return all
}

func AutoLogin(f bool) {
	if f {
		data := map[string][]string{
			//"username": username,
			//"password": password,
			//"captcha":  "",
			"username": {"admin014"},
			"password": {"baogongtou"},
			"captcha":  {""},
		}
		vjudgeApi(apiURL+"/user/login", "POST", data)
	} else {
		data := map[string][]string{
			//"username": username,
			//"password": password,
			//"captcha":  "",
			"username": {"admin013"},
			"password": {"baogongtou"},
			"captcha":  {""},
		}
		vjudgeApi(apiURL+"/user/login", "POST", data)
	}
	log.Println("JSESSIONID:", jSESSIONID)
	log.Println("JSESSlONID:", jSESSlONID)
	log.Println("Jax.Q:", jax_Q)
}

func VerifyVjudgeOnline() bool {
	res := string(vjudgeApi(apiURL+"/user/checkLogInStatus", "POST", nil))
	if strings.Contains(res, "true") || strings.Contains(res, "success") {
		return true
	} else {
		return false
	}
}

func GetVjudgeCaptcha() {
	// 返回jpeg
	imgBytes := vjudgeApi(apiURL+"/util/captcha?1722858493282.5", "GET", nil)
	err := utils.SaveImageToFile(imgBytes, "captcha.jpg")
	if err != nil {
		log.Println("saveImageToFile error: " + err.Error())
	}
}

func LoginVjudge(captcha string) bool {
	data := map[string][]string{
		"username": {username},
		"password": {password},
		"captcha":  {captcha},
	}
	res := string(vjudgeApi(apiURL+"/user/login", "POST", data))
	if strings.Contains(res, "true") || strings.Contains(res, "success") {
		return true
	} else {
		return false
	}

}

func Pending() error {
	for true {
		fmt.Println("New Round Pending")
		var pendingId []int
		db.DB.Model(&db.Submit{}).Where("statuscanonical = ?", "PENDING").Pluck("id", &pendingId)
		for _, id := range pendingId {
			resBytes := vjudgeApi(apiURL+"/solution/data/"+strconv.Itoa(id), "POST", nil)
			var res map[string]interface{}
			err := json.Unmarshal(resBytes, &res)
			if err != nil {
				return err
			}
			if res["statusCanonical"] != "PENDING" {
				db.DB.Model(&db.Submit{}).Where("id = ?", id).Update("statuscanonical", res["statusCanonical"])
			}
			time.Sleep(time.Second * 10)
		}
		time.Sleep(time.Second * 300)
		fmt.Println("Pending done")
	}
	return nil
}

const ApiHost = "vjudge.net"

//const proxyHost = "liamaking.eu.org"

func ReverseProxyHandler(target *url.URL) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Cookies()
		//找到cookie中的token
		token := ""
		for _, cookie := range cookies {
			if cookie.Name == "token" {
				token = cookie.Value
				break
			}
		}
		log.Println("token:", token)
		if CheckToken(token) != 1 {
			body := "你是怎么找到这里的?~!"
			w.Write([]byte(body))
			return
		}
		// 日志记录表单提交
		log.Println(r.Method, r.URL.String())
		proxyHost := r.URL.Hostname()
		log.Println("Cookies:")
		for _, v := range cookies {
			if v.Name == "JSESSIONID" {
				jSESSIONID = v.Value
				log.Println("JSESSIONID: " + jSESSIONID)
			} else if v.Name == "JSESSlONID" {
				jSESSlONID = v.Value
				log.Println("JSESSlONID: " + jSESSlONID)
			} else if v.Name == "Jax.Q" {
				jax_Q = v.Value
				log.Println("Jax.Q: " + jax_Q)
			}
		}
		// 复制请求体以保证后续使用
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		r.Body = io.NopCloser(strings.NewReader(string(body)))

		// 解析表单数据
		log.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
		err = r.ParseForm()
		if err != nil {
			log.Println("ParseForm error:", err)
		}
		log.Println(r.Form)
		//log.Println(r.PostForm)
		for key, values := range r.PostForm {
			log.Println(key, values)
		}
		log.Println("############################")
		//encodedString := r.Form["source"]
		//log.Println("encodedString:", encodedString)
		//decodedBytes, err := base64.StdEncoding.DecodeString(encodedString[0])
		//if err != nil {
		//	log.Println("Decode error:", err)
		//}
		//log.Println(string(decodedBytes))

		// 创建一个自定义的 RoundTripper
		tr := &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}

		// 使用自定义的 RoundTripper
		proxy := &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				req.URL.Scheme = target.Scheme
				req.URL.Host = target.Host
				req.Body = io.NopCloser(strings.NewReader(string(body)))
			},
			Transport: tr,
			// 自定义响应处理器
			ModifyResponse: func(resp *http.Response) error {
				for i, h := range resp.Header["Set-Cookie"] {
					resp.Header["Set-Cookie"][i] = strings.Replace(h, ApiHost, proxyHost, -1)
					log.Println(resp.Header["Set-Cookie"][i])
				}
				return nil
			},
		}
		proxy.ServeHTTP(w, r)
	}
}
