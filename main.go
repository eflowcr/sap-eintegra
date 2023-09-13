package main

import (
	"fmt"

	"github.com/eflowcr/sap-eintegra/database"
	logiclayer "github.com/eflowcr/sap-eintegra/logic_layer"
)

func main() {

	database.DSN = fmt.Sprintf("sqlserver://%s:%s@%s:%v?database=%s",
		"usreprac",
		"usreprac",
		"20.118.233.183",
		1433,
		"EINTEGRA_SL")

	cll := logiclayer.SapCompanyLL{}

	companies, err := cll.GetCompanies()

	if err != nil {
		panic(err)
	}

	for _, company := range *companies {
		println(company.CompanyDB)
	}

}
