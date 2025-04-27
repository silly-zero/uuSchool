package model

type User struct {
	BaseModel
	OpenID   string   `json:"openID" gorm:"size:36;unique;not null"`
	UserInfo UserInfo `json:"userInfo"`
}

func SelectUserByOpenID(openID string) (*User, error) {
	user := User{}
	err := DB.Where("open_id = ?", openID).Preload("UserInfo").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 创建用户
func InsertUser(user *User) {
	DB.Create(&user)
}
