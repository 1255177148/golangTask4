package po

type User struct {
	ID                 uint    `gorm:"primary_key" json:"id" form:"id"`
	UserName           string  `json:"userName" form:"userName"`
	Password           string  `json:"password" form:"password"`
	Email              *string `json:"email" form:"email"`
	AuthenticationFlag *string `json:"authenticationFlag" form:"authenticationFlag"`
}
