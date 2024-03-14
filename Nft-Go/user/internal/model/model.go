package model

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	Avatar     string `json:"avatar"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRole struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

type Notice struct {
	Id          int32  `json:"id"`           // 通知ID
	Title       string `json:"title"`        // 标题
	Description string `json:"description"`  // 描述
	PublishTime string `json:"publish_time"` // 发布时间
	UserAddress string `json:"user_address"` // 对应的用户地址
	Type        int32  `json:"type"`         // 类型
}
