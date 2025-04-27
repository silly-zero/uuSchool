package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silly-zero/uuSchool/model"
	"github.com/silly-zero/uuSchool/request"
	"github.com/silly-zero/uuSchool/response"
	util "github.com/silly-zero/uuSchool/utils"
	"gorm.io/gorm"
)

const (
	DefaultNickName = "微信用户"
	DefaultAvg      = "URL_ADDRESS"
)

// 微信登录
func WxLogin(ctx *gin.Context) {
	wxLogin := request.ReqWxLogin{}
	err := ctx.ShouldBindJSON(&wxLogin)
	if err != nil {
		response.BindErr(ctx)
		return
	}
	openID, err := util.GetOpenID(wxLogin.Code)
	fmt.Println(openID)
	if err != nil {
		response.Fail(ctx, 500, "获取openId失败！！！")
	}
	userByOpenID, err := model.SelectUserByOpenID(openID)
	//判断是否有该用户
	userID := ""
	userInfo := model.UserInfo{}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user := model.User{
			OpenID: openID,
			UserInfo: model.UserInfo{
				NickName: DefaultNickName,
				Avatar:   DefaultAvg,
			},
		}
		model.InsertUser(&user)
		userID = user.ID
		userInfo = user.UserInfo
	} else {
		userID = userByOpenID.ID
		userInfo = userByOpenID.UserInfo
	}
	// 生成token
	token, err := util.GenToken(userID)
	if err != nil {
		response.Fail(ctx, 500, "token生成失败！！！")
		return
	}
	res := struct {
		Token    string         `json:"token"`
		UserInfo model.UserInfo `json:"userInfo"`
	}{
		Token:    token,
		UserInfo: userInfo,
	}

	response.Success(ctx, res)
}
