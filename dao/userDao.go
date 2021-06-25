package dao

import (
	"fmt"
	"go_gin/global"
	"go_gin/models"
)

var users []models.User
var user models.User

// GetUserList 获取用户列表(page第几页,page_size每页几条数据)
func GetUserList(page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据
	userList := make([]interface{}, 0, len(users))
	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Offset(offset).Limit(page_size).Find(&users)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, userList
	}
	// 获取user总数
	total := len(users)
	// 查询数据
	result.Offset(offset).Limit(page_size).Find(&users)
	//
	for _, useSingle := range users {
		birthday := ""
		if useSingle.Birthday == nil {
			birthday = ""
		}else {
			// 给未设置生日的初始值
			birthday = useSingle.Birthday.Format("2006-01-02")
		}
		//
		userItemMap := map[string]interface{}{
			"id":        useSingle.ID,
			"nick_name": useSingle.NickName,
			"head_url":  useSingle.HeadUrl,
			"birthday":  birthday,
			"address":   useSingle.Address,
			"desc":      useSingle.Desc,
			"gender":    useSingle.Gender,
			"role":      useSingle.Role,
			"mobile":    useSingle.Mobile,
		}
		fmt.Println(userItemMap)
		userList = append(userList, userItemMap)
	}
	return total, userList
}

// UsernameFindUserInfo 通过username找到用户信息
func FindUserInfo(username string, password string) (*models.User, bool) {
	// 查询用户
	rows := global.DB.Where(&models.User{NickName: username, Password: password}).Find(&user)
	fmt.Println(&user)
	if rows.RowsAffected < 1 {
		return &user, false
	}
	return &user, true
}
