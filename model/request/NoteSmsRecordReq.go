package request

type NoteSmsRecordReq struct {
	Sms          string    `json:"sms" form:"sms" gorm:"column:sms"`
	Sender       string    `json:"sender" form:"sender" gorm:"column:sender"`
}