package data

import (
	"context"
	"database/sql"
	"github.com/KwokGH/kratos/internal/biz"
	"github.com/KwokGH/kratos/internal/entity/be"
	"github.com/KwokGH/kratos/internal/entity/de"
	"github.com/KwokGH/kratos/pkg/dbo"
	"github.com/KwokGH/kratos/pkg/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper

	dbo.BaseDA
}

func (u *userRepo) CreateUser(ctx context.Context, input *be.RegisterInput) (string, error) {
	newUser := &de.User{
		Base:     de.NewBase(),
		NickName: input.NickName,
		Phone:    input.Phone,
		Password: utils.GetMd5(input.Password),
	}
	if _, err := u.Insert(ctx, newUser); err != nil {
		return "", err
	}

	return newUser.ID, nil
}

func (u *userRepo) FindByPhone(ctx context.Context, phone string) (*de.User, error) {
	cond := &UserCondition{
		Phone: sql.NullString{
			String: phone,
			Valid:  true,
		},
	}
	var userInfo = new(de.User)
	if err := u.First(ctx, cond, userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.IUserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type UserCondition struct {
	Phone sql.NullString

	OrderBy UserOrderBy
	Pager   dbo.Pager

	DeletedAt sql.NullInt64
}

func (u *UserCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if u.Phone.Valid {
		wheres = append(wheres, "phone = ?")
		params = append(params, u.Phone.String)
	}

	if u.DeletedAt.Valid {
		wheres = append(wheres, "deleted_at>0")
	} else {
		wheres = append(wheres, "(deleted_at=0)")
	}

	return wheres, params
}

func (u *UserCondition) GetPager() *dbo.Pager {
	return &u.Pager
}

func (u *UserCondition) GetOrderBy() string {
	return u.OrderBy.ToSQL()
}

type UserOrderBy string

func (o UserOrderBy) ToSQL() string {
	switch o {
	default:
		return "id desc"
	}
}
