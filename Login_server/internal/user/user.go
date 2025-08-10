package user

type User struct {
	ID       int
	Username string `json:"username"`
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}
