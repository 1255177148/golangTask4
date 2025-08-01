package po

type User struct {
	ID                 uint   `gorm:"primary_key" json:"id"`
	UserName           string `json:"userName"`
	Password           string `json:"password"`
	Email              string `json:"email"`
	AuthenticationFlag string `json:"authenticationFlag"`
}
