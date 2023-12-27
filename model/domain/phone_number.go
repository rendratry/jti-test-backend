package domain

type PhoneNumber struct {
	ID       int    `gorm:"primaryKey;autoIncrement:true;column:id"`
	Number   string `gorm:"column:number"`
	Provider string `gorm:"column:provider"`
}

func (PhoneNumber) TableName() string {
	return "phone_number"
}
