package entities

type Note struct {
	ID          int    `gorm:"primaryKey;column:id"`
	ActivityNo  string `gorm:"column:activity_no;size:256"`
	Title       string `gorm:"column:title;size:256"`
	Description string `gorm:"column:description;type:text"`
	Status      int    `gorm:"column:status;"`
	UserID      int    `gorm:"column:user_id;"`
	User        User   `gorm:"foreignKey:UserID;"`
}
