package aggregate

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/domain/repository"
	staffRepoImpl "flight_system/internal/infrastructure/repository/impl/staff/database"
)

type StaffAggregate interface {
	CreateStaff(staff *entity.Staff) (int64, error)
	Login(staff *entity.Staff) (*entity.Staff, error)
}

type staffAggregate struct {
	staffRepoImpl repository.StaffRepository
}

func NewStaffAggregate() *staffAggregate {
	return &staffAggregate{
		staffRepoImpl: staffRepoImpl.NewStaffRepoImpl(),
	}
}

func (s *staffAggregate) CreateStaff(staff *entity.Staff) (int64, error) {
	return s.staffRepoImpl.CreateStaff(staff)
}

func (s *staffAggregate) Login(staff *entity.Staff) (*entity.Staff, error) {
	return s.staffRepoImpl.Login(staff)
}
