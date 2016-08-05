package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (r *ReplyController) Get() {

}

func (r *ReplyController) Add() {
	tid := r.Input().Get("tid")
	err := models.AddReply(tid, r.Input().Get("nickname"), r.Input().Get("content"))
	if err != nil {
		beego.Error(err)
		r.Redirect("/", 302)
		return
	}

	r.Redirect("/topic/view/"+tid, 302)
}

func (r *ReplyController) Del() {
	if !checkAccount(r.Ctx) {
		return
	}
	tid := r.Input().Get("tid")
	err := models.DeleteReply(r.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}
	r.Redirect("/topic/view/"+tid, 302)

}
