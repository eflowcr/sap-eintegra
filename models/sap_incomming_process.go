package models

type SapIncommingProcess struct {
	ProcessId   int    `gorm:"column:process_id;primary_key" json:"process_id"`
	ProcessCode string `gorm:"column:process_code" json:"process_code"`
	InterfaceId int    `gorm:"column:interface_id" json:"interface_id"`
	Description string `gorm:"column:description" json:"description"`
}

// TableName sets the insert table name for this struct type

func (s *SapIncommingProcess) TableName() string {
	return "incomming_process"
}
