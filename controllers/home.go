package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "home.html"
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics, err := models.GetAllTopics(c.Input().Get("cate"), c.Input().Get("label"), true)
	if err != nil {
		beego.Error(err)
	}

	//无数据时导致的异常怎么解决
	c.Data["Topics"] = topics

	//获取分类数据
	categories, err := models.GetAllCategories() //err居然不是重复定义
	if err != nil {
		beego.Error(err)
	}

	c.Data["Categories"] = categories

}
