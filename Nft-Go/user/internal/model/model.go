package model

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
	Email      string `json:"email"`
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
