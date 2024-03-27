package dto

type LoginInfo struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignupInfo struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}
