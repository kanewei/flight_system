package service

import (
	"flight_system/internal/application/assembler"
	"flight_system/internal/application/model"
	"flight_system/internal/domain/aggregate"
)

type StaffService struct {
	StaffRequest   *assembler.StaffRequest
	StaffResponse  *assembler.StaffResponse
	StaffAggregate aggregate.StaffAggregate
}

func NewStaffService() *StaffService {
	return &StaffService{
		StaffRequest:   assembler.NewStaffRequest(),
		StaffResponse:  assembler.NewStaffResponse(),
		StaffAggregate: aggregate.NewStaffAggregate(),
	}
}

func (s *StaffService) CreateStaff(req *model.CreateStaffRequest) (int64, error) {
	staff := s.StaffRequest.StaffRequestToEntity(req)
	id, err := s.StaffAggregate.CreateStaff(staff)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *StaffService) Login(req *model.StaffLoginRequest) (*model.StaffLoginResponse, error) {
	staff := s.StaffRequest.LoginRequestToEntity(req)
	staffEntity, err := s.StaffAggregate.Login(staff)
	if err != nil {
		return nil, err
	}

	// Todo generate token
	token := ""

	return s.StaffResponse.LoginToResponse(staffEntity, token), nil
}
