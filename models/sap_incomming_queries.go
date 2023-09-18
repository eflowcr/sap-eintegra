package models

type SapIncommingQueries struct {
	QueryId          int    `gorm:"column:query_id;primary_key" json:"query_id"`
	InterfaceId      int    `gorm:"column:interface_id" json:"interface_id"`
	QueryDescription string `gorm:"column:query_description" json:"query_description"`
	QueryValue       string `gorm:"column:query_value" json:"query_value"`
}

// TableName sets the insert table name for this struct type
func (s *SapIncommingQueries) TableName() string {
	return "incomming_interfaces_queries"
}
