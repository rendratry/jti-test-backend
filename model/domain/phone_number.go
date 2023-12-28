package domain

type PhoneNumber struct {
	ID       int    `gorm:"primaryKey;autoIncrement:true;column:id" json:"id"`
	Number   string `gorm:"column:number" json:"number"`
	Provider string `gorm:"column:provider" json:"provider"`
}

func (PhoneNumber) TableName() string {
	return "phone_number"
}
