package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/Response"
	"go_gin/dao"
	"go_gin/forms"
	"go_gin/models"
	"go_gin/utils"
	"time"
)

func HandleUserModelToMap(user *models.User) map[string]interface{} {
	birthday := ""
	if user.Birthday == nil {
		birthday = ""
	} else {
		birthday = user.Birthday.Format("2006-01-02")
	}
	userItemMap := map[string]interface{}{
		"id":        user.ID,
		"nick_name": user.NickName,
		"head_url":  user.HeadUrl,
		"birthday":  birthday,
		"address":   user.Address,
		"desc":      user.Desc,
		"gender":    user.Gender,
		"role":      user.Role,
		"mobile":    user.Mobile,
	}
	return userItemMap
}

// PasswordLogin 登录
func PasswordLogin(c *gin.Context) {
	PasswordLoginForm := forms.PasswordLoginForm{

	}
	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	//// 数字验证码验证失败    store.Verify(验证码id,验证码,验证后是否关闭)
	//if !store.Verify(PasswordLoginForm.CaptchaId, PasswordLoginForm.Captcha, true) {
	//	Response.Err(c, 400, 400, "验证码错误", "")
	//	return
	//}
	//查询数据库是否有改用户
	user, ok := dao.FindUserInfo(PasswordLoginForm.Username, PasswordLoginForm.PassWord)
	if !ok {
		Response.Err(c, 401, 401, "未注册该用户", "")
		return
	}
	//
	token := utils.CreateToken(c, user.ID, user.NickName, user.Role)
	userinfoMap := HandleUserModelToMap(user)
	userinfoMap["token"] = token
	Response.Success(c, 200, "success", userinfoMap)
}

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	// 获取参数
	UserListForm := forms.UserListForm{}
	if err := c.ShouldBind(&UserListForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	total, userlist := dao.GetUserList(UserListForm.Page, UserListForm.PageSize)
	// 判断
	if (total + len(userlist)) == 0 {
		Response.Err(c, 400, 400, "未获取到到数据", map[string]interface{}{
			"total":    total,
			"userlist": userlist,
		})
		return
	}
	Response.Success(c, 200, "获取用户列表成功", map[string]interface{}{
		"total":    total,
		"userlist": userlist,
	})
}

// PutHeaderImage 上传用户头像
func PutHeaderImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileObj, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 把文件上传到minio对应的桶中
	ok := utils.UploadFile("userheader", file.Filename, fileObj, file.Size)
	if !ok {
		Response.Err(c, 401, 401, "头像上传失败", "")
		return
	}
	headerUrl := utils.GetFileUrl("userheader", file.Filename, time.Second*24*60*60)
	if headerUrl == "" {
		Response.Success(c, 400, "获取用户头像失败", "headerMap")
		return
	}
	//TODO 把用户的头像地址存入到对应user表中head_url 中
	Response.Success(c, 200, "头像上传成功", map[string]interface{}{
		"userheaderUrl": headerUrl,
	})
}
