package controllers

import (
	"ZhaiLuBaiKe/app/models"
	"fmt"
	"github.com/robfig/revel"
)

type Account struct {
	*revel.Controller
}

func (c *Account) Register() revel.Result {
	return c.Render()
}

func (c *Account) RegisterSuccessful() revel.Result {
	return c.Render()
}

func (c *Account) Login() revel.Result {
	return c.Render()
}

func (c *Account) PostRegister(user *models.MockUser) revel.Result {
	c.Validation.Email(user.Email).Message("电子邮件格式无效")
	c.Validation.Required(user.Nickname).Message("用户昵称不能为空")
	c.Validation.Required(user.Password).Message("密码不能为空")
	c.Validation.MinSize(user.Password, 6).Message("密码长度不短于6位")
	c.Validation.Required(user.ConfirmPassword == user.Password).Message("两次输入的密码不一致")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect((*Account).Register)
	}

	manager, err := models.NewDbManager()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer manager.Close()

	err = manager.RegisterUser(user)
	if err != nil {
		// c.Validation.Keep()
		// c.FlashParams()
		c.Flash.Error(err.Error())
		return c.Redirect((*Account).Register)
	}

	return c.Redirect((*Account).RegisterSuccessful)
}

func (c *Account) PostLogin(user *models.LoginUser) revel.Result {
	c.Validation.Email(user.Email).Message("电子邮件格式无效")
	c.Validation.Required(user.Password).Message("请输入密码")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect((*Account).Login)
	}

	manager, err := models.NewDbManager()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer manager.Close()

	var u *models.User
	u, err = manager.LoginUser(user)
	if err != nil {
		//c.Validation.Keep()
		// c.FlashParams()
		c.Flash.Error(err.Error())
		return c.Redirect((*Account).Login)
	}
	c.Session["email"] = u.Email
	c.Session["nickName"] = u.Nickname
	fmt.Println("Login successful with email: ", user.Email)
	fmt.Println("Nickname is: ", u.Nickname)

	return c.Redirect((*App).Index)
}
