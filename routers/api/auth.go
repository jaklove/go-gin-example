package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jacklove/go-gin-example/models"
	"github.com/jacklove/go-gin-example/pkg/e"
	"github.com/jacklove/go-gin-example/pkg/logging"
	"github.com/jacklove/go-gin-example/pkg/util"
	"net/http"
)

type Auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := Auth{Username: username,Password: password}

	ok,_ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if !ok{
		for _, err := range valid.Errors {
			logging.Info(err.Key,err.Message)
		}
		c.JSON(http.StatusOK,gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data" : data,
		})
	}

	isAuth := models.CheckAuth(username, password)
	if isAuth{
		token, err := util.GenerateToken(username, password)
		if err != nil{
			code = e.ERROR_AUTH_TOKEN
		}else {
			data["token"] = token
			code = e.SUCCESS
		}
	}else {
		code = e.ERROR_AUTH
	}

	c.JSON(http.StatusOK,gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}