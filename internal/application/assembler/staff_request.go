package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"
)

type StaffRequest struct{}

func NewStaffRequest() *StaffRequest {
	return new(StaffRequest)
}

func (s *StaffRequest) StaffRequestToEntity(req *model.CreateStaffRequest) *entity.Staff {
	return &entity.Staff{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}
}

func (s *StaffRequest) LoginRequestToEntity(req *model.StaffLoginRequest) *entity.Staff {
	return &entity.Staff{
		Email:    req.Email,
		Password: req.Password,
	}
}
