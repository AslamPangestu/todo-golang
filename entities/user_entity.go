package entities

type User struct {
	ID       int    `gorm:"primaryKey;column:id"`
	Name     string `gorm:"column:name;size:256"`
	UserID   string `gorm:"column:user_id;size:256;uniqueIndex"`
	Password string `gorm:"column:password;size:256"`
}
