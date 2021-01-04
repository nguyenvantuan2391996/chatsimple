package entities

type SenderReceive struct {
	id int `gorm:"primarykey"`
	MessageId, SenderId, ReceiveId int
}

func (SenderReceive) TableName() string {
	return "sender_receive"
}
