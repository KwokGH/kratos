package be

type LoginInput struct {
	Phone    string
	Password string
}

type LoginOutput struct {
	Token string
}

type RegisterInput struct {
	NickName string
	Phone    string
	Password string
}
