package database

import "github.com/eflowcr/sap-eintegra/models"

type SapCompaniesDB struct{}

var DSN string

func (db *SapCompaniesDB) GetCompanies() (*[]models.SapCompany, error) {

	dbRepository, err := GetLocalConnection(DSN)

	if err != nil {
		return nil, err
	}

	defer dbRepository.db.Close()

	var companies []models.SapCompany

	dbRepository.db.Find(&companies)

	return &companies, nil

}
