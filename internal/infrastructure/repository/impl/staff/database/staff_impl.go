package database

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/global"
	"flight_system/internal/infrastructure/repository/converter"
	"flight_system/pkg/encrypt"

	"gorm.io/gorm"
)

type StaffRepo struct {
	StaffTable string
	db         *gorm.DB
}

func NewStaffRepoImpl() *StaffRepo {
	return &StaffRepo{
		StaffTable: "staffs",
		db:         global.Db,
	}
}

func (s *StaffRepo) CreateStaff(staff *entity.Staff) (int64, error) {
	po := converter.StaffEntityToPo(staff)
	hashedPassword, err := encrypt.EncryptPassword(po.Password)
	if err != nil {
		return 0, err
	}

	po.Password = hashedPassword

	result := s.db.Table(s.StaffTable).Create(&po)
	if result.Error != nil {
		return 0, result.Error
	}
	return staff.ID, nil
}

func (s *StaffRepo) Login(staff *entity.Staff) (*entity.Staff, error) {
	po := converter.StaffEntityToPo(staff)
	result := s.db.Table(s.StaffTable).Where("email = ? AND password = ?", po.Email, po.Password).First(&po)
	if result.Error != nil {
		return nil, result.Error
	}
	return converter.StaffPoToEntity(po), nil
}
