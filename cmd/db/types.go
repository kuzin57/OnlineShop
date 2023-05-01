package db

const (
	usersTable = "bshop.user"
)

type User struct {
	Id          uint32 `json:"id"`
	Firstname   string `json:"firstname"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
	Birthday    string `json:"birthday"`
	Password    string `json:"password"`
}
