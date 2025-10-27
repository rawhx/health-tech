package dto

type UserCreateRequest struct {
	Nama     string `form:"nama" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=8"`
}

type UserLoginRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required,min=8"`
}

type UserParams struct {
	UserID string `form:"user_id"`
	Email  string `form:"email"`
	Nama   string `form:"password"`
}

type LoginResponse struct {
	User  DataUser `json:"data_user"`
	Token string   `json:"token"`
}

type DataUser struct {
	UserID string `json:"user_id"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
}
