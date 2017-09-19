package main

import (
	"backend/handler"

	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

var (
	key = "zc_develop_2008"
)

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())
	app.Adapt(iris.DevLogger())

	sess := sessions.New(sessions.Config{Cookie: key})
	app.Adapt(sess)

	app.Adapt(view.HTML("./templates", ".html").Reload(true))
	app.StaticWeb("/manage/static/", "./static/")
	m := app.Party("/manage")
	{
		m.Get("/logout/", handler.Logout)
		m.Get("/login/", handler.Login)

		// Index
		m.Get("/index/", handler.Index)

		m.Post("/login/", handler.LoginP)
		m.Post("/statistics/", handler.Statistics)
		m.Post("/user_detail/", handler.UserDetail)
		m.Post("/user_account_detail", handler.UserAccountDetail)
		m.Post("/ticket_auto", handler.TicketAuto)
	}
	app.Listen(":8080") //cfg.Port
}
