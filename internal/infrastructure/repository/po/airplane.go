package po

import "gorm.io/datatypes"

type Airplane struct {
	ID        int64          `gorm:"primaryKey;column:id"`
	Model     string         `gorm:"model"`
	Seats     datatypes.JSON `gorm:"seats"`
	Available bool           `gorm:"available"`
}
