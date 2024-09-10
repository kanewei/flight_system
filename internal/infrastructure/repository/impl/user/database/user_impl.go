package database

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/global"
	"flight_system/internal/infrastructure/repository/converter"
	"flight_system/internal/infrastructure/repository/po"
	"flight_system/pkg/encrypt"

	"gorm.io/gorm"
)

type UserRepo struct {
	UserTable string
	db        *gorm.DB
}

func NewUserRepoImpl() *UserRepo {
	return &UserRepo{
		UserTable: "user",
		db:        global.Db,
	}
}

func (u *UserRepo) SignUp(user *entity.SignUp) (int64, error) {
	po := converter.SignUpToPo(user)
	hashedPassword, err := encrypt.EncryptPassword(po.Password)
	if err != nil {
		return 0, err
	}

	po.Password = hashedPassword
	result := u.db.Table(u.UserTable).Create(&po)
	if result.Error != nil {
		return 0, err
	}
	return po.ID, nil
}

func (u *UserRepo) Login(login *entity.Login) (*entity.User, error) {
	var user po.User

	hashedPassword, err := encrypt.EncryptPassword(login.Password)
	if err != nil {
		return nil, err
	}

	password := hashedPassword
	result := u.db.Table(u.UserTable).Where("email = ? AND password = ?", login.Email, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return converter.UserPoToEntity(&user), nil
}
