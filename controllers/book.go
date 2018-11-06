package controllers

import (
	"blog/util"
	"path/filepath"
	"strconv"
	"strings"
)

type BookController struct {
	baseController
}

//后台用户登录
func (c *BookController) Login() {
	name := "第一本Docker书"
	c.Data["Name"] = name
	c.TplName = c.controllerName + "/main.html"
}

const PATH = ""

var lastIdx int

func (c *BookController) Book() {
	//name := "第一本Docker书"
	str := c.GetString("id")

	str = strings.TrimSpace(str)
	idx, err := strconv.ParseInt(str, 10, 32)
	if 0 != idx {
		lastIdx = int(idx)
	}

	dir := getPdfs()
	if nil != err {
		//return
	}

	// for _, _ := range dir {
	// 	//println(tmp)
	// }

	book := dir[lastIdx]
	_, book = filepath.Split(book)
	//fmt.Println(dir, book)
	println(str, err, idx, len(str), c.Ctx.Input.URI())
	c.Data["Name"] = book
	c.TplName = c.controllerName + "/book.tpl"
}

func getPdfs() []string {
	dir, err := util.ListDir(`G:\02 go\src\blog\static\books`, "pdf")
	if nil != err {
		return nil
	}

	var temp []string = make([]string, len(dir)-8)

	copy(temp, dir[:10])

	//fmt.Println(temp, dir)
	copy(temp[10:], dir[18:])
	//fmt.Println(temp, dir)
	var tmp = make([]string, 8)
	copy(tmp, dir[10:18])
	temp = append(temp, tmp...)
	return temp
}
