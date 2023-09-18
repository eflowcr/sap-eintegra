package database

import "github.com/eflowcr/sap-eintegra/models"

type IncommingConnections struct{}

func (IncommingConnections) GetConnections(interface_id int) (*models.IncommingConnectionsView, error) {

	dbRepository, err := GetLocalConnection(DSN)

	if err != nil {
		return nil, err
	}

	defer dbRepository.db.Close()

	var incommingProcess models.IncommingConnectionsView

	result := dbRepository.db.Find(&incommingProcess, "interface_id = ?", interface_id).Error

	if result != nil {
		return nil, result
	}

	return &incommingProcess, nil

}
