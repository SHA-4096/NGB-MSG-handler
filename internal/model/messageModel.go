package model

type Message struct {
	ContentType string `json:"contentType"`
	Body        string `json:"body"`
	TargetUid   string `json:"targetUid"`
}

func CreateMessage(contentType, body, targetUid string) error {
	var newMsg Message
	newMsg.ContentType = contentType
	newMsg.Body = body
	newMsg.TargetUid = targetUid
	err := db.Create(&newMsg).Error
	return err
}

func QueryMessageByUid(Uid string) ([]Message, error) {
	var msgs []Message
	err := db.Find(&msgs, "target_uid=?", Uid).Error
	return msgs, err
}
