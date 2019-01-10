package controllers

import (
	"beego_blog/models"
	"beego_blog/util"
	"fmt"
	"strconv"
	"time"
)

type UserController struct {
	baseController
}

//配置信息
func (c *UserController) Login() {
	switch c.Ctx.Request.Method {
	case "GET":
		// c.loginGet()
		println("get")
	case "POST":
		c.loginPost()
		println("post")
	default:
		c.History("", "/")
		println("defualt")
	}

	c.TplName = c.controllerName + "/login.html"
}

//配置信息
func (c *UserController) Reg() {
	switch c.Ctx.Request.Method {
	case "GET":
		c.regGet()
		println("get")
	case "POST":
		c.regPost()
		println("post")
	default:
		c.History("", "/")
		println("defualt")
	}

}

func (c *UserController) regGet() {
	c.TplName = c.controllerName + "/reg.html"
}

func (c *UserController) regPost() {
	data := &models.RegForm{}
	err := c.ParseForm(data)
	if nil != err || data.Username == "" {
		str := data.Username
		if nil != err {
			str = err.Error()
		}
		c.msgString(-1, str)
		return
	}

	fmt.Println(data)
	if data.Pass == "" || data.Pass != data.Repass {
		c.msgString(-1, "密码不相同")
		return
	}

	user := c.getUser()
	if user.Username != "" {
		c.msgString(-1, "短时间内无法多次注册")
		return
	}

	user.Username = data.Username
	c.o.Read(user, "username")

	if user.Password != "" {
		c.msgString(-1, "alread registed.")
		return
	}

	user.Password = util.Md5(data.Pass)
	user.Name = data.Name
	user.Created = time.Now()

	_, err = c.o.Insert(user)
	if nil != err {
		c.msgString(-1, err.Error())
		return
	}

	c.msgString(0, "login")
}

func (c *UserController) msgString(code int, msg string) {
	c.Data["json"] = map[string]interface{}{"status": code, "msg": msg}
	c.ServeJSON()
}

func (c *UserController) loginGet() {
	// fmt.Println(c.TplName)
}

func (c *UserController) loginPost() {
}

func (c *UserController) doesLogin() bool {
	//sess := util.UniqueId()
	auth := c.Ctx.GetCookie("auth")
	if auth == "" {
		return false
	}

	isess := c.GetSession(auth)
	sess, ok := isess.(models.User)
	if !ok {
		return false
	}

	_ = sess
	return true
}

func (c *UserController) getUser() *models.User {
	//sess := util.UniqueId()
	auth := c.Ctx.GetCookie("_s")
	if auth == "" {
		ts := time.Now().UnixNano()
		v := c.getClientIp() + strconv.FormatInt(ts, 10)
		c.Ctx.SetCookie("_s", v, 60)

		sess := &models.User{}
		c.SetSession(v, sess)
		return sess
	}

	isess := c.GetSession(auth)
	sess, ok := isess.(*models.User)
	if !ok {
		panic("not sess")
	}

	return sess
}
