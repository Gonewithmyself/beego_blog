package controllers

type BookController struct {
	baseController
}

//后台用户登录
func (c *BookController) Login() {
	name := "第一本Docker书"
	c.Data["Name"] = name
	c.TplName = c.controllerName + "/main.html"
}

func (c *BookController) Book() {
	name := "第一本Docker书"
	idx, err := c.GetInt("id")
	println(idx, err)
	c.Data["Name"] = name
	c.TplName = c.controllerName + "/book.tpl"
}
