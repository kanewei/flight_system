package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"
)

type StaffResponse struct{}

func NewStaffResponse() *StaffResponse {
	return new(StaffResponse)
}

func (s *StaffResponse) StaffToResponse(staff *entity.Staff) *model.CreateStaffResponse {
	return &model.CreateStaffResponse{
		ID: staff.ID,
	}
}

func (s *StaffResponse) LoginToResponse(staff *entity.Staff, token string) *model.StaffLoginResponse {
	return &model.StaffLoginResponse{
		ID:    staff.ID,
		Name:  staff.Name,
		Token: token,
	}
}
