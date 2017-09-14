package handler

import (
	"fmt"
	iris "gopkg.in/kataras/iris.v6"
	"time"

	
)

var (
	auth = "authenticated"
)

// Login ...
func Login(ctx *iris.Context) {
	ctx.MustRender("login.html", nil)
}


// index
func Index(ctx *iris.Context) {
	ctx.MustRender("index.html", nil)
}

type slogin struct {
	Mobile string `form:"mobile"`
	Pwd    string `form:"pwd"`
}

type myTest struct {
	StartDate string  `form:"startDate"`
	EndDate   string  `form:"endDate"`
	CheckType string  `form:"checktype"`
	PhoneNumber string `form:"phoneNumber"`
}

// Test
func Test(ctx *iris.Context) {
	var t myTest
	err := ctx.ReadForm(&t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	ctx.JSON(iris.StatusOK, iris.Map{
		"tzl": 2333,
		"tzje": 3333,
		"tzyh": 4444,
		"chart": [...] int {20, 52, 200, 334, 390, 330, 220, 240},
	})
}

// LoginP ...
func LoginP(ctx *iris.Context) {
	var l slogin
	err := ctx.ReadForm(&l)
	if err != nil {
		fmt.Println("loginp:", err)
		ctx.WriteString("参数错误")
		return
	}

	if l.Pwd != "12345" {
		// ctx.JSON(iris.StatusOK, common.ResData(false, "密码错误", "", "", nil))
		// ctx.WriteString("手机或密码错误")
		ctx.Redirect("/manage/login/")
		time.Sleep(time.Second)
		return
	}
	sess := ctx.Session()
	sess.Set(auth, true)
	sess.Set("name", "555")
	ctx.Redirect("/manage/index/")
}

// Logout ...
func Logout(ctx *iris.Context) {
	sess := ctx.Session()
	sess.Set(auth, false)
}
