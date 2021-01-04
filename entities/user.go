package entities

type User struct {
	Id       int `gorm:"primarykey"`
	UserName string
	Password string
	Name     string
}

func (User) TableName() string {
	return "user"
}
