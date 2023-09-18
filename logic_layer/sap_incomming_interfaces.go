package logiclayer

import (
	"github.com/eflowcr/sap-eintegra/database"
	"github.com/eflowcr/sap-eintegra/models"
)

type SapIncommingInterfacesLL struct{}

func (SapIncommingInterfacesLL) GetIncommingInterfacesByProcess(processId string) (*[]models.SapIncommingProcess, error) {

	db := database.SapIncommingDB{}

	data, err := db.GetIncommingInterfacesByProcess(processId)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func (SapIncommingInterfacesLL) GetQueriesByInterface(InterfaceId int) (*[]models.SapIncommingInterfacesView, error) {

	db := database.SapIncommingDB{}

	data, err := db.GetQueriesByInterface(InterfaceId)

	if err != nil {
		return nil, err
	}

	return data, nil

}
