package middleware

import (
	commons "blog-go/src/common"
	"blog-go/src/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var optionService = new(service.OptionService)

func StaticPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Index(c.Request.URL.Path, ".") != -1 {
			var uri = ""
			key, err := optionService.GetByKey(commons.OptionWebTheme)
			if err != nil {
				log.Fatalf("StaticPathMiddleware, error: %v", err)
			}
			if strings.Index(c.Request.URL.Path, key.Value+"Static") != -1 {
				uri = strings.ReplaceAll(c.Request.URL.Path, key.Value+"Static", "static/"+key.Value+"/static/")
				c.Redirect(http.StatusMovedPermanently, uri)
			}
		}
		c.Next()
	}
}
