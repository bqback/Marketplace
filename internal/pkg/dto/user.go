package dto

type NewUserInfo struct {
	ID    uint64 `json:"id"`
	Login string `json:"login"`
	Token string `json:"-"`
}

type DBUser struct {
	ID           uint64 `json:"id"`
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
}
