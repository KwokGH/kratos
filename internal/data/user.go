package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "github.com/KwokGH/kratos/api/v1/user"
	"github.com/go-redis/redis/v8"

	"github.com/KwokGH/kratos/internal/biz"
	"github.com/KwokGH/kratos/internal/dao"
	"github.com/KwokGH/kratos/internal/entity"
	"github.com/KwokGH/kratos/pkg/dbo"
	"github.com/KwokGH/kratos/pkg/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
	"time"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

var (
	ErrAccountLocked       = errors.New("您的账户已被锁定，请2小时后重试，或者使用忘记密码功能解锁")
	ErrAccountLoginFailed  = errors.New("登录失败")
	ErrAccountAlreadyLogin = errors.New("用户已登录")
)

func (u *userRepo) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	if err := u.data.redisClient.Ping(context.Background()).Err(); err != nil {
		fmt.Println("redis连接失败")
		panic(err)
	} else {
		fmt.Println("redis连接成功")
	}
	locKey := entity.UserKeyLoginLockPrefix.String() + req.Mobile
	lockVal, err := u.data.redisClient.Get(ctx, locKey).Result()
	if err == redis.Nil {

	} else if err != nil {
		u.log.Errorw("err", err, "locKey", locKey)
		return nil, err
	} else if lockVal != "" {
		u.log.Errorw("msg", "用户已被锁定", "err", err, "locKey", locKey)
		return nil, ErrAccountLocked
	}

	var userInfos []*dao.User
	err = dao.GetUserDao().Query(ctx, &dao.UserCondition{
		Mobile: sql.NullString{
			String: req.Mobile,
			Valid:  true,
		},
	}, &userInfos)
	if err == dbo.ErrRecordNotFound {
		u.log.Errorw("msg", "用户不存在", "err", err, "req.Mobile", req.Mobile)
		return nil, ErrAccountLoginFailed
	}
	if err != nil {
		u.log.Errorw("err", err, "req.Mobile", req.Mobile)
		return nil, err
	}

	userInfo := userInfos[0]
	failedCount := 0
	// 密码验证
	md5Password := utils.GetMd5(req.Password + userInfo.PasswordSalt)
	logx.Infof("请求密码：%s, 数据库密码：%s", md5Password, userInfo.Password)
	if strings.ToLower(userInfo.Password) != strings.ToLower(md5Password) {
		u.log.Error("密码错误,登录失败")

		failedCount, err = u.processLoginFailed(ctx, req.Mobile)
		if err != nil {
			u.log.Error(err)
			return nil, err
		}

		return nil, ErrAccountLoginFailed
	}

	//if req.Type != 0 {
	//	userInfo.AppType = req.Type
	//	userInfo.DeviceID = req.DeviceId
	//	if _, err := l.svcCtx.UserDao.Update(l.ctx, userInfo); err != nil {
	//		l.Error(err)
	//		return nil, err
	//	}
	//}

	now := time.Now().Unix()
	token, err := u.getJwtToken("123", now, 10000, "999999")
	if err != nil {
		u.log.Error(err)
		return nil, err
	}

	resp := new(pb.LoginReply)
	resp.Authentication = token
	resp.FailedCount = int32(failedCount)
	resp.UserId = "999999"

	return resp, nil
}

func (u *userRepo) getJwtToken(secretKey string, iat int64, seconds int64, userID string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (u *userRepo) processLoginFailed(ctx context.Context, mobile string) (int, error) {
	fckey := entity.UserKeyLoginFailedCount.String() + mobile

	val, err := u.data.redisClient.Get(ctx, fckey).Result()
	if err != nil {
		return 0, err
	}

	failedCount := 0
	if val != "" {
		failedCount, err = strconv.Atoi(val)
		if err != nil {
			return 0, err
		}
	}

	lockKey := entity.UserKeyLoginLockPrefix.String() + mobile
	failedCount++
	if failedCount >= 5 {
		if err = u.data.redisClient.SetEX(ctx, lockKey, "1", 1*time.Minute).Err(); err != nil {
			return 0, err
		}

		if err := u.data.redisClient.Del(ctx, fckey).Err(); err != nil {
			return 0, err
		}

		u.log.Infof("手机号为：%s 的用户名或密码不正确，错误次数太多，您的账户已被锁定，请2小时后重试，或者使用忘记密码功能解锁", mobile)

		return failedCount, nil
	} else {
		if err := u.data.redisClient.SetEX(ctx, fckey, strconv.Itoa(failedCount), 10*time.Minute).Err(); err != nil {
			return 0, err
		}

		return failedCount, nil
	}
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.IUserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
