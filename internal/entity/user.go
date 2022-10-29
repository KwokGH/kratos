package entity

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
