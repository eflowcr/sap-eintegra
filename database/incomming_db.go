package database

import "github.com/eflowcr/sap-eintegra/models"

type SapIncommingDB struct{}

func (SapIncommingDB) GetIncommingInterfacesByProcess(processId string) (*[]models.SapIncommingProcess, error) {

	dbRepository, err := GetLocalConnection(DSN)

	if err != nil {
		return nil, err
	}

	defer dbRepository.db.Close()

	var incommingProcess []models.SapIncommingProcess

	result := dbRepository.db.Debug().Where("process_code = ?", processId).Find(&incommingProcess).Error

	if result != nil {
		return nil, result
	}

	return &incommingProcess, nil

}

func (SapIncommingDB) GetQueriesByInterface(InterfaceId int) (*[]models.SapIncommingInterfacesView, error) {

	dbRepository, err := GetLocalConnection(DSN)

	if err != nil {
		return nil, err
	}

	defer dbRepository.db.Close()

	var incommingInterfaces []models.SapIncommingInterfacesView

	result := dbRepository.db.Debug().Where("interface_id = ?", InterfaceId).Find(&incommingInterfaces).Error

	if result != nil {
		return nil, result
	}

	return &incommingInterfaces, nil

}
