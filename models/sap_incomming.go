package models

type SapIncommingInterfaces struct {
	InterfaceId          int    `gorm:"column:interface_id;primary_key" json:"interface_id"`
	InterfaceCode        string `gorm:"column:interface_code" json:"interface_code"`
	InterfaceDescription string `gorm:"column:interface_description" json:"interface_description"`
	CompanyDb            string `gorm:"column:company_db" json:"company_db"`
	EflowConnection      string `gorm:"column:eflow_connection" json:"eflow_connection"`
	IsActive             bool   `gorm:"column:is_active" json:"is_active"`
}

// TableName sets the insert table name for this struct type

func (s *SapIncommingInterfaces) TableName() string {
	return "incomming_interfaces"
}
