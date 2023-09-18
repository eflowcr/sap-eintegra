package logiclayer

import (
	"github.com/eflowcr/sap-eintegra/database"
	"github.com/eflowcr/sap-eintegra/models"
)

type SapCompanyLL struct {
}

func (SapCompanyLL) GetCompanies() (*[]models.SapCompany, error) {

	db := database.SapCompaniesDB{}

	data, err := db.GetCompanies()

	if err != nil {
		return nil, err
	}

	return data, nil

}

func (SapCompanyLL) GetCompany(id string) (*models.SapCompany, error) {

	db := database.SapCompaniesDB{}

	data, err := db.GetCompany(id)

	if err != nil {
		return nil, err
	}

	return data, nil

}
