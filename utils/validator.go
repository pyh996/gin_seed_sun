package utils

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_gin/Response"
	"go_gin/global"
	"net/http"
	"regexp"
	"strings"
)

// HandleValidatorError 处理字段校验异常
func HandleValidatorError(c *gin.Context, err error) {
	//如何返回错误信息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		Response.Err(c, http.StatusInternalServerError, 500, "字段校验错误", err.Error())
	}
	msg := removeTopStruct(errs.Translate(global.Trans))
	Response.Err(c, http.StatusBadRequest, 400, "字段校验错误", msg)
	return
}

//   removeTopStruct定义一个去掉结构体名称前缀的自定义方法：
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		// 从文本的逗号开始切分   处理后"mobile": "mobile为必填字段"  处理前: "PasswordLoginForm.mobile": "mobile为必填字段"
		rsp[field[strings.Index(field, ".")+1:]] = err
		color.Red(err)
	}
	return rsp
}

// ValidateMobile 校验手机号
func ValidateMobile(fl validator.FieldLevel) bool {
	// 利用反射拿到结构体tag含有mobile的key字段
	mobile := fl.Field().String()
	//使用正则表达式判断是否合法
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
