package controllers

import (
	"ZhaiLuBaiKe/app/models"
	"fmt"
	"github.com/robfig/revel"
)

type Quotation struct {
	*revel.Controller
}

func (q *Quotation) Add() revel.Result {
	email := q.Session["email"]
	nickName := q.Session["nickName"]
	return q.Render(email, nickName)
}

func (q *Quotation) PostAdd(quotation *models.Quotation) revel.Result {
	q.Validation.Required(quotation.Tag).Message("请选择一个标签")
	q.Validation.Required(quotation.Content).Message("摘录内容不能为空")
	q.Validation.Required(quotation.Author).Message("作者不能为空")

	fmt.Println("摘录标签： ", quotation.Tag)
	fmt.Println("摘录被容： ", quotation.Content)
	fmt.Println("原文： ", quotation.Original)
	fmt.Println("作者： ", quotation.Author)

	if q.Validation.HasErrors() {
		q.Validation.Keep()
		q.FlashParams()
		return q.Redirect((*Quotation).Add)
	}

	manager, err := models.NewDbManager()
	if err != nil {
		q.Response.Status = 500
		return q.RenderError(err)
	}
	defer manager.Close()

	err = manager.AddQuotation(quotation)
	if err != nil {
		// q.Validation.Keep()
		// q.FlashParams()
		q.Flash.Error(err.Error())
		return q.Redirect((*Quotation).Add)
	}

	return q.Redirect((*Account).RegisterSuccessful)
	//return q.Redirect((*Quotation).AddSuccessful)
}