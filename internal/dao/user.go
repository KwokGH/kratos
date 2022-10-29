package dao

import (
	"database/sql"
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type User struct {
	ID                 string    `gorm:"column:id;primary_key"`
	Time               time.Time `gorm:"column:time;default:CURRENT_TIMESTAMP"`
	Name               string    `gorm:"column:name"`
	HeadPic            string    `gorm:"column:head_pic"`
	HeadPicType        int       `gorm:"column:head_pic_type;default:0;NOT NULL"` // 头像类型 0-默认头像 1-藏品图片
	CountryCode        string    `gorm:"column:country_code"`                     // 手机号国家代码
	Mobile             string    `gorm:"column:mobile"`                           // 手机号
	Sex                int       `gorm:"column:sex;default:0;NOT NULL"`           // 性别 0-未填写 1-男 2-女
	Birth              time.Time `gorm:"column:birth"`                            // 生日
	WechatMpOpenid     string    `gorm:"column:wechat_mp_openid"`                 // 微信小程序openid
	WechatMpSessionKey string    `gorm:"column:wechat_mp_session_key"`            // 微信小程序解密key
	Status             int       `gorm:"column:status;default:0;NOT NULL"`        // 状态：0-尚未注册 1-正常
	MetapassID         string    `gorm:"column:metapass_id"`
	Password           string    `gorm:"column:password"`                          // 密码（加密后）
	PasswordSalt       string    `gorm:"column:password_salt"`                     // 密码加密混淆盐
	RealNameAuth       int       `gorm:"column:real_name_auth;default:0;NOT NULL"` // 是否实名认证
	RealName           string    `gorm:"column:real_name"`                         // 真实姓名
	IDCard             string    `gorm:"column:id_card"`                           // 身份证
	AppType            int       `gorm:"column:app_type;default:0;NOT NULL"`       // 最近APP登录类型 0-无 1-安卓 2-iOS
	DeviceID           string    `gorm:"column:device_id"`
	Amount             float64   `gorm:"column:amount;default:0.00;NOT NULL"`        // 总余额
	FreezeAmount       float64   `gorm:"column:freeze_amount;default:0.00;NOT NULL"` // 冻结余额
	WtsSignNum         string    `gorm:"column:wts_sign_num"`                        // 通联会员号
	PayPassword        string    `gorm:"column:pay_password"`                        // 支付密码
	PayPasswordSalt    string    `gorm:"column:pay_password_salt"`                   // 支付密码混淆值
}

func (m *User) TableName() string {
	return "user"
}

type IUserDao interface {
	dbo.DataAccesser
}

type UserDao struct {
	dbo.BaseDA
}

var (
	_UserOnce sync.Once
	_UserDao  IUserDao
)

func GetUserDao() IUserDao {
	_UserOnce.Do(func() {
		_UserDao = &UserDao{}
	})
	return _UserDao
}

type UserCondition struct {
	Mobile sql.NullString

	OrderBy int
	Pager   dbo.Pager
}

func (c UserCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.Mobile.Valid {
		wheres = append(wheres, "mobile =?")
		params = append(params, c.Mobile.String)
	}

	return wheres, params
}

func (c UserCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c UserCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
