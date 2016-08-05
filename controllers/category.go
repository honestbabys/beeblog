package controllers

import (
	"beeblog/models"
	"fmt"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {

	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		fmt.Println("name空了")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	}

	var err error
	c.Data["Categories"], err = models.GetAllCategories()

	if err != nil {
		beego.Error(err)
	}
	c.Data["IsCategory"] = true
	c.TplName = "category.html"
}
