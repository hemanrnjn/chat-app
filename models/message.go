import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	Id        int    `json:"id"`
	User_Id   int    `json:"user_id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type MessageRecipient struct {
	Id           int  `json:"id"`
	Recepient_Id int  `json:"recepient_id"`
	Message_Id   int  `json:"message_id"`
	Is_Read      bool `json:"is_read"`
}