package routers

import (
	commons "blog-go/src/common"
	"fmt"
	"github.com/gin-gonic/gin"
	b64 "github.com/mojocn/base64Captcha"
	"net/http"
)

type SystemRouter struct {
	BaseRouter
	commons.Response
}

func (s *SystemRouter) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "default/index.tmpl", gin.H{})
}
func (s *SystemRouter) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "default/login.tmpl", gin.H{})
}

func (s *SystemRouter) Captcha(c *gin.Context) {
	// 以下是生成验证码的代码
	driver := b64.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := b64.NewCaptcha(driver, b64.DefaultMemStore)
	id, b64s, err := captcha.Generate()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.SetCookie("captcha_id", id, 300, "/", "127.0.0.1", false, true)
	//c.Data(200, "image/png", []byte(b64s))
	c.String(http.StatusOK, b64s)
	//s.SuccessData(c, b64s)
}
