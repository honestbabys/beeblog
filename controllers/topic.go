package controllers

import (
	"beeblog/models"
	"path"
	"strings"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (l *TopicController) Get() {
	l.Data["IsLogin"] = checkAccount(l.Ctx)
	l.Data["IsTopic"] = true
	l.TplName = "topic.html"
	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err)
	} else {
		l.Data["Topics"] = topics
	}
}
func (l *TopicController) Post() {
	if !checkAccount(l.Ctx) {
		l.Redirect("/login", 302)
		return
	}

	//解析POST获得的表单数据
	title := l.Input().Get("title")
	content := l.Input().Get("content")
	categary := l.Input().Get("category")
	tid := l.Input().Get("tid")
	label := l.Input().Get("label")

	//判断用户是否上传附件
	_, fh, err := l.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	var attachment string

	if fh != nil {
		//上传了附件
		attachment = fh.Filename
		beego.Info(attachment)
		err = l.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err = models.AddTopic(title, content, label, categary, attachment)
	} else {
		err = models.ModifyTopic(tid, title, content, label, categary, attachment)
	}

	if err != nil {
		beego.Error(err)
	}
	l.Redirect("/topic", 302)
}

func (l *TopicController) Add() {
	l.TplName = "topic_add.html"
}

func (l *TopicController) Modify() {
	if !checkAccount(l.Ctx) {
		l.Redirect("/login", 302)
		return
	}
	l.TplName = "topic_modify.html"
	tid := l.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		l.Redirect("/", 302)
		return
	}
	l.Data["Topic"] = topic
	l.Data["Tid"] = tid
}

func (l *TopicController) View() {
	l.TplName = "topic_view.html"
	//智能路由 参数会被解析到这个map中 key为0即为第一个参数
	topic, err := models.GetTopic(l.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		l.Redirect("/", 302)
		return
	}
	l.Data["Topic"] = topic
	l.Data["Lables"] = strings.Split(topic.Lables, " ")
	l.Data["Tid"] = l.Ctx.Input.Param("0") //这种适用于Post获取数据
	replies, err := models.GetAllReplies(l.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}

	l.Data["Replies"] = replies
	l.Data["IsLogin"] = checkAccount(l.Ctx)
}

func (l *TopicController) Delete() {
	if !checkAccount(l.Ctx) {
		l.Redirect("/login", 302)
		return
	}
	err := models.DelTopic(l.Input().Get("tid")) //这种方式适用于Get方式传递数据
	if err != nil {
		beego.Error(err)
	}
	l.Redirect("/", 302)
}
