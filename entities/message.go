package entities

type Message struct {
	Id int `gorm:"primarykey"`
	MessageContent string
	//MessageTime time.Time
	IsRead string
}

func (Message) TableName() string {
	return "message"
}
