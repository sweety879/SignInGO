package domain

type Insertuser struct {
	Id       string `json:"id"`
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type Deleteuser struct {
	Id       string `json:"id"`
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type EditUserType struct {
	Email       string `json:"email"       binding:"required"`
	OldPassword string `json:"oldpassword" binding:"required"`
	NewPassword string `json:"newpassword"  binding:"required"`
}
