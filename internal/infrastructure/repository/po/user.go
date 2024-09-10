package po

type User struct {
	ID       int64  `gorm:"primaryKey;column:id"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	Name     string `gorm:"name"`
}
