package po

type Staff struct {
	ID       int64  `gorm:"primaryKey;column:id"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	Name     string `gorm:"name"`
}
