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

// Statistics平台数据
type statistics struct {
	StartDate   string `form:"startDate"`
	EndDate     string `form:"endDate"`
	CheckType   string `form:"checktype"`
	PhoneNumber string `form:"phoneNumber"`
}

// Statistics平台数据
func Statistics(ctx *iris.Context) {
	var s statistics
	err := ctx.ReadForm(&s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	ctx.JSON(iris.StatusOK, iris.Map{
		"tzl":   2333,
		"tzje":  3333,
		"tzyh":  4444,
		"chart": [...]int{20, 52, 200, 334, 390, 330, 220, 240},
	})
}

type user_detail struct {
	PhoneNumber string `form:"phoneNumber"`
}

// 用户详情
func UserDetail(ctx *iris.Context) {
	var u user_detail
	err := ctx.ReadForm(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
	ctx.JSON(iris.StatusOK, iris.Map{
		"wx1": [...]string{"zhang", "13412341234", "32", "500", "15-3-5", "已完善", "1231254345"},
		/*"account":      "zhang",
		"phone":        13412341234,
		"balance":      32,
		"earn":         500,
		"registerTime": "15-3-5",
		"infoMsg":      "已完善",
		"id":           1231254345,*/
		"wx2":  [...]string{"zhang", "13412341234", "32", "500", "15-3-5", "已完善", "1231254345"},
		"wx3":  [...]string{"zhang", "13412341234", "32", "500", "15-3-5", "已完善", "1231254345"},
		"wx4":  [...]string{"zhang", "13412341234", "32", "500", "15-3-5", "已完善", "1231254345"},
		"wx5":  [...]string{"zhang", "13412341234", "32", "500", "15-3-5", "已完善", "1231254345"},
		"wx6":  [...]string{"zhang", "13412341234", "32", "500", "15-3-5", "已完善", "1231254345"},
		"wx7":  [...]string{"zhang", "13412341234", "32", "500", "15-3-5", "已完善", "1231254345"},
		"wx8":  [...]string{"zhang", "13412341234", "32", "500", "15-3-5", "已完善", "1231254345"},
		"wx9":  [...]string{"zhang", "13412341234", "32", "500", "15-3-5", "已完善", "1231254345"},
		"wx10": [...]string{"zhang", "13412341234", "32", "500", "15-3-5", "已完善", "1231254345"},
	})
}

type line struct {
	Account  string
	OrderNum string
	Xxx      string
	Limit    string
	Time     string
}
type table struct {
	Lines []*line
}

// 用户账户详情
func UserAccountDetail(ctx *iris.Context) {
	var u statistics
	err := ctx.ReadForm(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)

	var t table
	t.Lines = make([]*line, 0, 10)
	l := &line{
		Account:  "zhang",
		OrderNum: "234234234234",
		Xxx:      "24",
		Limit:    "500",
		Time:     "17-4-3",
	}
	for i := 0; i < 100; i++ {
		t.Lines = append(t.Lines, l)
	}
	ctx.JSON(iris.StatusOK, t)
}

type platform_activity struct {
	Activity  string `form:"activityName"`
	StartDate string `form:"startDate"`
	EndDate   string `form:"endDate"`
}

// 自动派奖获得期次
func TicketAuto(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, iris.Map{
		"qici": [...]int{20170203, 20170204, 20170205, 20170206, 20170207, 20170208, 20170209},
	})

}

//平台活动
func PlatformActivity(ctx *iris.Context) {
	var p platform_activity
	err := ctx.ReadForm(&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	ctx.JSON(iris.StatusOK, iris.Map{
		"name": "test",
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
