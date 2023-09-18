package models

type IncommingConnectionsView struct {
	InterfaceId              int    `gorm:"column:interface_id"`
	InterfaceCode            string `gorm:"column:interface_code"`
	InterfaceDescription     string `gorm:"column:interface_description"`
	EflowServer              string `gorm:"column:eflow_server"`
	EflowDatabase            string `gorm:"column:eflow_database"`
	EflowUserName            string `gorm:"column:eflow_user_name"`
	EflowUserPassword        string `gorm:"column:eflow_user_password"`
	EflowConnectionString    string `gorm:"column:eflow_connection_string"`
	EflowUseConnectionString bool   `gorm:"column:eflow_use_connection_string"`
	SapCompanyDb             string `gorm:"column:sap_company_db"`
	SapUserName              string `gorm:"column:sap_user_name"`
	SapUserPassword          string `gorm:"column:sap_user_password"`
	SapServiceLayer          string `gorm:"column:sap_service_layer"`
}

func (IncommingConnectionsView) TableName() string {
	return "view_incomming_connections"
}
