package user

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"xframe/internal/service/user/entity"
)

type IUserRepository interface {
	i() //接口中所有方法只能在本包中实现
	Info(context.Context, int) (*entity.User, error)
	GetUserByName(context.Context, string) (*entity.User, error)
	CreateUSer(context.Context, *entity.User) error
}

type Repository struct {
	Db *gorm.DB
}

func New(db *gorm.DB) IUserRepository {
	return &Repository{Db: db}
}
func (*Repository) i() {}

func (d *Repository) Info(ctx context.Context, id int) (res *entity.User, err error) {
	userPo := User{}
	res = &entity.User{}
	err = d.Db.WithContext(ctx).First(&userPo, id).Error
	if err != nil {
		return
	}
	if err = copier.Copy(res, &userPo); err != nil {
		return
	}
	return
}

func (d *Repository) GetUserByName(ctx context.Context, username string) (res *entity.User, err error) {
	userPo := User{}
	err = d.Db.WithContext(ctx).Where("user_name=?", username).First(&userPo).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		return
	}

	res = &entity.User{}
	if err = copier.Copy(res, &userPo); err != nil {
		return
	}
	return
}

func (d *Repository) CreateUSer(ctx context.Context, data *entity.User) (err error) {
	userPo := &User{}
	if err = copier.Copy(userPo, data); err != nil {
		return
	}
	return d.Db.Omit("last_login_time", "create_time", "update_time").Create(userPo).Error
}
