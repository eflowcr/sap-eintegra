package logiclayer

import (
	"github.com/eflowcr/sap-eintegra/database"
	"github.com/eflowcr/sap-eintegra/models"
)

type IncommingConnectionsLL struct{}

func (IncommingConnectionsLL) GetConnections(interface_id int) (*models.IncommingConnectionsView, error) {

	db := database.IncommingConnections{}

	data, err := db.GetConnections(interface_id)

	if err != nil {
		return nil, err
	}

	return data, nil

}
