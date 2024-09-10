package repository

import "flight_system/internal/domain/entity"

type StaffRepository interface {
	CreateStaff(staff *entity.Staff) (int64, error)
	Login(staff *entity.Staff) (*entity.Staff, error)
}
