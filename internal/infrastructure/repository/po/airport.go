package po

type Airport struct {
	Code      string `gorm:"primaryKey;column:code"`
	Terminal  string `gorm:"primaryKey;column:terminal"`
	City      string `gorm:"city"`
	Name      string `gorm:"name"`
	Available bool   `gorm:"available"`
}
