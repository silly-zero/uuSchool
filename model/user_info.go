package model

// 用户信息
type UserInfo struct {
	BaseModel
	UserID   string `json:"userID,omitempty" gorm:"size:36"`
	NickName string `json:"nickName" gorm:"size:20"`
	Avatar   string `json:"avatar" gorm:"size:255"`
}
