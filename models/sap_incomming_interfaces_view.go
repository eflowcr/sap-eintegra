package models

type SapIncommingInterfacesView struct {
	InterfaceId          int    `gorm:"column:interface_id;primary_key" json:"interface_id"`
	InterfaceCode        string `gorm:"column:interface_code" json:"interface_code"`
	InterfaceDescription string `gorm:"column:interface_description" json:"interface_description"`
	CompanyDb            string `gorm:"column:company_db" json:"company_db"`
	EflowConnection      string `gorm:"column:eflow_connection" json:"eflow_connection"`
	IsActive             bool   `gorm:"column:is_active" json:"is_active"`
	QueryId              int    `gorm:"column:query_id" json:"query_id"`
	QueryDescription     string `gorm:"column:query_description" json:"query_description"`
	QueryValue           string `gorm:"column:query_value" json:"query_value"`
}

// TableName sets the insert table name for this struct type

func (s *SapIncommingInterfacesView) TableName() string {
	return "view_incomming_interface"
}
