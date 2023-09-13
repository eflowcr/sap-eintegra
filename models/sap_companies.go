package models

type SapCompany struct {
	CompanyDB       string `gorm:"column:company_db" json:"company_db"`
	UserName        string `gorm:"column:user_name" json:"user_name"`
	UserPassword    string `gorm:"column:user_password" json:"user_password"`
	ServiceLayerURL string `gorm:"column:service_layer_url" json:"service_layer_url"`
	IsActive        bool   `gorm:"column:is_active" json:"is_active"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (SapCompany) TableName() string {
	return "sap_companies"
}
