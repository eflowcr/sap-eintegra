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

func (db *SapCompaniesDB) GetCompany(id string) (*models.SapCompany, error) {

	dbRepository, err := GetLocalConnection(DSN)

	if err != nil {
		return nil, err
	}

	defer dbRepository.db.Close()

	var company models.SapCompany

	result := dbRepository.db.Where("company_db = ?", id).Find(&company).Error

	if result != nil {
		return nil, result
	}

	return &company, nil

}
