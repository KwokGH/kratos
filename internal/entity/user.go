package entity

type User struct {
	ID       string `gorm:"id"`
	UserName string `gorm:"user_name"`
	Password string `gorm:"password"`
	Phone    string `gorm:"phone"`
	NickName string `gorm:"nick_name"`
}

type UserLoginRedisKeyPrefix string

const (
	UserKeyLoginLockPrefix  UserLoginRedisKeyPrefix = "Login:Lock:"
	UserKeyLoginFlag        UserLoginRedisKeyPrefix = "Login:Token:"
	UserKeyLoginFailedCount UserLoginRedisKeyPrefix = "Login:FailedCount:"
)

func (k UserLoginRedisKeyPrefix) String() string {
	return string(k)
}

type SmsCodeRedisKeyPrefix string

const (
	SmsCodeRedisKeyFreq             SmsCodeRedisKeyPrefix = "SmsCode:Freq:"
	SmsCodeRedisKeyCode             SmsCodeRedisKeyPrefix = "SmsCode:Code:"
	SmsCodeRedisKeyMatchFailedCount SmsCodeRedisKeyPrefix = "SmsCode:MatchFailedCount:"
)

func (k SmsCodeRedisKeyPrefix) String() string {
	return string(k)
}
