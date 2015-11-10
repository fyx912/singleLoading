package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"singleLoading/models"
	"singleLoading/utils"
	"singleLoading/utils/session"
)

var user map[string]interface{}
var globalSessions *session.Manager

//然后在init函数中初始化
func init() {
	//使用内存保存session,默认值是 3600 秒
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go globalSessions.GC()
}

// http post 登陆接口获取token
func Login(rw http.ResponseWriter, req *http.Request) {
	str := ""
	sesions, err := globalSessions.SessionStart(rw, req) //根据当前请求返回 session 对象
	models.CheckError(err)
	defer sesions.SessionRelease(rw)

	fmt.Println("request PostForm:   ", req.PostForm)
	fmt.Println("request Header :", req.Header)
	fmt.Println("request Header  Content-Type  :", req.Header.Get("Content-Type"))
	fmt.Println("request host :", req.Host)
	fmt.Println("request Form :", req.Form)
	if req.PostForm == nil {
		fmt.Println("requset 获取的form为: 空")
	} else {
		// username = req.FormValue("username")
		// password = req.FormValue("password")
	}
	if GetPostType(req) {
		//result := utils.GetDataString(req)
		body, err := ioutil.ReadAll(req.Body)
		models.CheckError(err)
		defer req.Body.Close()

		//fmt.Println("request req.Body string:", result)
		// fmt.Println("request req.Body string:", body)//打印字节:[123 34 117 115 101 114 110 97 109 101 34 58 34 115 121 115 116 101 109 34 44 34 112 97 115 115 119 111 114 100 34 58 34 49 50 51 52 53 54 34 125]

		err = json.Unmarshal(body, &user)
		models.CheckError(err)
		// fmt.Println("json", user)
		// fmt.Println("获取json中的username:", user["username"])
		//fmt.Println("获取json中的password:", user["password"].(string))
		// fmt.Println("username 接收的类型:", reflect.TypeOf(user["username"]))
		// fmt.Println("username 接收的类型:", reflect.TypeOf(user["username"]).Kind() == reflect.String)
		rw.Header().Add("Content-Type", "application/json")
		if reflect.TypeOf(user["username"]).Kind() == reflect.String {
			username := user["username"].(string)
			if len(username) != 0 {
				falg := models.ValidationUsername(username) //验证username
				if falg {
					if reflect.TypeOf(user["password"]).Kind() == reflect.String {
						password := user["password"].(string)
						if len(password) != 0 {
							// fmt.Println(" username:", username)
							// fmt.Println(" password:", password)
							falg := models.ValidationLogin(username, password)
							seesionTime := sesions.Get("time")
							if falg == true {
								if seesionTime == nil { //session time 等于空,则说明用户没有登录或登录已过期
									timeAgo, token := utils.GetToken() //timeAgo,token表示第一次请求的时间,token
									sesions.Set("username", username)
									sesions.Set("time", timeAgo)
									sesions.Set("password", password)
									sesions.Set("token", token)
								}
								fmt.Println("session ID=", sesions.SessionID())
								fmt.Println("session token=", sesions.Get("token"))
								str = "{\"code\":0,\"msg\":\"success\",\"token\":\"" + sesions.Get("token").(string) + "\"}"
							} else {
								str = utils.GetErrorJsonData(1, "密码错误!")
							}
						} else {
							str = utils.GetErrorJsonData(1, "密码不能为空!")
						}
					} else {
						str = utils.GetErrorJsonData(1, "输入的参数password必须为字符串!")
					}
				} else {
					str = utils.GetErrorJsonData(1, "用户名不存在 !")
				}
			} else {
				str = utils.GetErrorJsonData(1, "用户名不能为空!")
			}
		} else {
			str = utils.GetErrorJsonData(1, "输入的参数username必须为字符串!")
		}
	} else {
		str = utils.GetErrorJsonData(1, "请求的必须为POST方法!")
	}
	fmt.Println(str)
	rw.Write([]byte(str))
}

//退出接口
func Logout(w http.ResponseWriter, r *http.Request) {
	str := ""
	w.Header().Set("Content-Type", "application/json")
	if GetPostType(r) {
		data, err := ioutil.ReadAll(r.Body)
		models.CheckError(err)
		defer r.Body.Close()
		err = json.Unmarshal(data, &user)
		models.CheckError(err)
		token := user["token"]
		if reflect.TypeOf(token).Kind() == reflect.String { //判断用户输入的是否是string
			token := token.(string)
			if len(token) != 0 {
				falg := ValidationLogoutToken(w, r, token) //验证系统退出token
				if falg {
					str = utils.GetErrorJsonData(0, "系统退出成功!")
				} else {
					str = utils.GetErrorJsonData(1, "系统退出失败!")
				}
			} else {
				str = utils.GetErrorJsonData(1, "token不能为空!")
			}
		} else {
			str = utils.GetErrorJsonData(1, "输入的参数token必须为字符串!")
		}
	} else {
		str = utils.GetErrorJsonData(1, "请求的必须为POST方法!")
	}
	fmt.Println(str)
	w.Write([]byte(str))
}

//验证系统退出token
func ValidationLogoutToken(w http.ResponseWriter, r *http.Request, token string) bool {
	falg := false
	sees, err := globalSessions.SessionStart(w, r) //根据当前请求返回 session 对象
	models.CheckError(err)
	defer sees.SessionRelease(w)
	if token == sees.Get("token") { //验证成功,才能退出,并删除session中的token
		sees.Delete("time")
		sees.Delete("token")
		sees.Delete("username")
		sees.Delete("password")
		falg = true
	}
	return falg
}

//判断是否时post请求
func GetPostType(r *http.Request) bool {
	if r.Method == "POST" {
		return true
	} else {
		return false
	}
}

//判断是否时get请求
func GetGetType(r *http.Request) bool {
	if r.Method == "GET" {
		return true
	} else {
		return false
	}
}
