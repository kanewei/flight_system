package converter

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/infrastructure/repository/po"
)

func StaffEntityToPo(staff *entity.Staff) *po.Staff {
	return &po.Staff{
		ID:       staff.ID,
		Email:    staff.Email,
		Name:     staff.Name,
		Password: staff.Password,
	}
}

func StaffPoToEntity(staff *po.Staff) *entity.Staff {
	return &entity.Staff{
		ID:       staff.ID,
		Email:    staff.Email,
		Name:     staff.Name,
		Password: staff.Password,
	}
}
