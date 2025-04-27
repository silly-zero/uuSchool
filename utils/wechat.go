package util

import (
	"encoding/json"
	"fmt"
	"github.com/silly-zero/uuSchool/configs"
	"net/http"
)

type ResOpenID struct {
	OpenID  string `json:"openid"`
	Unionid string `json:"unionid"`
	Errmsg  string `json:"errmsg"`
}

// 获取openId
func GetOpenID(code string) (string, error) {
	wechat := config.GConfig.Wechat
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url = fmt.Sprintf(url, wechat.AppId, wechat.AppSecret, code)
	// 发起请求
	res, _ := http.Get(url)
	// 成功后获取openId
	resOpenID := ResOpenID{}
	err := json.NewDecoder(res.Body).Decode(&resOpenID)
	if err != nil {
		return "", err
	}
	return resOpenID.OpenID, nil
}
